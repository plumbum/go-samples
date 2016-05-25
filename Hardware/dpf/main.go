package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deadsy/libusb"
	"image"
	"math/rand"
)

const (
	AX206_VID = 0x1908 // Hacked frames USB Vendor ID
	AX206_PID = 0x0102 // Hacked frames USB Product ID

	USBCMD_SETPROPERTY = 0x01 // USB command: Set property
	USBCMD_BLIT        = 0x12 // USB command: Blit to screen

	ENDP_OUT = 0x01
	ENDP_IN  = 0x81
)

func main() {

	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	libusb.Set_Auto_Detach_Kernel_Driver(hdl, true)

	if err = libusb.Claim_Interface(hdl, 0); err != nil {
		log.Fatal("Claim interface: ", err)
		return
	}
	defer libusb.Release_Interface(hdl, 0)

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

	w, h, err := GetParams(hdl)
	if err != nil {
		log.Fatal("LcdParams: ", err)
	}
	fmt.Printf("Got LCD dimensions: %dx%d\n", w, h)

	for i := range [8]struct{}{} {
		if err := SetBrightness(hdl, i); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Microsecond * 500)
	}

	rand.Seed(time.Now().Unix())
	r := image.Rect(0, 0, 319, 239)
	buf := make([]byte, (r.Max.X-r.Min.X+1)*(r.Max.Y-r.Min.Y+1)*2)

	for _ = range [50]struct{}{} {
		for i := range buf {
			buf[i] = byte(rand.Int())
		}
		err = Blit(hdl, r, buf)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func GetParams(udev libusb.Device_Handle) (width, height int, err error) {
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
	width = int(data[0]) | int(data[1])<<8
	height = int(data[2]) | int(data[3])<<8
	return width, height, nil
}

func SetBrightness(udev libusb.Device_Handle, lvl int) error {
	if lvl < 0 {
		lvl = 0
	}
	if lvl > 7 {
		lvl = 7
	}

	cmd := []byte{
		0xcd, 0, 0, 0,
		0, 6, USBCMD_SETPROPERTY,
		1, 0, // PROPERTY_BRIGHTNESS
		byte(lvl), byte(lvl >> 8),
		0, 0, 0, 0, 0,
	}

	return ScsiWrite(udev, cmd, nil)
}

func Blit(udev libusb.Device_Handle, r image.Rectangle, data []byte) error {

	cmd := []byte{
		0xcd, 0, 0, 0,
		0, 6, USBCMD_BLIT,
		byte(r.Min.X), byte(r.Min.X >> 8),
		byte(r.Min.Y), byte(r.Min.Y >> 8),
		byte(r.Max.X), byte(r.Max.X >> 8),
		byte(r.Max.Y), byte(r.Max.Y >> 8),
		0,
	}
	return ScsiWrite(udev, cmd, data)
}

/*
static
unsigned char g_excmd[16] = {
0xcd, 0, 0, 0,
0, 6, 0, 0,
0, 0, 0, 0,
0, 0, 0, 0
};

void dpf_ax_screen_blit(DPFAXHANDLE h, const unsigned char *buf, short rect[4])
{
unsigned long len = (rect[2] - rect[0]) * (rect[3] - rect[1]);
len <<= 1;
unsigned char *cmd = g_excmd;

cmd[6] = USBCMD_BLIT;
cmd[7] = rect[0];
cmd[8] = rect[0] >> 8;
cmd[9] = rect[1];
cmd[10] = rect[1] >> 8;
cmd[11] = rect[2] - 1;
cmd[12] = (rect[2] - 1) >> 8;
cmd[13] = rect[3] - 1;
cmd[14] = (rect[3] - 1) >> 8;
cmd[15] = 0;

wrap_scsi((DPFContext *) h, cmd, sizeof(g_excmd), DIR_OUT, (unsigned char *) buf, len);
}
*/

const scsiTimeout = 1000

func scsiCmdPrepare(cmd []byte, blockLen int, out bool) []byte {
	var bmCBWFlags byte
	if out {
		bmCBWFlags = 0x00
	} else {
		bmCBWFlags = 0x80
	}
	buf := []byte{
		0x55, 0x53, 0x42, 0x43, // dCBWSignature
		0xde, 0xad, 0xbe, 0xef, // dCBWTag
		byte(blockLen), byte(blockLen >> 8), byte(blockLen >> 16), byte(blockLen >> 24), // dCBWLength (4 byte)
		bmCBWFlags,     // bmCBWFlags: 0x80: data in (dev to host), 0x00: Data out
		0x00,           // bCBWLUN
		byte(len(cmd)), // bCBWCBLength

		// SCSI cmd: (15)
		0xcd, 0x00, 0x00, 0x00,
		0x00, 0x06, 0x11, 0xf8,
		0x70, 0x00, 0x40, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	copy(buf[15:], cmd)

	log.Print("SCSI cmd: ", cmd)
	log.Print("SCSI command: ", buf)
	return buf
}

func scsiGetAck(udev libusb.Device_Handle) error {
	buf := []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0,
	}
	// Get ACK
	log.Print("[ACK] Read ACK from device")
	ack, err := libusb.Bulk_Transfer(udev, ENDP_IN, buf, scsiTimeout)
	if err != nil {
		return err
	}
	log.Print("[ACK] data ", ack)

	if string(ack[:4]) != "USBS" {
		return fmt.Errorf("Got invalid reply")
	}
	// pass back return code set by peer:
	// return ansbuf[12];
	return nil
}

func ScsiWrite(hdl libusb.Device_Handle, cmd []byte, data []byte) error {
	var err error

	// Write command to device
	log.Print("[WRITE] Write command to device")
	_, err = libusb.Bulk_Transfer(hdl, ENDP_OUT, scsiCmdPrepare(cmd, len(data), true), scsiTimeout)
	if err != nil {
		return err
	}

	// Write data to device
	if data != nil {
		log.Print("[WRITE] Write data to device")
		out, err := libusb.Bulk_Transfer(hdl, ENDP_OUT, data, scsiTimeout)
		if err != nil {
			return err
		}
		log.Print(out)
	}

	return scsiGetAck(hdl)
}

func ScsiRead(hdl libusb.Device_Handle, cmd []byte, blockLen int) ([]byte, error) {
	var err error

	// Write command to device
	log.Print("[READ] Write command to device")
	_, err = libusb.Bulk_Transfer(hdl, ENDP_OUT, scsiCmdPrepare(cmd, blockLen, false), scsiTimeout)
	if err != nil {
		return nil, err
	}

	log.Print("[READ] Read data from device")
	// Read data from device
	data1 := make([]byte, blockLen, blockLen)
	data, err := libusb.Bulk_Transfer(hdl, ENDP_IN, data1, scsiTimeout)
	if err != nil {
		return nil, err
	}
	log.Print("[READ] data ", data)

	err = scsiGetAck(hdl)
	if err != nil {
		return data, err
	}

	return data, nil
}
