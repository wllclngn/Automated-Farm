package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    num := 0
    for {
        num++
        if (num % 2) == 0 {
            led.Low()
            time.Sleep(time.Millisecond * 250)
        
            led.High()
            time.Sleep(time.Millisecond * 250)

            led.Low()
            time.Sleep(time.Millisecond * 250)
        
            led.High()
            time.Sleep(time.Millisecond * 250)
        } else {
            led.Low()
            time.Sleep(time.Millisecond * 500)
        
            led.High()
            time.Sleep(time.Millisecond * 500)
        }
    }
}
