package main

type YouTubeChannelImpl struct {
	channelName string
	subscribers []Subscriber
	video       string
}

func NewYouTubeChannelImpl(channelName string) *YouTubeChannelImpl {
	return &YouTubeChannelImpl{
		channelName: channelName,
		subscribers: make([]Subscriber, 0),
	}
}

func (c *YouTubeChannelImpl) AddSubscriber(s Subscriber) {
	c.subscribers = append(c.subscribers, s)
}

func (c *YouTubeChannelImpl) RemoveSubscriber(subscriber Subscriber) {
	for i, s := range c.subscribers {
		if s == subscriber {
			c.subscribers = append(c.subscribers[:i], c.subscribers[i+1:]...)
		}
	}
}

func (c *YouTubeChannelImpl) NotifySubscribers() {
	for _, subscriber := range c.subscribers {
		subscriber.Update(c.video)
	}
}

func (c *YouTubeChannelImpl) UploadVideo(video string) {
	c.video = video
	c.NotifySubscribers()
}
