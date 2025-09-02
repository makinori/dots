//go:build laptop

package main

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

func setEasyEffectsBypass(bypass bool) error {
	log.Printf("setting easy effects bypass to %v\n", bypass)

	bypassArg := ""
	if bypass {
		bypassArg = "1"
	} else {
		bypassArg = "2"
	}

	return exec.Command("easyeffects", "-b", bypassArg).Run()
}

func (state *State) checkEasyEffects() error {
	defaultSink, err := state.client.DefaultSink()
	if err != nil {
		return errors.New("failed to get default sink: " + err.Error())
	}

	defaultSinkName := defaultSink.Name()
	defaultSinkID := defaultSink.ID()

	bypass := !strings.Contains(defaultSinkName, EASY_EFFECTS_ONLY_NAME) ||
		strings.Contains(defaultSinkName, "Monitor of")

	if state.lastDefaultSinkID != defaultSinkID {
		setEasyEffectsBypass(bypass)
		state.lastDefaultSinkID = defaultSinkID
	}

	return nil
}
