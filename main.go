package main

import (
	"fmt"
	"strings"
)

const (
	apiURL              = "https://api.github.com/octocat"
	tokenPath           = "/.config/gh-pradar/"
	tokenFile           = "config.json"
)

var (
    quitting bool = false
)

func main() {
	// Set the alternate screen for output aesthetics
	// enterAltScreen()

	// currentToken, err := loadToken()
	// if err != nil {
	// 	print(errorText("Error loading token: " + err.Error()))
	// 	return
	// }

	// if currentToken == nil {
	// 	currentToken, err = oauth2Flow()
	// 	if err != nil {
	// 		print(errorText("Error in OAuth 2.0 flow: " + err.Error()))
	// 		return
	// 	}

	// 	err = saveToken(currentToken)
	// 	if err != nil {
	// 		print(errorText("Error saving token: " + err.Error()))
	// 		return
	// 	}
	// } else {
	// 	var refreshToken, err = oauth2RefreshFlow(currentToken)
	// 	if err != nil {
	// 		print(errorText("Error refreshing OAuth 2.0: " + err.Error()))
	// 		return
	// 	}

	// 	if refreshToken != nil {
	// 		err = saveToken(refreshToken)
	// 		if err != nil {
	// 			print(errorText("Error saving token (refresh): " + err.Error()))
	// 			return
	// 		}
	// 		currentToken, err = loadToken()
	// 		if err != nil {
	// 			print(errorText("Error loading token (refresh): " + err.Error()))
	// 			return
	// 		}
	// 	}
	// }

	// Load modules
	// err = startRealtimeLoader("Loading modules...", requestModules)
	// if err != nil {
	// 	print(errorText("Error loading modules: " + err.Error()))
	// 	return
	// }

	if quitting {
		print(warnText("Exiting module loading..."))
		return
	}

	// // Obtains modes
	// err = requestModes()
	// if err != nil {
	// 	print(errorText("Error getting modes: " + err.Error()))
	// 	return
	// }
	// // Main loop to allow repeated selection until user quits with 'q'
	// for {
	// 	// Print modes
	// 	result, err := promptModes()
	// 	if err != nil {
	// 		print(errorText("Error loading modes: " + err.Error()))
	// 		return
	// 	}
	// 	if quitting {
	// 		print(warnText("Exiting mode selection..."))
	// 		return
	// 	}

	// 	// Get the selected mode
	// 	for _, mode := range modeList {
	// 		if strings.Contains(result, mode.Name) {
	// 			selectedMode = mode
	// 		}
	// 	}

	// 	// Confirm selected mode
	// 	if selectedMode == (Mode{}) {
	// 		print(errorText("Error getting mode..."))
	// 		return
	// 	}

	// 	switch selectedMode.Id {
	// 	case 0:
	// 		startPromptingOG()
	// 	case 1:
	// 		startPrompting()
	// 	default:
	// 		fmt.Println("No valid choice made.")
	// 		return
	// 	}
	// }
}

// Function to check for the `q` key in any input and exit the program
func checkForQuit(input string) bool {
	if strings.ToLower(input) == "q" {
		quitting = true
		print(warnText("Exiting program..."))
		return true
	}
	return false
}

// Original startPrompting function
// func startPromptingOG() {
// 	// Prompt modules
// 	result, err := promptModules()
// 	if err != nil {
// 		print(errorText("Module prompt failed: " + err.Error()))
// 		return
// 	}

// 	// Get the selected module
// 	for _, module := range catalog {
// 		if strings.Contains(result, module.Name) {
// 			selectedModule = module
// 		}
// 	}

// 	// Confirm selected module
// 	if selectedModule == (Module{}) {
// 		print(errorText("Error getting module..."))
// 		return
// 	}

// 	var selectedTag = ""
// 	if selectTagMode {
// 		err = startRealtimeLoader("Loading tags...", requestTags)
// 		if err != nil {
// 			print(errorText("Error loading tags: " + err.Error()))
// 			return
// 		}
// 		if quitting {
// 			return
// 		}

// 		selectedTag, err = promptTags()
// 		if err != nil {
// 			print(errorText("Tag prompt failed: " + err.Error()))
// 			return
// 		}
// 		if quitting {
// 			return
// 		}
// 	}

// 	err = startFileLoader(selectedModule, selectedTag)
// 	if err != nil {
// 		print(errorText("Error loading file: " + err.Error()))
// 		return
// 	}

// 	fmt.Println("  " + infoTitle("File module."+selectedModule.Name+".tf created!"))
// }
// func startPrompting() {
// 	// Clear the screen before starting the cloning process
// 	enterAltScreen()

// 	// Prompt modules
// 	result, err := promptModules()
// 	if err != nil {
// 		print(errorText("Module prompt failed: " + err.Error()))
// 		return
// 	}

// 	// If you exit the prompt, cancel
// 	if quitting || checkForQuit(result) {
// 		return
// 	}

// 	// Get the selected module
// 	for _, module := range catalog {
// 		if strings.Contains(result, module.Name) {
// 			selectedModule = module
// 		}
// 	}

// 	// Confirm selected module
// 	if selectedModule == (Module{}) {
// 		print(errorText("Error getting module..."))
// 		return
// 	}

// 	// Create the destination directory path inside modules
// 	destination := filepath.Join(".", "modules", selectedModule.Name)

// 	// Create the directory if it doesn't exist
// 	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
// 		print(errorText(fmt.Sprintf("Error creating directory for module: %v", err)))
// 		return
// 	}

// 	// Clone the module using the cloneRepo function
// 	if err := cloneRepo(selectedModule.CloneURL, destination); err != nil {
// 		print(errorText(fmt.Sprintf("Error cloning repository: %v", err)))
// 		return
// 	}

// 	fmt.Println(infoTitle("Module successfully cloned to " + destination))

// 	// Call startFileLoader to display the progress bar while processing the file
// 	err = startFileLoader(selectedModule, "")
// 	if err != nil {
// 		print(errorText("Error loading file: " + err.Error()))
// 		return
// 	}

// 	// Path to the .tf file within modules/templates
// 	tfFilePath := filepath.Join(destination, "files", "templates", "module.example.tf")

// 	// Check if the .tf file inside modules exists
// 	if _, err := os.Stat(tfFilePath); os.IsNotExist(err) {
// 		print(errorText("TF file does not exist: " + tfFilePath))
// 		return
// 	} else if err != nil {
// 		print(errorText("Error checking TF file: " + err.Error()))
// 		return
// 	}

// 	// Read the content of the template file inside modules/templates
// 	content, err := os.ReadFile(tfFilePath)
// 	if err != nil {
// 		print(errorText("Error reading template TF file: " + err.Error()))
// 		return
// 	}

// 	// Split the content into lines to modify the `source` line
// 	lines := strings.Split(string(content), "\n")
// 	var newContent []string

// 	// Modify the line containing `source =`
// 	for _, line := range lines {
// 		if strings.HasPrefix(strings.TrimSpace(line), "source =") {
// 			newSourceLine := fmt.Sprintf("source = \"%s\"", destination)
// 			newContent = append(newContent, newSourceLine)
// 		} else {
// 			newContent = append(newContent, line)
// 		}
// 	}

// 	// Join the modified content back into a string
// 	finalContent := strings.Join(newContent, "\n")

// 	// Overwrite the file in modules/templates with the updated content
// 	err = os.WriteFile(tfFilePath, []byte(finalContent), 0644)
// 	if err != nil {
// 		print(errorText("Error overwriting TF file in templates: " + err.Error()))
// 		return
// 	}

// 	fmt.Println(infoText("File in modules/templates successfully updated with new source value."))

// 	// Also write the updated content to the file outside modules
// 	cloneOutsidePath := filepath.Join(".", fmt.Sprintf("module.%s.tf", selectedModule.Name))
// 	err = os.WriteFile(cloneOutsidePath, []byte(finalContent), 0644)
// 	if err != nil {
// 		print(errorText("Error writing TF file outside modules: " + err.Error()))
// 		return
// 	}

// 	fmt.Println(infoText("File " + cloneOutsidePath + " created and updated with new source value!"))
// }

// Clear and position the output at the top
func enterAltScreen() {
	fmt.Print("\033[H\033[2J")
}
