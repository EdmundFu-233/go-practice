package main

import "fmt"

type Observer interface {
	Update(data string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(data string) {
	for _, o := range s.observers {
		o.Update(data)
	}
}

type EmailObserver struct{}

func (e EmailObserver) Update(data string) {
	fmt.Println("Email:", data)
}

type LogObserver struct{}

func (l LogObserver) Update(data string) {
	fmt.Println("Log:", data)
}

func main() {
	subject := &Subject{}
	subject.Attach(EmailObserver{})
	subject.Attach(LogObserver{})
	subject.Notify("Event occurred")
}
