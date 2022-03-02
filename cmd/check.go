package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func checkPi() {
	CPUINFO := "/proc/cpuinfo"
	MODEL := "Raspberry Pi 4"

	data, err := os.ReadFile(CPUINFO)

	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(data), MODEL) {
		color.Red("It is not possible to run the DeksPiPro-CLI on this device. \nIt only can be run on Raspberry Pi 4 with DeskPiPro Case.")
		os.Exit(1)
	}
}
