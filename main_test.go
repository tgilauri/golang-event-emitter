package main

import (
	"fmt"
	"testing"
)

func TestEmitterShouldAddListeners(t *testing.T) {
	emitter := NewEventEmitter[string](10)

	emitter.On("event", func(event string) {
		fmt.Println(event)
	})

	emitter.Emit("event", "Event emitted")
}
