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
		fmt.Println("Usage: go run script.go input_path output_path")
		os.Exit(1)
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
				problemName := snakeToUcfirst(folderName)
				// Write the contents of the main.py file wrapped in a code tag
				f.WriteString("---\ntitle: \"")
				f.WriteString(problemName)
				f.WriteString("\"\ndate: 2022-11-20T09:03:20-08:00\n---\n\n")
				f.WriteString("```python\n")
				f.Write(data)
				f.WriteString("\n```\n")
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
