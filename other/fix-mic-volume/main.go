package main

import (
	"errors"
	"log"
	"math"
	"strings"
	"time"

	"github.com/jfreymuth/pulse"
	"github.com/jfreymuth/pulse/proto"
)

const (
	MIC_NAME      = "Yeti Nano"
	MIC_VOLUME    = 0.6
	INTERVAL_TIME = 500 // ms
)

var (
	ErrClientNotInitalized = errors.New("client not initialized")
	ErrFailedToGetSources  = errors.New("failed to get sources")
	ErrFailedToFindMic     = errors.New("failed to find " + MIC_NAME)
)

func clamp(val float64, min float64, max float64) float64 {
	return math.Max(min, math.Min(val, max))
}

type State struct {
	Client *pulse.Client
}

func updateMicVolume(state *State) error {
	if state.Client == nil {
		return ErrClientNotInitalized
	}

	sources, err := state.Client.ListSources()
	if err != nil {
		return ErrFailedToGetSources
	}

	var foundSource *pulse.Source

	for _, source := range sources {
		name := source.Name()

		if strings.Contains(name, MIC_NAME) &&
			!strings.Contains(name, "Monitor of") {

			foundSource = source
		}
	}

	if foundSource == nil {
		return ErrFailedToFindMic
	}

	volumeUint := uint32(
		clamp(MIC_VOLUME, 0, 1) * float64(0xffff),
	)

	channelVolumes := make([]uint32, len(foundSource.Channels()))
	for i := range channelVolumes {
		channelVolumes[i] = volumeUint
	}

	state.Client.RawRequest(&proto.SetSourceVolume{
		SourceIndex:    foundSource.SourceIndex(),
		ChannelVolumes: channelVolumes,
	}, nil)

	return nil
}

func initPulse(state *State) bool {
	if state.Client == nil {
		client, err := pulse.NewClient()
		state.Client = client

		if err != nil {
			log.Println("failed to connect to pulseaudio")
			return false
		}

		log.Println("connected to pulseaudio")
	}

	return true
}

func loop(state *State) {
	if !initPulse(state) {
		return
	}

	err := updateMicVolume(state)

	if err == nil {
		return
	}

	log.Println(err)

	switch err {
	case ErrClientNotInitalized:
	case ErrFailedToGetSources:
		state.Client.Close()
		state.Client = nil

		if !initPulse(state) {
			return
		}
	}

}

func main() {
	var state State

	cleanup := func() {
		if state.Client != nil {
			state.Client.Close()
		}
	}

	defer cleanup()

	for {
		loop(&state)
		time.Sleep(time.Millisecond * INTERVAL_TIME)
	}
}
