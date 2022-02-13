package cmd

import (
	"fmt"
	"io/ioutil"
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
