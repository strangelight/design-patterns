package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var singleton *Singleton

type Singleton struct {
	name string
}

func (s *Singleton) SetName(name string) {
	s.name = name
}

func getSingletonInstance() *Singleton {
	//thread safety: http://marcio.io/2015/07/singleton-pattern-in-go/
	once.Do(func() {
		singleton = new(Singleton)
	})

	return singleton
}

func (s Singleton) Hello() {
	fmt.Printf("Hello %v\n", s.name)
}

func main() {
	s1 := getSingletonInstance()
	s1.SetName("Maicon")
	s1.Hello()
	//Maicon

	s2 := getSingletonInstance()
	s2.SetName("Patty")
	s2.Hello()
	//Patty

	s1.Hello()
	//Patty (again!)
}
