package cors

type listener func(*event, *eventDispatcher)

type eventDispatcher struct {
  listeners map[string][]listener
}

func (ed *eventDispatcher) dispatch(e *event, eventName string) {
  for _, l := range ed.listeners[eventName] {
    if e.isPropagationStopped() {
      break
    }

    l(e, ed)
  }
}

func (ed *eventDispatcher) addListener(eventName string, l listener) {
  ed.listeners[eventName] = append(ed.listeners[eventName], l)
}
