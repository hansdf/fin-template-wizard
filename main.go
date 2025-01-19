package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MessageSection struct {
	Name       string `json:"Name"`
	Content    string `json:"Content"`
	IsSelected bool   `json:"IsSelected"`
}

func main() {
	// Load templates
	data, err := os.ReadFile("templates/default.json")
	if err != nil {
		fmt.Println("Error reading template file:", err)
		os.Exit(1)
	}

	var sections []MessageSection
	if err := json.Unmarshal(data, &sections); err != nil {
		fmt.Println("Error parsing templates:", err)
		os.Exit(1)
	}

	// Display template options
	fmt.Println("Select sections to include in your message:")
	for i, section := range sections {
		fmt.Printf("[%d] %s\n", i+1, section.Name)
	}

	fmt.Print("Enter the numbers of the sections you want to include (e.g., 1,3): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	selectedIndexes := strings.Split(input, ",")
	var selectedSections []MessageSection

	for _, index := range selectedIndexes {
		i, err := strconv.Atoi(index)
		if err == nil && i > 0 && i <= len(sections) {
			selectedSections = append(selectedSections, sections[i-1])
		}
	}

	// Combine selected sections
	fmt.Println("\nGenerated Message:")
	for _, section := range selectedSections {
		fmt.Println(section.Content)
	}
}
