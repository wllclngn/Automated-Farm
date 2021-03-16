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
        if (num % 2) == 0 {
            led.Low()
            time.Sleep(time.Millisecond * 250)
        
            led.High()
            time.Sleep(time.Millisecond * 275)

            led.Low()
            time.Sleep(time.Millisecond * 100)
        
            led.High()
            time.Sleep(time.Millisecond * 275)

            led.Low()
            time.Sleep(time.Millisecond * 250)
            if (num == 100) {
                num = 0
            }
        } else {
            led.Low()
            time.Sleep(time.Millisecond * 250)
        
            led.High()
            time.Sleep(time.Millisecond * 500)

            led.Low()
            time.Sleep(time.Millisecond * 250)
        
        }
        num++
    }
}
