package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?
func analyzeMessage(alytics *Analytics, msg Message) {
	switch msg.Success {
	case true:
		alytics.MessagesSucceeded += 1
	case false:
		alytics.MessagesFailed += 1
	}
	alytics.MessagesTotal += 1
}
