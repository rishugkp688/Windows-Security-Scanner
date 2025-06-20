# 🔐 Windows Security Scanner

A simple yet powerful Go-based tool to scan your Windows system for common security misconfigurations and risks.  
It checks key system settings such as UAC, Firewall, BitLocker, Defender, and more — and helps identify potential hardening opportunities.

---

## 🚀 Features

- ✅ Detects if **UAC (User Account Control)** is enabled  
- ✅ Checks **Windows Firewall** status across all profiles (Domain, Private, Public)  
- ✅ Verifies if the **Guest account** is disabled  
- ✅ Detects **Auto-login** configuration (which is insecure)  
- ✅ Confirms if **Windows Defender** is active  
- ✅ Reports **BitLocker** encryption status  
- ✅ Detects active **antivirus software** via WMI  

---

## 📦 Installation

### 1. **Clone this repository**
```bash
git clone https://github.com/yourname/win-security-scanner.git
cd win-security-scanner
go build -o scanner.exe
./scanner.exe
```
And that's all for now I will say that I made this project with the help of a llm but I will try to make it better.
