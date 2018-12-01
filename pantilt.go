package pantilt

import (
//    "reflect"
//    "fmt"
    "time"
//    "strconv"
    "encoding/binary"
    "log"
    "periph.io/x/periph/host"
//    "periph.io/x/periph/host/rpi"
    "periph.io/x/periph/conn/i2c/i2creg"
    "periph.io/x/periph/conn/i2c"
)

var dev i2c.Dev
var bus i2c.BusCloser
var confReg byte

const servoMin uint16 = 575
const servoMax uint16 = 2375
const degMin int = -90
const degMax int =  90
const addr uint16 = 0x15
const i2cbus string = "1"

type srvo struct {
    reg byte
    ang int
    bit byte
    abl bool
}

var srvos = map[string]*srvo {
"pan"  : &srvo{0x01, 0, 0, false},
"tilt" : &srvo{0x03, 0, 1, false},
}

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func valid(name string) (ok bool) {
    if _, ok := srvos[name]; ok {
        return ok
    }
    return false
}

func usToDeg(us uint16) (ang int) {
    if us < servoMin {us = servoMin}
    if us > servoMax {us = servoMax}
    rng := float32(servoMax - servoMin)
    pos := us - servoMin
    ang = int( float32(pos) / rng * 180.0)
    ang = ang - 90
    return ang
}

func degToUs(ang int) (us int) {
    if ang < degMin {ang = degMin}
    if ang > degMax {ang = degMax}
    ang +=90
    rng:= float32(servoMax - servoMin)
    us = int((rng / 180.0) * float32(ang))
    return int(servoMin) + us
}

func Open() {
    var err error
// initialise periph
        if _, err := host.Init(); err != nil {
            log.Fatal(err)
    }
        bus, err = i2creg.Open( i2cbus )
        if err != nil {
            log.Fatal(err)
    }
//    fmt.Println("opened ok: ", bus, reflect.TypeOf(bus))
    dev = i2c.Dev{bus, addr}
//    fmt.Println("dev: ", dev, reflect.TypeOf(dev))
    confReg = 0x00
}

func Close() {
    bus.Close()
}

func i2cWriteByte(reg byte, data byte) {
    read := make([]byte, 0)
    write := []byte{reg}
    write = append(write, data)
// write to i2c
    if err := dev.Tx(write, read); err != nil {
        log.Fatal(err)
    }
}

func i2cReadWord(reg byte, litendup bool)(res uint16) {
    read := make([]byte, 2)
    write := []byte{reg}
// get the data
    if err := dev.Tx(write, read); err != nil {
        log.Fatal(err)
    }
// bytes into word
    if litendup {
        res = uint16(int(read[1]) * 256 + int(read[0]))
    } else {
        res = uint16(int(read[0]) * 256 + int(read[1]))
    }
    return res
}

func i2cWriteWord(reg byte, data uint16, litendup bool){
    read := make([]byte, 0)
// load the register
    write := []byte{reg}
// load the data
    x := make([]byte, 2)
    if litendup {
        binary.LittleEndian.PutUint16(x, data)
    } else {
        binary.BigEndian.PutUint16(x, data)
    }
// write to i2c
    write = append(write, x...)
        if err := dev.Tx(write, read); err != nil {
        log.Fatal(err)
    }
    delay(50)
}

func servo(reg byte, ang int) (res string) {
    us := degToUs(ang)
    i2cWriteWord(reg, uint16(us), true)
    return "done"
}

func GetServo(name string) (ang int) {
    if valid(name) {
        reg := srvos[name].reg
        wd := i2cReadWord(reg, true)
        ang = usToDeg(wd)
        return ang
        } else {
            return -255
        }
    }

func ServoEnable(name string, state bool) (res string) {
    if valid(name) {
        srvos[name].abl = state
        if state == true {
            confReg |= (1 << srvos[name].bit)
        } else {
                var mask byte
                mask = ^(1 << srvos[name].bit)
                confReg &= mask
        }
        i2cWriteByte(0x00, confReg)
        delay(50)
        return "done"
    } else {
        return "PtServoEnable: Servo Name Invalid: " + name
    }
}

func Go(pan int, tilt int) {
    servo(srvos["pan"].reg,   pan)
    srvos["pan"].ang = pan
    servo(srvos["tilt"].reg,  tilt)
    srvos["tilt"].ang = tilt
}

func Delta(name string, dlt int) (res string) {
    if valid(name) {
        if dlt > 0 {
            for x := srvos[name].ang ; x < srvos[name].ang + dlt ; x++{
                servo(srvos[name].reg, x)
            }
        } else if  dlt < 0 {
            for x := srvos[name].ang ; x > srvos[name].ang + dlt ; x--{
                servo(srvos[name].reg, x)
            }
        }
        srvos[name].ang = srvos[name].ang + dlt
        return "done"
    } else {
        return "Delta: Servo Name Invalid: " +name
    }
}
