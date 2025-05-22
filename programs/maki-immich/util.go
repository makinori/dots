package main

import (
	"time"
)

func retryNoFail[T any](
	output *T,
	attempts int, waitDuration time.Duration,
	fn func() (T, error),
	// errorMessagePrefix string,
) bool {
	var currentOutput T
	var lastError error

	attempt := 0

	for attempt < attempts {
		currentOutput, lastError = fn()

		if lastError == nil {
			// success
			*output = currentOutput
			return true
		}

		// final := attempt == attempts-1

		// if final {
		// 	log.Errorf("%s: %s", errorMessagePrefix, lastError.Error())
		// } else {
		// 	log.Warnf(
		// 		"%s (attempt %d): %s",
		// 		errorMessagePrefix, attempt+1, lastError.Error(),
		// 	)
		// }

		attempt++
		time.Sleep(waitDuration)
	}

	return false
}

func retryNoFailNoOutput(
	attempts int, waitDuration time.Duration,
	fn func() error,
	// errorMessagePrefix string,
) bool {
	var discard struct{}
	return retryNoFail(
		&discard, attempts, waitDuration,
		func() (struct{}, error) {
			err := fn()
			return discard, err
		},
		// errorMessagePrefix,
	)
}
