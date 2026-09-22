package main

import "fmt"

type SMSSubscriber struct {
	phoneNumber string
}

func NewSMSSubscriber(phoneNumber string) *SMSSubscriber {
	return &SMSSubscriber{phoneNumber: phoneNumber}
}

func (s *SMSSubscriber) Update(video string) {
	fmt.Printf("sending sms to %s for the video %s", s.phoneNumber, video)
}
