package cors

type listener func(*Event, *EventDispatcher)

type EventDispatcher struct {
  listeners map[string][]listener
}

func NewEventDispatcher() *EventDispatcher {
  ed := new(EventDispatcher)
  ed.listeners = map[string][]listener{}

  return ed
}

func (ed *EventDispatcher) dispatch(e *Event, eventName string) {
  for _, l := range ed.listeners[eventName] {
    if e.isPropagationStopped() {
      break
    }

    l(e, ed)
  }
}

func (ed *EventDispatcher) addListener(eventName string, l listener) {
  ed.listeners[eventName] = append(ed.listeners[eventName], l)
}
