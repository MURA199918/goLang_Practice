package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var host = make(chan bool, 2)

//ChopS a single chopstick can be used as a mutex object
type ChopS struct{ sync.Mutex }

//Philo philosopher has an ID, and 2 assigned chopsticks
type Philo struct {
	id      int
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) eat() {

	for i := 0; i < 3; i++ {

		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat: ", p.id+1)
		time.Sleep(2 * time.Second)
		fmt.Println("finishing eating: ", p.id+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<-host

	}
	wg.Done()
}

func main() {

	//create the chopsticks
	ChopSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)
	}

	// create the philosophers
	Philosophers := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philo{i, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	// let the philosophers eat now
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philosophers[i].eat()
	}

	wg.Wait()

}
