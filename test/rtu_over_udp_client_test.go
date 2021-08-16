// Copyright 2018 xft. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gamanlab/go-modbus"
)

const (
	rtuOverUDPDevice = "192.168.100.84:8001"
)

func TestRTUOverUDPClient(t *testing.T) {
	// Diagslave does not support broadcast id.
	handler := modbus.NewRTUUDPClientHandler(rtuOverUDPDevice)
	handler.SlaveID = 1
	client := modbus.NewClient(handler)
	res, err := client.ReadHoldingRegisters(187, 48)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
	// ClientTestAll(t, modbus.NewClient(handler))
}

func TestRTUOverUDPClientAdvancedUsage(t *testing.T) {
	handler := modbus.NewRTUOverTCPClientHandler(rtuOverUDPDevice)
	handler.Timeout = 5 * time.Second
	handler.SlaveID = 1
	handler.Logger = log.New(os.Stdout, "ascii over tcp: ", log.LstdFlags)
	handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadDiscreteInputs(15, 2)
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
	results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
	results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
}
