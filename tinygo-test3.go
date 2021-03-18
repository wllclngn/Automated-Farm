package main

import (
    "machine"
    "time"
    "sync"
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
        for i := 0; i < 10; i++ {
            led.Low()
            time.Sleep(time.Millisecond * 500)

            led.High()
            time.Sleep(time.Millisecond * 500)

            led.Low()
            time.Sleep(time.Millisecond * 500)
        }
    }()

    go func() {
        defer wg.Done()
        for j := 0; j < 10; j++ {

            led.Low()
            time.Sleep(time.Millisecond * 100)

            led.High()
            time.Sleep(time.Millisecond * 200)

            led.Low()
            time.Sleep(time.Millisecond * 100)

            led.High()
            time.Sleep(time.Millisecond * 200)

            led.Low()
            time.Sleep(time.Millisecond * 100)

            led.High()
            time.Sleep(time.Millisecond * 200)
        }
    }()
    wg.Wait()
}
