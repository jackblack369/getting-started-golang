package main

import "testing"

type MockSender struct {
	messages []string
}

func (m *MockSender) Send(message string) error {
	m.messages = append(m.messages, message)
	return nil
}

func TestNotificationService(t *testing.T) {
	mockSender := &MockSender{}
	notifier := NewNotificationService(mockSender)

	notifier.SendNotification("Test message")

	if len(mockSender.messages) != 1 || mockSender.messages[0] != "Test message" {
		t.Error("Notification was not sent correctly")
	}
}
