package main

import "fmt"

type VideoPlayer struct{}

func (vp *VideoPlayer) PlayVideo() {
	// Logic to play video
	fmt.Println("Playing video...")
}

func (v *VideoPlayer) LoadVideoFiles() {
	fmt.Println("Loading video files...")

}

func (v *VideoPlayer) SetupRenderingEnginer() {
	fmt.Println("Setting up rendering engine...")
}
