package graphmailer

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// ParseAddresses reads a log file and extracts email addresses
func ParseAddresses(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var emails []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`[\w\.-]+@[\w\.-]+\.\w+`)
		email := re.FindString(line)

		if email != "" {
			emails = append(emails, email)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return emails, nil
}
