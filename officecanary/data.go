package officecanary

import (
	"time"
)

type Datapoint struct {
	DeviceEUI EUI
	Timestamp time.Time

	AppEUI     EUI
	GatewayEUI EUI
	DataRate   string
	DevAddr    string
	Frequency  float64
	RSSI       int32
	SNR        float64
	Payload    []byte

	CO2PPM     uint16
	CO2Status  CO2Status
	Resistance uint32

	// Total Volatile Organic Compound equivalent in Parts Per Billion
	TVOCPPB uint16
}

type CO2Status uint8

const (
	CO2StatusOK    CO2Status = 0x00
	CO2StatusBusy            = 0x01
	CO2StatusRunin           = 0x10
	CO2StatusError           = 0x80
)
