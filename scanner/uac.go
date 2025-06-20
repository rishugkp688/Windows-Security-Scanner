package scanner

import (
	"golang.org/x/sys/windows/registry"
)

func CheckUAC() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
	if err != nil {
		return "⚠️ Could not open UAC registry key"
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue("EnableLUA")
	if err != nil {
		return "⚠️ Could not read EnableLUA (UAC status)"
	}
	if val == 1 {
		return "✅ UAC is enabled"
	}
	return "❌ UAC is disabled"
}
