package event_emitter

import (
	"fmt"
)

type Callback[T any] func(T)

type EventEmitter[T any] interface {
	On(string, Callback[T]) *Callback[T]
	Off(string, *Callback[T])
	Emit(string, T)
	ClearAllListeners()
}

type SEventEmitter[T any] struct {
	listeners    map[string]map[*Callback[T]]Callback[T]
	maxListeners int
}

func NewEventEmitter[T any](maxListeners int) EventEmitter[T] {
	emitter := new(SEventEmitter[T])
	emitter.maxListeners = 10
	emitter.listeners = make(map[string]map[*Callback[T]]Callback[T])

	return emitter
}

func (this *SEventEmitter[T]) On(eventName string, callback Callback[T]) *Callback[T] {
	fmt.Printf("ON. Pointer to callback is %d\n", &callback)
	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		this.listeners[eventName] = make(map[*Callback[T]]Callback[T])
	}
	if len(this.listeners[eventName]) >= this.maxListeners {
		panic("Can't add more listeners")
	}

	this.listeners[eventName][&callback] = callback
	return &callback
}

func (this *SEventEmitter[T]) Off(eventName string, callback *Callback[T]) {
	fmt.Printf("OFF. Pointer to callback is %d\n", callback)
	isEmpty := len(this.listeners) == 0
	if _, ok := this.listeners[eventName]; isEmpty || !ok {
		return
	}
	fmt.Println(this.listeners[eventName][callback])
	if _, ok := this.listeners[eventName][callback]; !ok {
		return
	}

	delete(this.listeners[eventName], callback)

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

	this.listeners = make(map[string]map[*Callback[T]]Callback[T])
}
