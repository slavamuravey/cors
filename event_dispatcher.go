package cors

type Listener func(*Event, *EventDispatcher)

type EventDispatcher struct {
  listeners map[string][]Listener
}

func NewEventDispatcher() *EventDispatcher {
  ed := new(EventDispatcher)
  ed.listeners = map[string][]Listener{}

  return ed
}

func (ed *EventDispatcher) dispatch(e *Event, eventName string) {
  for _, l := range ed.listeners[eventName] {
    if e.IsPropagationStopped() {
      break
    }

    l(e, ed)
  }
}

func (ed *EventDispatcher) addListener(eventName string, l Listener) {
  ed.listeners[eventName] = append(ed.listeners[eventName], l)
}
