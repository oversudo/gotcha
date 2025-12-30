package helpers

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	ENUM_CURRENT_SETTINGS         = 0xFFFFFFFF
	DISPLAY_DEVICE_ACTIVE         = 0x00000001
	DISPLAY_DEVICE_PRIMARY_DEVICE = 0x00000004
)

type DEVMODE struct {
	DeviceName       [32]uint16
	SpecVersion      uint16
	DriverVersion    uint16
	Size             uint16
	DriverExtra      uint16
	Fields           uint32
	Orientation      int16
	PaperSize        int16
	PaperLength      int16
	PaperWidth       int16
	Scale            int16
	Copies           int16
	DefaultSource    int16
	PrintQuality     int16
	Color            int16
	Duplex           int16
	YResolution      int16
	TTOption         int16
	Collate          int16
	FormName         [32]uint16
	LogPixels        uint16
	BitsPerPel       uint32
	PelsWidth        uint32
	PelsHeight       uint32
	DisplayFlags     uint32
	DisplayFrequency uint32
	ICMMethod        uint32
	ICMIntent        uint32
	MediaType        uint32
	DitherType       uint32
	Reserved1        uint32
	Reserved2        uint32
	PanningWidth     uint32
	PanningHeight    uint32
}

type DISPLAY_DEVICE struct {
	Size         uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	StateFlags   uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

type Display struct {
	Resolution string
	Primary    bool
	Name       string
}

func GetDisplays() []Display {
	user32 := syscall.NewLazyDLL("user32.dll")
	enumDisplayDevices := user32.NewProc("EnumDisplayDevicesW")
	enumDisplaySettings := user32.NewProc("EnumDisplaySettingsW")

	var displayDevice DISPLAY_DEVICE
	displayDevice.Size = uint32(unsafe.Sizeof(displayDevice))

	displayIndex := uint32(0)
	var displays []Display
	for {
		ret, _, _ := enumDisplayDevices.Call(
			0,
			uintptr(displayIndex),
			uintptr(unsafe.Pointer(&displayDevice)),
			0,
		)

		if ret == 0 {
			break
		}

		if displayDevice.StateFlags&0x00000001 != 0 {
			deviceName := syscall.UTF16ToString(displayDevice.DeviceName[:])
			isPrimary := displayDevice.StateFlags&DISPLAY_DEVICE_PRIMARY_DEVICE != 0

			var devMode DEVMODE
			devMode.Size = uint16(unsafe.Sizeof(devMode))

			ret2, _, _ := enumDisplaySettings.Call(
				uintptr(unsafe.Pointer(&displayDevice.DeviceName[0])),
				ENUM_CURRENT_SETTINGS,
				uintptr(unsafe.Pointer(&devMode)),
			)

			if ret2 != 0 {
				displays = append(displays, Display{
					Resolution: fmt.Sprintf("%dx%d @ %d Hz", devMode.PelsWidth, devMode.PelsHeight, devMode.DisplayFrequency),
					Primary:    isPrimary,
					Name:       deviceName,
				})
			}
		}

		displayIndex++
	}
	return displays
}
