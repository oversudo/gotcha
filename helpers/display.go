//go:build !linux
package helpers

type Display struct {
	Name        string
	Resolution  string
	RefreshRate float64
	Primary     bool
}

func GetDisplays() []Display {
	return nil
}