package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Fullname    string
	Phonenumber string
}

func Vcf_parser(contacts_path string, path string) {
	file, err := os.Open(contacts_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var contacts []Contact
	var currentContact *Contact

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "BEGIN") {
			currentContact = &Contact{}
		} else if strings.HasPrefix(line, "FN:") {
			currentContact.Fullname = strings.TrimPrefix(line, "FN:")
		} else if strings.HasPrefix(line, "TEL;") {
			var line_ = strings.ReplaceAll(line, " ", "")
			parts := strings.Split(line_, ":")
			currentContact.Phonenumber = parts[len(parts)-1]
		} else if strings.HasPrefix(line, "END") {
			if currentContact.Phonenumber == "" {
				Red("❌  No full name assigned for number : " + currentContact.Fullname)
				continue
			}
			if currentContact.Fullname == "" {
				Red("❌  No number assigned for contact: " + currentContact.Phonenumber)
				continue
			}
			contacts = append(contacts, *currentContact)
			currentContact = nil
		}
	}

	for _, contact := range contacts {
		err := SetConfig("contacts", contact.Fullname, contact.Phonenumber, path)
		if err != nil {
			Red("❌  Fail to add the contact to the configuration.")
		}
	}

	if err := scanner.Err(); err != nil {
		Red("❌  Error reading file : " + err.Error())
	}
}
