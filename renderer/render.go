package renderer

import (
	"fmt"
	"strings"

	"github.com/oversudo/gofetch/helpers"
)

func Render() {
	header := fmt.Sprintf("%s@%s\n", helpers.GetUsername(), helpers.GetHostname())
	fmt.Printf("%s@%s\n", helpers.GetUsername(), helpers.GetHostname())
	fmt.Println(strings.Repeat("-",len(header)))
	fmt.Printf("OS: %s\n", helpers.GetOSInfo())
	fmt.Printf("Uptime: %s",helpers.GetUptime())
}
