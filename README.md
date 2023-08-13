# golang-event-emitter
Event emitter to use in golang


## Usage

### Create new event listener

```golang
import "github.com/tgilauri/golang-event-emitter"

func main() {
    // Provide a type for event data and max amount of listeners
    eventListener := NewEventListener[string](10)

    // To make EventListener for unknown event data type use <any>
    eventListener := NewEventListener[any](10)
}
```

### Add event listener

To add event listener for specific event

```golang
eventListener.On("eventName", callback)
```

### Remove event listener

To add event listener for specific event

```golang
eventListener.Off("eventName", callback)
```

### Trigger event

To add event listener for specific event

```golang
eventListener.On("eventName", callback)

eventListener.Emit("eventName", data) // callback will be called on this event
```