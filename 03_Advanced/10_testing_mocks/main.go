package main

import "fmt"

type MessageSender interface {
	Send(msg string) error
}
type NotificationService struct {
	Sender MessageSender
}

func (n *NotificationService) Alert(msg string) error {
	return n.Sender.Send(msg)
}

func main() {
	fmt.Println("Testing and Mocking")
}
