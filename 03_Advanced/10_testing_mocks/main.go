package main

import "fmt"

// 1. THE INTERFACE (The Rulebook)
// We don't care if the sender is Twilio, SendGrid, or a Fake Test object.
// We only care that it has a Send() method that takes a string.
type MessageSender interface {
	Send(msg string) error
}

// 2. THE SERVICE (The target we want to test)
type NotificationService struct {
	// 3. DEPENDENCY INJECTION
	// Instead of hardcoding `TwilioClient` here, we leave a blank socket `MessageSender`.
	// This allows us to plug in a Fake client when testing!
	Sender MessageSender
}

// 4. THE CORE LOGIC
// Notice we blindly call `n.Sender.Send`. We don't know what is actually inside `Sender`!
func (n *NotificationService) Alert(msg string) error {
	return n.Sender.Send(msg)
}

func main() {
	fmt.Println("Testing and Mocking")
}
