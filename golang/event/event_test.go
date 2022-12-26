package event

import (
	"fmt"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	m := NewManger(16, 16)
	m.RegisterA("ceshi", "asdf1", func(event *Event) {
		fmt.Println("handler 1", event.Data)
	})
	m.RegisterA("ceshi", "asdf2", func(event *Event) {
		fmt.Println("handler 2", event.Data)
	})
	go func() {
		for i := 0; i < 100; i++ {
			m.CallA("ceshi", "a")
			time.Sleep(1 * time.Second)
		}
	}()
	for i := 0; i < 100; i++ {
		m.CallA("ceshi", i)
		time.Sleep(3 * time.Second)
	}
}
