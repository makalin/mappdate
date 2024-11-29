package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/manifoldco/promptui"
)

type Update struct {
	Name        string
	Version     string
	Size        string
	Description string
}

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("This program needs to be run with sudo privileges")
		os.Exit(1)
	}

	updates, err := listAvailableUpdates()
	if err != nil {
		fmt.Printf("Error getting updates: %v\n", err)
		os.Exit(1)
	}

	if len(updates) == 0 {
		fmt.Println("No updates available")
		os.Exit(0)
	}

	selected := selectUpdates(updates)
	if len(selected) > 0 {
		installUpdates(selected)
	}
}

func listAvailableUpdates() ([]Update, error) {
	cmd := exec.Command("mappdate", "-l")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list updates: %v", err)
	}

	return parseUpdateOutput(string(output))
}

func parseUpdateOutput(output string) ([]Update, error) {
	var updates []Update
	scanner := bufio.NewScanner(strings.NewReader(output))
	var currentUpdate *Update

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "* ") {
			if currentUpdate != nil {
				updates = append(updates, *currentUpdate)
			}
			currentUpdate = &Update{
				Name: strings.TrimPrefix(line, "* "),
			}
		} else if currentUpdate != nil {
			if strings.HasPrefix(line, "Version: ") {
				currentUpdate.Version = strings.TrimPrefix(line, "Version: ")
			} else if strings.HasPrefix(line, "Size: ") {
				currentUpdate.Size = strings.TrimPrefix(line, "Size: ")
			} else if line != "" {
				if currentUpdate.Description == "" {
					currentUpdate.Description = line
				} else {
					currentUpdate.Description += " " + line
				}
			}
		}
	}

	if currentUpdate != nil {
		updates = append(updates, *currentUpdate)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning output: %v", err)
	}

	return updates, nil
}

func selectUpdates(updates []Update) []Update {
	var items []string
	for _, update := range updates {
		items = append(items, fmt.Sprintf("%s (Version: %s, Size: %s)", 
			update.Name, update.Version, update.Size))
	}

	prompt := promptui.Select{
		Label: "Select updates to install (use space to select, enter to confirm)",
		Items: items,
		Size:  20,
	}

	var selected []Update
	for {
		idx, _, err := prompt.Run()
		if err != nil {
			break
		}
		selected = append(selected, updates[idx])
	}

	return selected
}

func installUpdates(updates []Update) {
	fmt.Println("\nInstalling selected updates...")
	
	for _, update := range updates {
		fmt.Printf("Installing %s...\n", update.Name)
		cmd := exec.Command("mappdate", "-i", update.Name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error installing %s: %v\n", update.Name, err)
			continue
		}
		
		fmt.Printf("Successfully installed %s\n", update.Name)
	}
}
