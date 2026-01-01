package helpers

/*
#cgo LDFLAGS: -framework CoreGraphics -framework Foundation
#include <CoreGraphics/CoreGraphics.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type DisplayCGO struct {
	ID          uint32
	Width       uint32
	Height      uint32
	RefreshRate float64
	IsMain      bool
}

type Display struct {
	Resolution string
	Primary    bool
}

type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func GetDarwinDisplaysCGO() ([]DisplayCGO, error) {
	const maxDisplays = 32
	displays := make([]C.CGDirectDisplayID, maxDisplays)
	var displayCount C.uint32_t

	result := C.CGGetActiveDisplayList(C.uint32_t(maxDisplays),
		(*C.CGDirectDisplayID)(unsafe.Pointer(&displays[0])),
		&displayCount)

	if result != C.kCGErrorSuccess {
		return nil, fmt.Errorf("failed to get display list: error code %d", result)
	}

	displayInfos := make([]DisplayCGO, displayCount)

	for i := C.uint32_t(0); i < displayCount; i++ {
		displayID := displays[i]
		displayInfos[i] = getDisplayInfo(displayID)
	}

	return displayInfos, nil
}

func getDisplayInfo(displayID C.CGDirectDisplayID) DisplayCGO {
	display := DisplayCGO{
		ID: uint32(displayID),
	}

	mode := C.CGDisplayCopyDisplayMode(displayID)

	if unsafe.Pointer(mode) != nil { //nolint:govet
		defer C.CGDisplayModeRelease(mode)
		display.Width = uint32(C.CGDisplayModeGetWidth(mode))
		display.Height = uint32(C.CGDisplayModeGetHeight(mode))
		display.RefreshRate = float64(C.CGDisplayModeGetRefreshRate(mode))
	}

	display.IsMain = C.CGDisplayIsMain(displayID) != 0

	return display
}

func GetDisplays() []Display {
	darwinDisplaysCGO, err := GetDarwinDisplaysCGO()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	var displays []Display
	for _, darwinDisplay := range darwinDisplaysCGO {
		displays = append(displays, Display{
			Resolution: fmt.Sprintf("%dx%d @ %.0fHz", darwinDisplay.Width, darwinDisplay.Height, darwinDisplay.RefreshRate),
			Primary:    darwinDisplay.IsMain,
		})
	}
	return displays
}
