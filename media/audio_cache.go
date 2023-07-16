package util

import (
	"sync"
)

// Minimal synchronization limited to cache access.
var (
	audioMapMutex = sync.RWMutex{}
)

type AudioCache interface {
	LoadAudioStreamer(file string)
	GetAudioStreamer(file string) AudioPlayer
}
type audioCacheData struct {
	audioCache map[string]AudioPlayer
}

func NewAudioCache() AudioCache {
	cache := audioCacheData{}
	cache.audioCache = make(map[string]AudioPlayer, 100)
	return &cache
}

func (c *audioCacheData) GetAudioStreamer(fileName string) AudioPlayer {
	//var streamer AudioPlayer
	//audioMapMutex.Lock()
	//streamer = c.audioCache[fileName]
	//audioMapMutex.Unlock()
	//return streamer
	return nil
}
func (c *audioCacheData) LoadAudioStreamer(fileName string) {
	//audioMapMutex.Lock()
	//if c.audioCache[fileName] == nil {
	//	f := GetFile(fileName)
	//	streamSeeker, format, err := mp3.Decode(f)
	//	player := audioPlayer{stream: streamSeeker, format: format, playing: true}
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer streamSeeker.Close()
	//
	//	c.audioCache[fileName] = &player
	//}
	//audioMapMutex.Unlock()
}
