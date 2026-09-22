package main 

type YouTubeChannel interface{
	AddSubscriber(s Subscriber)
	RemoveSubscriber(s Subscriber)
	NotifySubscribers()
}