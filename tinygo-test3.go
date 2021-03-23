package main

import (
	"machine"
	"sync"
	"time"
)

/*
  CONCURRENT LED
*/

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			led.Low()
			<-time.Tick(time.Millisecond * 250)

			led.High()
			<-time.Tick(time.Millisecond * 5000)

			led.Low()
			<-time.Tick(time.Millisecond * 250)
		}
	}()

	go func() {
		defer wg.Done()
		for j := 0; j < 100; j++ {

			led.Low()
			<-time.Tick(time.Millisecond * 250)

			led.High()
			<-time.Tick(time.Millisecond * 200)

			led.Low()
			<-time.Tick(time.Millisecond * 100)

			led.High()
			<-time.Tick(time.Millisecond * 200)

			led.Low()
			<-time.Tick(time.Millisecond * 250)
		}
	}()
	wg.Wait()
}
