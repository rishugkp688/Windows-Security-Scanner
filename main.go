package main

import (
	"fmt"

	"github.com/rishugkp688/WinConfigScanner/scanner"
)

func main() {
	fmt.Println("🔍 Windows Security Scanner")

	fmt.Println(scanner.CheckUAC())
	fmt.Println(scanner.CheckFirewall())
	fmt.Println(scanner.CheckGuestAccount())
	fmt.Println(scanner.CheckAutoLogin())
	fmt.Println(scanner.CheckDefender())
	fmt.Println(scanner.CheckBitLocker())
	fmt.Println(scanner.CheckAntivirus())
}
