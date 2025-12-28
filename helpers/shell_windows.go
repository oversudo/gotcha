package helpers

import (
	"os"
	"path/filepath"
)

func GetShellInfo() string {
	if os.Getenv("PSModulePath") != "" {
		if os.Getenv("PSEdition") == "Core" {
			return "pwsh.exe"
		}
		return "powershell.exe"
	}

	if comspec := os.Getenv("COMSPEC"); comspec != "" {
		return filepath.Base(comspec)
	}

	return "cmd.exe" // fallback
}
