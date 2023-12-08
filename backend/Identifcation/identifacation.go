package identifcation

import (
	"fmt"
	"runtime"

	"github.com/denisbrodbeck/machineid"
)

func GetHWID() (string, error) {
	switch os := runtime.GOOS; os {
	case "windows", "linux":
		return machineid.ProtectedID("MTGA-Analytics")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", os)
	}
}
