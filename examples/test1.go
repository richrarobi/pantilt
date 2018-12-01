package main

import (
    "fmt"
    "time"
    pt "github.com/richrarobi/pantilt"
)

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
    pt.Open()

// test invalid name
//    fmt.Println(pt.PtServoEnable("fred", true))
    pt.Go(0,0)
    fmt.Println("getServo (pan): ", pt.GetServo("pan"))
    fmt.Println("getServo (tilt): ", pt.GetServo("tilt"))

// camera is inverted i.e. -ve tilt is up
    pt.ServoEnable("pan", true)
    pt.Delta("pan", 45)
// disable (test)
//    pt.ServoEnable("pan", false)

    pt.ServoEnable("tilt", true)
    pt.Delta("tilt", 45)
    fmt.Println("getServo (pan): ", pt.GetServo("pan"))
    fmt.Println("getServo (tilt): ", pt.GetServo("tilt"))
//    pt.ServoEnable("tilt", false)

//    pt.ServoEnable("pan", true)
    pt.Delta("pan", -90)
//    pt.ServoEnable("tilt", true)
    pt.Delta("tilt", -90)
    fmt.Println("getServo (pan): ", pt.GetServo("pan"))
    fmt.Println("getServo (tilt): ", pt.GetServo("tilt"))
    
    pt.Delta("pan", 45)
    pt.Delta("tilt", 45)
// test invalid name
//    fmt.Println(pt.PtDelta("dave", 20))
//    delay(2000)
    pt.Go(0,0)
//    delay(2000)
    pt.ServoEnable("pan", false)
    pt.ServoEnable("tilt", false)
    fmt.Println("getServo (pan): ", pt.GetServo("pan"))
    fmt.Println("getServo (tilt): ", pt.GetServo("tilt"))
// close the i2c bus
    pt.Close()

// test after close
//    fmt.Println("getServo (pan): ", pt.GetServo("pan"))

}
