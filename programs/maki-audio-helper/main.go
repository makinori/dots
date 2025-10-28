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
	MIC_VOLUME = 0.6

	// bypass easy effects (framework 13)
	EASY_EFFECTS_ONLY_NAME = "Family 17h/19h/1ah HD Audio Controller"
)

var (
	client *pulse.Client
	// muteMic bool
)

func clamp(val float64, min float64, max float64) float64 {
	return math.Max(min, math.Min(val, max))
}

func fixMicVolume() error {
	sources, err := client.ListSources()
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

	volume := MIC_VOLUME
	// if muteMic {
	// 	volume = 0
	// }

	volumeUint := uint32(
		clamp(volume, 0, 1) * float64(0xffff),
	)

	channelVolumes := make([]uint32, len(foundSource.Channels()))
	for i := range channelVolumes {
		channelVolumes[i] = volumeUint
	}

	client.RawRequest(&proto.SetSourceVolume{
		SourceIndex:    foundSource.SourceIndex(),
		ChannelVolumes: channelVolumes,
	}, nil)

	return nil
}

func loop() {
	if client == nil {
		var err error
		client, err = pulse.NewClient()
		if err != nil {
			log.Println("failed to create pulse client: " + err.Error())
			return
		}
	}

	err := fixMicVolume()
	if err != nil {
		log.Println("failed to fix mic volume: " + err.Error())
		client.Close()
		client = nil
	}

	err = checkEasyEffects()
	if err != nil {
		log.Println("failed to check easy effects: " + err.Error())
		client.Close()
		client = nil
	}
}

// type dbusInterface struct {}

// func (export *dbusInterface) MuteMic(message dbus.Message) (string, error) {
// 	muteMic = !muteMic
// 	if muteMic {
// 		log.Println("muted mic")
// 	} else {
// 		log.Println("unmuted mic")
// 	}
// 	loop() // run early
// 	return "pass", nil
// }

// just use XF86AudioMicMute which shows a dialog on gnome
// mute is also different from adjusting volume, so above code is respected

func main() {
	// conn, err := dbus.ConnectSessionBus()
	// if err != nil {
	// 	log.Println("failed to connect to session bus:", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close()

	// err = conn.ExportAll(
	// 	&dbusInterface{},
	// 	"/cafe/maki/AudioHelper", "cafe.maki.AudioHelper",
	// )
	// if err != nil {
	// 	log.Println("failed to export:", err)
	// 	os.Exit(1)
	// }

	// _, err = conn.RequestName("cafe.maki.AudioHelper", dbus.NameFlagDoNotQueue)
	// if err != nil {
	// 	log.Println("failed to request name:", err)
	// 	os.Exit(1)
	// }

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	for {
		loop()
		time.Sleep(time.Millisecond * LOOP_INTERVAL)
	}
}
