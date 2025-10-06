package main

func addEmailsToQueue(emails []string) chan string {
	// ?

	length := len(emails)
	emailChan := make(chan string, length)

	for _, email := range emails {
		emailChan <- email
	}

	return emailChan
}
