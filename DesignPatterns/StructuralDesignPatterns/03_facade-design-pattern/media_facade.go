package main

type MediaFacade struct {
	musicPlayer *MusicPlayer
	videoPlayer *VideoPlayer
	ImageViewer *ImageViewer
}

func NewMediaFacade() *MediaFacade {
	return &MediaFacade{
		musicPlayer: &MusicPlayer{},
		videoPlayer: &VideoPlayer{},
		ImageViewer: &ImageViewer{},
	}
}

func (f *MediaFacade) PerformAction(action string) {
	switch action {
	case "play music":
		f.musicPlayer.InitializeAudioDrivers()
		f.musicPlayer.DecodeAudion()
		f.musicPlayer.PlayAudio()
	case "play video":
		f.videoPlayer.LoadVideoFiles()
		f.videoPlayer.SetupRenderingEnginer()
		f.videoPlayer.PlayVideo()
	case "view image":
		f.ImageViewer.LoadImage("image.jpg")
		f.ImageViewer.DisplayImage()

	default:
		println("Invalid action")
	}
}
