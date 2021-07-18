package main

import (
	"fmt"
	serial "github.com/NubeIO/nubeio-rubix-lib-serial-go/pkg"
	"github.com/NubeIO/nubeio-rubix-lib-serial-go/serial_config"
	"log"
)

func main() {

	var args serial_config.Params
	args.UseConfigFile = false

	var config serial_config.Serial
	config.Port = "/dev/ttyACM0"
	config.BaudRate = 38400

	err := serial_config.SetSerialConfig(config, args); if err != nil {
		log.Println(err)
		return
	}


	go serial.NewSerialConnection()
	msg := <- serial.CH
	fmt.Println(msg)

}
