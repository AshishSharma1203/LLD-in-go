package main

import "fmt"

func main() {
	MediaFacade := NewMediaFacade()

	fmt.Println("Welcome to the Media Player!")

	fmt.Println("chose an option: play video or view image or play music")

	for _, action := range []string{"play music", "play video", "view image"} {
		fmt.Println("Performing action:", action)
		MediaFacade.PerformAction(action)
	}
}
