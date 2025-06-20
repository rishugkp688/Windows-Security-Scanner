package scanner

import (
	"os/exec"
	"strings"
)

func CheckGuestAccount() string {
	out, err := exec.Command("net", "user", "guest").Output()
	if err != nil {
		return "⚠️ Could not check Guest account"
	}
	if strings.Contains(string(out), "Account active               Yes") {
		return "❌ Guest account is enabled"
	}
	return "✅ Guest account is disabled"
}
