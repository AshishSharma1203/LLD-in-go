package main

type MusicPlayer struct{}

func (m *MusicPlayer) InitializeAudioDrivers() {
	println("audio drivers initializing...")
}

func (m *MusicPlayer) DecodeAudion() {
	println("decoding audio...")
}

func (m *MusicPlayer) PlayAudio() {
	println("playing audio...")
}
