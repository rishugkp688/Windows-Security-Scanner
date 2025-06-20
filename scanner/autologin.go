package scanner

import (
	"golang.org/x/sys/windows/registry"
)

func CheckAutoLogin() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon`, registry.QUERY_VALUE)
	if err != nil {
		return "⚠️ Could not open Winlogon registry key"
	}
	defer key.Close()

	val, _, err := key.GetStringValue("AutoAdminLogon")
	if err != nil {
		return "✅ Auto-login is not configured"
	}
	if val == "1" {
		return "❌ Auto-login is enabled"
	}
	return "✅ Auto-login is disabled"
}
