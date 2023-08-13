package event_emitter

import (
	"fmt"
	"reflect"
)

type Callback[T any] func(T)

type EventEmitter[T any] interface {
	On(string, Callback[T])
	Off(string, Callback[T])
	Emit(string, T)
	ClearAllListeners()
}

type SEventEmitter[T any] struct {
	listeners    map[string]map[uintptr]Callback[T]
	maxListeners int
}

func NewEventEmitter[T any](maxListeners int) EventEmitter[T] {
	emitter := new(SEventEmitter[T])
	emitter.maxListeners = 10
	emitter.listeners = make(map[string]map[uintptr]Callback[T])

	return emitter
}

func (this *SEventEmitter[T]) On(eventName string, callback Callback[T]) {
	ptr := reflect.ValueOf(callback).Pointer()

	fmt.Printf("ON. Pointer to callback is %d\n", ptr)

	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		this.listeners[eventName] = make(map[uintptr]Callback[T])
	}
	if len(this.listeners[eventName]) >= this.maxListeners {
		panic("Can't add more listeners")
	}

	this.listeners[eventName][ptr] = callback
}

func (this *SEventEmitter[T]) Off(eventName string, callback Callback[T]) {
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

func (this *SEventEmitter[T]) Emit(eventName string, data T) {
	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		return
	}

	for _, callback := range this.listeners[eventName] {
		callback(data)
	}
}

func (this *SEventEmitter[T]) ClearAllListeners() {
	if len(this.listeners) == 0 {
		return
	}

	this.listeners = make(map[string]map[uintptr]Callback[T])
}
