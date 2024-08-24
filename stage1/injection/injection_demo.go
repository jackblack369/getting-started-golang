package main

import (
	"fmt"
	"time"
)

type MessageSender interface {
	Send(message string) error
}

type EmailSender struct{}

func (e *EmailSender) Send(message string) error {
	// Implementation for sending email
	fmt.Println("Sending email:", message)
	return nil
}

type SMSSender struct{}

func (s *SMSSender) Send(message string) error {
	// Implementation for sending SMS
	fmt.Println("Sending SMS:", message)
	return nil
}

type NotificationService struct {
	sender MessageSender
}

func NewNotificationService(sender MessageSender) *NotificationService {
	return &NotificationService{
		sender: sender,
	}
}

func (ns *NotificationService) SendNotification(message string) error {
	return ns.sender.Send(message)
}

func main() {
	// calculate the cost time of sending email and SMS
	begin := time.Now()
	emailSender := &EmailSender{}
	smsSender := &SMSSender{}

	emailNotifier := NewNotificationService(emailSender)
	smsNotifier := NewNotificationService(smsSender)

	emailNotifier.SendNotification("Hello via email")
	smsNotifier.SendNotification("Hello via SMS")
	fmt.Println("Time taken:", time.Since(begin))
}
