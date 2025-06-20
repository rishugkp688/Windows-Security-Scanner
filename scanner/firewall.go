package scanner

import (
	"os/exec"
	"strings"
)

func CheckFirewall() string {
	cmd := exec.Command("powershell", "-Command", "Get-NetFirewallProfile | Select-Object -ExpandProperty Enabled")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "⚠️ Error checking Windows Firewall with PowerShell"
	}
	results := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, r := range results {
		if strings.TrimSpace(r) == "True" {
			return "✅ Windows Firewall is active on at least one profile"
		}
	}
	return "❌ Windows Firewall is disabled on all profiles"
}
