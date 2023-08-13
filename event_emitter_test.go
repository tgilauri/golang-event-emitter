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

	assert.Equal(t, 1, callCounter, "Callback should be called 1 time")
}

func TestEmitterShouldRemoveListeners(t *testing.T) {
	emitter := NewEventEmitter[string](10)
	callCounter := 0
	cb := func(event string) {
		fmt.Println(event)
		callCounter += 1
	}

	emitter.On("event", cb)
	emitter.Off("event", cb)

	emitter.Emit("event", "Event emitted")

	assert.Equal(t, 0, callCounter, "Callback should be called 0 times")
}

func TestEmitterShouldRunAllListeners(t *testing.T) {
	emitter := NewEventEmitter[string](10)
	callCounter := 0
	cb := func(event string) {
		fmt.Println(event)
		callCounter += 1
	}

	cb2 := func(event string) {
		fmt.Println(event)
		callCounter += 1
	}

	emitter.On("event", cb)
	emitter.On("event", cb2)

	emitter.Emit("event", "Event emitted")

	assert.Equal(t, 2, callCounter, "Callback should be called 2 times")
}

func TestEmitterShouldRunListenersWithDifferentEventTypes(t *testing.T) {
	emitter := NewEventEmitter[any](10)
	callCounter := 0
	cb := func(event any) {
		fmt.Println(event)
		callCounter += 1
	}

	cb2 := func(event any) {
		fmt.Println(event)
		callCounter += 1
	}

	emitter.On("event", cb)
	emitter.On("someOther", cb2)

	emitter.Emit("event", "Event emitted")
	emitter.Emit("someOther", 1)

	assert.Equal(t, 2, callCounter, "Callback should be called 2 times")
}
