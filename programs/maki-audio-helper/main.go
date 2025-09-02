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
	LOOP_INTERVAL = 500 // ms

	// force mic volume
	MIC_NAME   = "Yeti Nano"
	MIC_VOLUME = 0.7

	// bypass easy effects (framework 13)
	EASY_EFFECTS_ONLY_NAME = "Family 17h/19h/1ah HD Audio Controller"
)

type State struct {
	client            *pulse.Client
	lastDefaultSinkID string
}

func clamp(val float64, min float64, max float64) float64 {
	return math.Max(min, math.Min(val, max))
}

func (state *State) fixMicVolume() error {
	sources, err := state.client.ListSources()
	if err != nil {
		return errors.New("failed to get sources: " + err.Error())
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
		return nil
	}

	volumeUint := uint32(
		clamp(MIC_VOLUME, 0, 1) * float64(0xffff),
	)

	channelVolumes := make([]uint32, len(foundSource.Channels()))
	for i := range channelVolumes {
		channelVolumes[i] = volumeUint
	}

	state.client.RawRequest(&proto.SetSourceVolume{
		SourceIndex:    foundSource.SourceIndex(),
		ChannelVolumes: channelVolumes,
	}, nil)

	return nil
}

func loop(state *State) {
	if state.client == nil {
		var err error
		state.client, err = pulse.NewClient()
		if err != nil {
			log.Println("failed to create pulse client: " + err.Error())
			return
		}
	}

	err := state.fixMicVolume()
	if err != nil {
		log.Println("failed to fix mic volume: " + err.Error())
		state.client.Close()
		state.client = nil
	}

	err = state.checkEasyEffects()
	if err != nil {
		log.Println("failed to check easy effects: " + err.Error())
		state.client.Close()
		state.client = nil
	}
}

func main() {
	var state State

	defer func() {
		if state.client != nil {
			state.client.Close()
		}
	}()

	for {
		loop(&state)
		time.Sleep(time.Millisecond * LOOP_INTERVAL)
	}
}
