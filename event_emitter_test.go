package event_emitter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmitterShouldAddListeners(t *testing.T) {
	emitter := NewEventEmitter[string](10)
	callCounter := 0
	cb := func(event string) {
		callCounter += 1
	}

	emitter.On("event", cb)

	emitter.Emit("event", "Event emitted")

	assert.Equal(t, callCounter, 1, "Callback should be called 1 time")
}

func TestEmitterShouldRemoveListeners(t *testing.T) {
	emitter := NewEventEmitter[string](10)
	callCounter := 0
	cb := func(event string) {
		fmt.Println(event)
	}

	emitter.On("event", cb)
	emitter.Off("event", cb)

	emitter.Emit("event", "Event emitted")

	assert.Equal(t, callCounter, 0, "Callback should be called 0 times")
}
