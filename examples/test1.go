package main

import (
    "os"
    "fmt"
    "time"
    "os/signal"
    "syscall"
    pt "github.com/richrarobi/pantilt"
)

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
    running := true
// initialise getout
    signalChannel := make(chan os.Signal, 2)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
    go func() {
        sig := <-signalChannel
        switch sig {
        case os.Interrupt:
//            fmt.Println("Stopping on Interrupt")
            running = false
            return
        case syscall.SIGTERM:
//            fmt.Println("Stopping on Terminate")
            running = false
            return
        }
    }()
    
    pt.PtOpen()
    pt.PtServoEnable("pan", true)
    pt.PtServoEnable("tilt", true)
    fmt.Println(pt.PtServoEnable("fred", true))
    pt.PtHome()
//    delay(2000)
// camera is inverted
// e.g -ve tilt is up
    pt.PtDelta("pan", 20)
    pt.PtDelta("tilt", 20)
    pt.PtDelta("pan", -40)
    pt.PtDelta("tilt", -40)
    pt.PtDelta("pan", 20)
    pt.PtDelta("tilt", 20)
    fmt.Println(pt.PtDelta("dave", 20))
//    delay(2000)
    pt.PtHome()
//    delay(2000)
    pt.PtServoEnable("pan", false)
    pt.PtServoEnable("tilt", false)
// close the i2c bus
    pt.PtClose()

}
