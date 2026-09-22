package main

import "fmt"

type EmailSubscriber struct {
	email string
}

func NewEmailSubscriber(email string) *EmailSubscriber {
	return &EmailSubscriber{
		email: email,
	}
}

func (s *EmailSubscriber) Update(video string) {
	fmt.Printf("Sending email to %s for video %s", s.email, video)
}
