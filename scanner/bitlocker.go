package scanner

import (
	"os/exec"
	"strings"
)

func CheckBitLocker() string {
	cmd := exec.Command("powershell", "-Command", "Get-BitLockerVolume | Select-Object -ExpandProperty ProtectionStatus")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "⚠️ BitLocker check failed (requires PowerShell & admin rights)"
	}

	output := strings.TrimSpace(string(out))
	if output == "" {
		return "❌ BitLocker is not enabled (no volumes protected)"
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "1" {
			return "✅ BitLocker is enabled on at least one drive"
		}
	}
	return "❌ BitLocker is not enabled"
}
