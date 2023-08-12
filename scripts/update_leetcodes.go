package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// snakeToUcfirst converts a snake case string to an ucfirst english string
func snakeToUcfirst(s string) string {
	// Split the string by underscores
	parts := strings.Split(s, "-")
	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	// Join the parts with spaces
	return strings.Join(parts, " ")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Get the input and output paths from command-line arguments
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run update_leetcodes.go input_path output_path")
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	// Create the output directory if it does not exist
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		err = os.Mkdir(outputPath, 0755)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			os.Exit(1)
		}
	}

	problems, err := readProblems("leetcodes.json")
	if err != nil {
		log.Fatal(err)
	}
	mkdir(outputPath + "/leetcode-hard")
	mkdir(outputPath + "/leetcode-medium")
	mkdir(outputPath + "/leetcode-easy")

	// Walk through the input directory and find all folders with main.py files
	err = filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != inputPath {
			// Check if the folder contains a main.py file
			mainPyPath := filepath.Join(path, "main.py")
			if _, err := os.Stat(mainPyPath); err == nil {
				// Read the contents of the main.py file
				data, err := ioutil.ReadFile(mainPyPath)
				if err != nil {
					log.Fatal(err)
				}

				if isExerciseNotReadyForPublish(string(data)) {
					return nil
				}

				// Get the name of the folder
				folderName := filepath.Base(path)

				if _, ok := problems[folderName]; !ok {
					log.Printf("Couldn't find problem with name [%v]!!\n", folderName)
					return nil
				}
				problem := problems[folderName]

				// Create a markdown file with the same name in the output directory
				mdPath := filepath.Join(outputPath+fmt.Sprintf("/leetcode-%v", strings.ToLower(problem.Difficulty)), folderName+".md")
				f, err := os.Create(mdPath)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				f.WriteString(createMarkdownContent(folderName, string(data), problems))
			}
			return filepath.SkipDir // Skip subdirectories
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through input directory:", err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}

func isExerciseNotReadyForPublish(content string) bool {
	firstLine := strings.ToLower(strings.Split(content, "\n")[0])
	notReadyMarkers := []string{"wip", "not ready", "unpublished"}

	for _, marker := range notReadyMarkers {
		if strings.Contains(firstLine, marker) {
			return true
		}
	}
	return false
}

func mkdir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(path, os.ModePerm)
	}
	return nil
}

func createMarkdownContent(problemKebabName, pythonFileContent string, problems map[string]Problem) string {
	aux := `---
title: %s
date: 2022-11-20T09:03:20-08:00
---

%s

%s

## Algorithm

%spython
%s
%s

%s
`

	preMarkdown, contents, postMarkdown := splitFile(pythonFileContent)

	link := ""
	if problem, ok := problems[problemKebabName]; ok {
		link = problem.Link
	}

	return fmt.Sprintf(aux, snakeToUcfirst(problemKebabName), link, preMarkdown, "```", contents, "```", postMarkdown)
}

// splitFile splits a string that represents a file with Python code and comments
// into three strings: pre-markdown, contents and post-markdown.
// Pre-markdown contains the initial lines of the file that are prefixed by a hash and a space,
// but with those initial characters removed. It stops at the first line that does not have that prefix.
// Contents contains all lines after that, until the post-markdown part.
// Post-markdown contains the final lines of the file that are prefixed by a hash and a space,
// but with those initial characters removed.
func splitFile(pythonFileContent string) (preMarkdown, contents, postMarkdown string) {
	// Split the file into lines
	lines := strings.Split(pythonFileContent, "\n")

	// Find the index of the first line that is not pre-markdown
	i := 0
	for i < len(lines) && strings.HasPrefix(lines[i], "#") {
		i++
	}
	// Also remove empty lines before start of code
	for i < len(lines) && lines[i] == "" {
		i++
	}

	// Join the pre-markdown lines
	preMarkdown = strings.Join(stripPrefixes(lines[:i]), "\n")

	// Find the index of the last line that is not post-markdown
	j := len(lines) - 2 // -2 because the last line is always empty
	for j >= i && strings.HasPrefix(lines[j], "#") {
		j--
	}

	// Join the post-markdown lines
	postMarkdown = strings.Join(stripPrefixes(lines[j+1:]), "\n")

	// Join the contents lines
	contents = strings.Join(lines[i:j+1], "\n")

	return
}

func stripPrefixes(lines []string) []string {
	stripped := []string{}
	for _, line := range lines {
		stripped = append(stripped, stripPrefix(line))
	}
	return stripped
}

func stripPrefix(line string) string {
	if line == "#" {
		return ""
	}
	if strings.HasPrefix(line, "# ") {
		return line[2:]
	}
	return line
}

// Problem represents a leetcode problem
type Problem struct {
	Number         int     `json:"number"`
	Name           string  `json:"name"`
	KebabName      string  `json:"kebabName"`
	Link           string  `json:"link"`
	AcceptanceRate float64 `json:"acceptanceRate"`
	Difficulty     string  `json:"difficulty"`
}

// readProblems reads a json file into a map of problems
func readProblems(filename string) (map[string]Problem, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content as bytes
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the json array into a slice of anonymous structs
	var problems []struct {
		Name           string `json:"name"`
		AcceptanceRate string `json:"acceptanceRate"`
		Difficulty     string `json:"difficulty"`
	}
	err = json.Unmarshal(data, &problems)
	if err != nil {
		return nil, err
	}

	// Create a map of problems with the name as the key
	result := make(map[string]Problem)
	for _, p := range problems {
		// Split the name by ". " to get the number and the name
		split := strings.Split(p.Name, ". ")
		number, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		name := split[1]

		// Convert the name to kebab case and replace any single quotes with dashes
		kebabName := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name, " ", "-"), "'", "-"))

		// Prepend the leetcode url to the kebab name to get the link
		link := "https://leetcode.com/problems/" + kebabName

		// Parse the acceptance rate as a float
		ar, err := strconv.ParseFloat(p.AcceptanceRate[:len(p.AcceptanceRate)-1], 64)
		if err != nil {
			return nil, err
		}

		// Create a problem struct and add it to the map
		result[kebabName] = Problem{
			Number:         number,
			Name:           name,
			KebabName:      kebabName,
			Link:           link,
			AcceptanceRate: ar,
			Difficulty:     p.Difficulty,
		}
	}

	return result, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// go run update_leetcodes.go /Users/marianol/Code/leetcode/2022/ ../content/
