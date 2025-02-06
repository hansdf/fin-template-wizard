package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
)

type MessageSection struct {
	Name       string `json:"Name"`
	Content    string `json:"Content"`
	IsSelected bool   `json:"IsSelected"`
}

// Data structure to hold placeholder values
type TemplateData struct {
	Name string
}

func main() {
	generateMessage()
}

func generateMessage() {
	// Load templates
	sections, err := loadTemplates("templates/default.json")
	if err != nil {
		fmt.Println("Error loading templates:", err)
		return
	}

	// Display template options
	fmt.Println("\nSelect sections to include in your message:")
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

	// Collect placeholder values
	data := TemplateData{}
	fmt.Print("Enter recipient's name: ")
	data.Name, _ = reader.ReadString('\n')
	data.Name = strings.TrimSpace(data.Name)

	// Combine selected sections and replace placeholders
	var message string
	for _, section := range selectedSections {
		tmpl, err := template.New("message").Parse(section.Content)
		if err != nil {
			fmt.Printf("Error parsing template for section '%s': %v\n", section.Name, err)
			continue
		}

		var result strings.Builder
		if err := tmpl.Execute(&result, data); err != nil {
			fmt.Printf("Error executing template for section '%s': %v\n", section.Name, err)
			continue
		}

		message += result.String() + "\n\n"
	}

	fmt.Println("\nGenerated Message:")
	fmt.Println(message)

	// Copy to clipboard
	if err := clipboard.WriteAll(message); err != nil {
		fmt.Println("Error copying message to clipboard:", err)
	} else {
		fmt.Println("Message copied to clipboard! You can now paste it anywhere.")
	}
}

func loadTemplates(filename string) ([]MessageSection, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var sections []MessageSection
	if err := json.Unmarshal(data, &sections); err != nil {
		return nil, err
	}

	return sections, nil
}
