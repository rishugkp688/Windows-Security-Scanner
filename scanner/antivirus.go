package scanner

import (
	"os/exec"
	"strings"
)

func CheckAntivirus() string {
	cmd := exec.Command("powershell", "-Command",
		"Get-CimInstance -Namespace root/SecurityCenter2 -ClassName AntivirusProduct | Select-Object displayName,productState")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "⚠️ Could not query antivirus status (WMI issue)"
	}

	output := strings.TrimSpace(string(out))
	if output == "" {
		return "❌ No antivirus detected via WMI (SecurityCenter2)"
	}

	if strings.Contains(output, "Windows Defender") || strings.Contains(output, "enabled") || strings.Contains(output, "On") {
		return "✅ Antivirus is detected and likely running"
	}
	return "⚠️ Antivirus detected, but state uncertain:\n" + output
}
