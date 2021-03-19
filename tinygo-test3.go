package main

import (
    "machine"
    "time"
    "sync"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})

    wg := &sync.WaitGroup{}
    wg.Add(2)

    go func() {
        defer wg.Done()
        for i := 0; i < 100; i++ {
            led.Low()
            time.Sleep(time.Millisecond * 250)

            led.High()
            time.Sleep(time.Millisecond * 500)

            led.Low()
            time.Sleep(time.Millisecond * 250)
        }
    }()

    go func() {
        defer wg.Done()
        for j := 0; j < 100; j++ {

            led.Low()
            time.Sleep(time.Millisecond * 250)

            led.High()
            time.Sleep(time.Millisecond * 200)

            led.Low()
            time.Sleep(time.Millisecond * 100)

            led.High()
            time.Sleep(time.Millisecond * 200)

            led.Low()
            time.Sleep(time.Millisecond * 250)
        }
    }()
    wg.Wait()
}
