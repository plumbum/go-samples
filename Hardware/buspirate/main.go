package main

import (
	"log"
	"github.com/plumbum/buspirate"
	"time"
)

func main() {
	bp, err := buspirate.Open("/dev/ttyACM0")
	if err != nil {
		log.Fatal(err)
	}

	bp.PowerOn()

	chk(bp.SpiEnter())
	chk(bp.SpiConfigure2(true, false, true, false))
	chk(bp.SpiSpeed(buspirate.SpiSpeed1mhz))
	chk(bp.SpiConfigure(true, true, false, true))

	Max7219cmd(bp, 0x09, 0xFF)
	Max7219cmd(bp, 0x0B, 0x07)
	Max7219cmd(bp, 0x0A, 0x08)
	Max7219cmd(bp, 0x0C, 0x01)

	Max7219cmd(bp, 0x01, 0x00)
	Max7219cmd(bp, 0x02, 0x01)
	Max7219cmd(bp, 0x03, 0x02)
	Max7219cmd(bp, 0x04, 0x03)
	Max7219cmd(bp, 0x05, 0x04)
	Max7219cmd(bp, 0x06, 0x05)
	Max7219cmd(bp, 0x07, 0x06)
	Max7219cmd(bp, 0x08, 0x07)


	time.Sleep(time.Second * 3)

	bp.SpiLeave()
	bp.PowerOff()

}

func Max7219cmd(bp *buspirate.BusPirate, cmd, data byte) error {
	var err error

	err = bp.SpiCs(false)
	if err != nil {
		return err
	}
	bp.SpiTransfer([]byte{cmd, data})
	err = bp.SpiCs(true)
	return err
}

func chk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
