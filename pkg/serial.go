package serial

import (
	"bufio"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-serial-go/serial_config"
	"go.bug.st/serial"
	"log"
)

type Connection struct {
	Port           serial.Port
	Enable         bool
	Connected      bool
	Error          bool
	ActivePortList []string
}

var Port Connection

func (c *Connection) Disconnect() error {
	return c.Port.Close()
}

type TMicroEdge struct {
	Sensor        string
}

var CH = make(chan TMicroEdge)

func NewSerialConnection() {

	c := serial_config.GetSerialConfig()
	portName := c.Port
	baudRate := c.BaudRate
	parity := c.Parity
	stopBits := c.StopBits
	dataBits := c.DataBits


	if Port.Connected {
		log.Println("Existing serial port connection by this app is open So! close existing connection")
		err := Port.Disconnect()
		if err != nil {
			return
		}
	}

	m := &serial.Mode{
		BaudRate: baudRate,
		Parity:   parity,
		DataBits: dataBits,
		StopBits: stopBits,
	}

	ports, err := serial.GetPortsList()
	Port.ActivePortList = ports
	log.Println("SerialPort try and connect to", portName)
	log.Println(ports)
	port, err := serial.Open(portName, m)
	Port.Port = port

	if err != nil {
		Port.Error = true
		log.Fatal("ERROR Connected to serial port", err)
	}
	Port.Connected = true
	log.Println("Connected to serial port", portName)

	scanner := bufio.NewScanner(port)
	count := 0



	for scanner.Scan() {
		var data = scanner.Text()
		count = count + 1
		//var mychannel chan json.Decoder

		message := TMicroEdge{Sensor: data}
		//jsonValue, _ := json.Marshal(message)
		CH <- message
		log.Println("MQTT messages, topic:", " ", "data:", data)
		fmt.Println(data)
		fmt.Println(len(data), "size")
	}

}
