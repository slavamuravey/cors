package cors

import (
  "testing"
)

func TestDispatch(t *testing.T)  {
  ed := newEventDispatcher()
  var listenerInvoked bool
  ed.addListener("event", func(e *event, ed *eventDispatcher) {
    listenerInvoked = true
  })

  ed.dispatch(newEvent(nil, nil, nil), "event")

  assertTrue(t, listenerInvoked, "listener must be invoked")
}

func TestDispatchStopPropagation(t *testing.T)  {
  ed := newEventDispatcher()
  var listener1Invoked, listener2Invoked, listener3Invoked bool

  ed.addListener("event", func(e *event, ed *eventDispatcher) {
    listener1Invoked = true
  })

  ed.addListener("event", func(e *event, ed *eventDispatcher) {
    e.stopPropagation()
    listener2Invoked = true
  })

  ed.addListener("event", func(e *event, ed *eventDispatcher) {
    listener3Invoked = true
  })

  ed.dispatch(newEvent(nil, nil, nil), "event")

  assertTrue(t, listener1Invoked, "listener must be invoked")
  assertTrue(t, listener2Invoked, "listener must be invoked")
  assertFalse(t, listener3Invoked, "listener mustn't be invoked")
}
