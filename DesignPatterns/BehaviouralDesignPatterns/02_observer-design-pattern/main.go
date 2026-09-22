package main

func main() {
	sms1 := NewSMSSubscriber("82123435")
	sms2 := NewSMSSubscriber("2345245")

	email1 := NewEmailSubscriber("akr@gmail.com")
	email2 := NewEmailSubscriber("sharma@gmail.com")

	channel1 := NewYouTubeChannelImpl("AshishCodes")

	channel1.AddSubscriber(sms1)
	channel1.AddSubscriber(sms2)
	channel1.AddSubscriber(email1)
	channel1.AddSubscriber(email2)

	channel1.UploadVideo("Golang tutorial")
	channel1.RemoveSubscriber(sms1)
	channel1.UploadVideo("Desing pattern tutorial")

}
