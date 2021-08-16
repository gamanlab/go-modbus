package main

import (
	"fmt"
	"time"

	"github.com/gamanlab/go-modbus"
)

func main() {
	handler := modbus.NewRTUUDPClientHandler("192.168.1.31:8001")
	handler.SlaveID = 1
	client := modbus.NewClient(handler)

	for {
		res, err := client.ReadHoldingRegisters(187, 48)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
		time.Sleep(time.Second)
	}
}
