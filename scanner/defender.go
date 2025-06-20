package scanner

import (
	"golang.org/x/sys/windows/registry"
)

func CheckDefender() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows Defender`, registry.QUERY_VALUE)
	if err != nil {
		return "⚠️ Could not read Defender registry"
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue("DisableAntiSpyware")
	if err != nil {
		return "⚠️ Could not determine Defender status"
	}
	if val == 1 {
		return "❌ Windows Defender is disabled"
	}
	return "✅ Windows Defender is active"
}
