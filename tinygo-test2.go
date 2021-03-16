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
            time.Sleep(time.Millisecond * 200)

            led.Low()
            time.Sleep(time.Millisecond * 100)
        
            led.High()
            time.Sleep(time.Millisecond * 200)

            led.Low()
            timing()
            if (num == 100) {
                num = 0
            }
        } else {
            led.Low()
            timing()
        
            led.High()
            time.Sleep(time.Millisecond * 500)

            led.Low()
            timing()
        }
        num++
    }
}
