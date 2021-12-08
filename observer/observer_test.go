package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(&ObserverOne{})
	sub.Register(&ObserverTwo{})
	sub.Notify("hi")
}
