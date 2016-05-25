package main

import (
	"fmt"
	"github.com/deadsy/libusb"
	"github.com/k0kubun/pp"
	"log"
	"os"
)

const (
	AX206_VID = 0x1908 // Hacked frames USB Vendor ID
	AX206_PID = 0x0102 // Hacked frames USB Product ID

	USBCMD_SETPROPERTY = 0x01 // USB command: Set property
	USBCMD_BLIT        = 0x12 // USB command: Blit to screen

)

func main() {

	log.SetOutput(os.Stderr)

	var ctx libusb.Context
	err := libusb.Init(&ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer libusb.Exit(ctx)

	hdl := libusb.Open_Device_With_VID_PID(ctx, AX206_VID, AX206_PID)
	if hdl == nil {
		log.Fatal("Failed to open device")
	}
	defer libusb.Close(hdl)

	/*
		dev := libusb.Get_Device(hdl)
		dd, err := libusb.Get_Device_Descriptor(dev)
		if err != nil {
			log.Fatal(err)
		}
		pp.Println(dd)

		// record information on input endpoints
		for i := 0; i < int(dd.BNumConfigurations); i++ {
			cd, err := libusb.Get_Config_Descriptor(dev, uint8(i))
			if err != nil {
				log.Fatal(err)
			}
			pp.Println(cd)

			libusb.Free_Config_Descriptor(cd)
		}
	*/

	/*
		if err := libusb.Claim_Interface(hdl, 0); err != nil {
			log.Fatal(err)
		}
		defer libusb.Release_Interface(hdl, 0)
	*/

	w, h, err := GetLcdParams(hdl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got LCD dimensions: %dx%d\n", w, h)

}

func GetLcdParams(udev libusb.Device_Handle) (width, height int, err error) {
	cmd := []byte{
		0xcd, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
	data, err := ScsiRead(udev, cmd, 5)
	if err != nil {
		return 0, 0, err
	}
	pp.Println(data)
	width = int(data[0]) | int(data[1])<<8
	height = int(data[2]) | int(data[3])<<8
	return width, height, nil
}

var exCmd = []byte{
	0xcd, 0, 0, 0,
	0, 6, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
}

var scsiCmd = []byte{
	0x55, 0x53, 0x42, 0x43, // dCBWSignature
	0xde, 0xad, 0xbe, 0xef, // dCBWTag
	0x00, 0x80, 0x00, 0x00, // dCBWLength
	0x00, // bmCBWFlags: 0x80: data in (dev to host), 0x00: Data out
	0x00, // bCBWLUN
	0x10, // bCBWCBLength

	// SCSI cmd:
	0xcd, 0x00, 0x00, 0x00,
	0x00, 0x06, 0x11, 0xf8,
	0x70, 0x00, 0x40, 0x00,
	0x00, 0x00, 0x00, 0x00,
}

func scsiCmdPrepare(cmd []byte, blockLen int) []byte {
	buf := scsiCmd
	buf[8] = byte(blockLen)
	buf[9] = byte(blockLen >> 8)
	buf[10] = byte(blockLen >> 16)
	buf[11] = byte(blockLen >> 24)

	copy(buf[15:], cmd)

	pp.Println(buf)
	return buf
}

const scsiTimeout = 1000

func ScsiWrite(udev libusb.Device_Handle, cmd []byte, data []byte) error {
	var err error

	// Write command to device
	log.Print("Write command to device")
	_, err = libusb.Bulk_Transfer(udev, 0x01, scsiCmdPrepare(cmd, len(data)), scsiTimeout)
	if err != nil {
		return err
	}

	// Write data to device
	log.Print("Write data to device")
	_, err = libusb.Bulk_Transfer(udev, 0x01, data, scsiTimeout)
	if err != nil {
		return err
	}

	// Get ACK
	log.Print("Read ACK from device")
	ack, err := libusb.Bulk_Transfer(udev, 0x81, []byte{}, scsiTimeout)
	if err != nil {
		return err
	}

	if string(ack[:4]) != "USBS" {
		return fmt.Errorf("Got invalid reply")
	}
	log.Print("Write reply: ", ack)
	// pass back return code set by peer:
	// return ansbuf[12];
	return nil
}

func ScsiRead(hdl libusb.Device_Handle, cmd []byte, blockLen int) ([]byte, error) {
	// Write command to device
	log.Print("Write command to device")
	_, err := libusb.Bulk_Transfer(hdl, 0x01, scsiCmdPrepare(cmd, blockLen), scsiTimeout)
	if err != nil {
		return nil, err
	}

	log.Print("Read data from device")
	// Read data from device
	return libusb.Bulk_Transfer(hdl, 0x81, []byte{0, 0, 0, 0, 0}, scsiTimeout)
}
