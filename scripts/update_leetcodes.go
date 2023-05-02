package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	// Walk through the input directory and find all folders with main.py files
	err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
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
				// Get the name of the folder
				folderName := filepath.Base(path)
				// Create a markdown file with the same name in the output directory
				mdPath := filepath.Join(outputPath, folderName+".md")
				f, err := os.Create(mdPath)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				f.WriteString(createMarkdownContent(snakeToUcfirst(folderName), string(data)))
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

func createMarkdownContent(problemName, pythonFileContent string) string {
	aux := `---
title: %s
date: 2022-11-20T09:03:20-08:00
---

%s

## Algorithm

%spython
%s
%s

%s
`

	preMarkdown, contents, postMarkdown := splitFile(pythonFileContent)

	return fmt.Sprintf(aux, problemName, preMarkdown, "```", contents, "```", postMarkdown)
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
	j := len(lines) - 1
	for j >= i && strings.HasPrefix(lines[j], "# ") {
		j--
	}

	// Join the post-markdown lines
	postMarkdown = strings.Join(stripPrefixes(lines[j+1:]), "\n")

	// Join the contents lines
	contents = strings.Join(lines[i:j+1], "\n")

	// Remove the hash and space prefix from pre-markdown and post-markdown
	preMarkdown = strings.Replace(preMarkdown, "# ", "", 1)
	postMarkdown = strings.Replace(postMarkdown, "# ", "", 1)

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

// go run update_leetcodes.go /Users/marianol/Code/leetcode/2022/ ../content/leetcode
