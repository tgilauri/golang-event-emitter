package event_emitter

import (
	"fmt"
	"reflect"
)

type Callback[EventDataType any] func(EventDataType)

type EventEmitter[EventDataType any] interface {
	On(string, Callback[EventDataType])
	Off(string, Callback[EventDataType])
	Emit(string, EventDataType)
	ClearAllListeners()
}

type SEventEmitter[EventDataType any] struct {
	listeners    map[string]map[uintptr]Callback[EventDataType]
	maxListeners int
}

func NewEventEmitter[EventDataType any](maxListeners int) EventEmitter[EventDataType] {
	emitter := new(SEventEmitter[EventDataType])
	emitter.maxListeners = maxListeners
	emitter.listeners = make(map[string]map[uintptr]Callback[EventDataType])

	return emitter
}

func (this *SEventEmitter[EventDataType]) On(eventName string, callback Callback[EventDataType]) {
	ptr := reflect.ValueOf(callback).Pointer()

	fmt.Printf("ON. Pointer to callback is %d\n", ptr)

	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		this.listeners[eventName] = make(map[uintptr]Callback[EventDataType])
	}
	if len(this.listeners[eventName]) >= this.maxListeners {
		panic("Can't add more listeners")
	}

	this.listeners[eventName][ptr] = callback
}

func (this *SEventEmitter[EventDataType]) Off(eventName string, callback Callback[EventDataType]) {
	ptr := reflect.ValueOf(callback).Pointer()

	fmt.Printf("OFF. Pointer to callback is %d\n", ptr)

	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		return
	}
	fmt.Println(this.listeners[eventName][ptr])
	if _, ok := this.listeners[eventName][ptr]; !ok {
		return
	}

	delete(this.listeners[eventName], ptr)

}

func (this *SEventEmitter[EventDataType]) Emit(eventName string, data EventDataType) {
	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		return
	}

	for _, callback := range this.listeners[eventName] {
		callback(data)
	}
}

func (this *SEventEmitter[EventDataType]) ClearAllListeners() {
	if len(this.listeners) == 0 {
		return
	}

	this.listeners = make(map[string]map[uintptr]Callback[EventDataType])
}
