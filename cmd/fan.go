package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

func fanSet(value string) {
	DATA := "pwm_"
	COM := "/dev/ttyUSB0"

	for i := 0; i < 3-len(value); i++ {
		DATA += "0"
	}

	DATA += value

	err := ioutil.WriteFile(COM, []byte(DATA), 0660)
	if err != nil {
		fmt.Println(err)
	}
}

func getTemp() int {
	COM := "/sys/class/thermal/thermal_zone0/temp"

	data, err := os.ReadFile(COM)
	if err != nil {
		log.Fatal(err)
	}
	tempInt, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		log.Fatal("Error when converting Temperature to Float")
	}

	tempInt = math.Round(tempInt / 1000) // Convert to Celsius and get integer number
	return int(tempInt)
}
