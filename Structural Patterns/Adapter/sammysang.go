package main

import "fmt"

type SammysangTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *SammysangTV) getVolume() int {
	fmt.Println("SammysangTV volume is ", tv.currentVolume)
	return tv.currentVolume
}

func (tv *SammysangTV) setVolume(vol int) {
	fmt.Println("Setting SammysangTV volume to ", vol)
	tv.currentVolume = vol
}
func (tv *SammysangTV) getChannel() int {
	fmt.Println("SammysangTV Channel is ", tv.currentChan)
	return tv.currentChan
}
func (tv *SammysangTV) setChannel(channel int) {
	fmt.Println("Setting SammysangTV channel to ", channel)
	tv.currentChan = channel
}
func (tv *SammysangTV) setOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("SammysangTv is on")
	} else {
		fmt.Println("SammysangTv is off")
	}
	tv.tvOn = tvOn
}
