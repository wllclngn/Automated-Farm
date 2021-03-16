package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    num := 0
    timing := func() {
        time.Sleep(time.Millisecond * 250)
    }
    for {
        if (num % 2) == 0 {
            led.Low()
            timing()
        
            led.High()
            timing()

            led.Low()
            timing()
        
            led.High()
            timing()

            led.Low()
            timing()
            if (num == 100) {
                num = 0
            }
        } else {
            led.Low()
            timing()
        
            led.High()
            timing()

            led.Low()
            timing()
        }
        num++
    }
}
