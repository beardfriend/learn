package main

import (
	"fmt"
	"sync"
	"time"
)


type Car struct {
	Body string
	Tire string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()
func main () {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh,paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the Factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After((10 * time.Second))

	for {
		select {
		case <- tick:
			//Make a Body
			car := Car{}
			car.Body = "Sports car"
			tireCh <- &car

		case <- after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan * Car) {
	for car:= range tireCh {
			// Make a body
		time.Sleep(time.Second)
		car.Tire ="Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintcCh chan *Car) { //도색
	for car:= range paintcCh {
		//Make a body
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body,car.Tire,car.Color)
	}
	wg.Done()
}