package officecanary

import (
	"context"
	"crypto/tls"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	aviary "github.com/telenordigital/office-canary-go/proto/aviary"
	officecanary "github.com/telenordigital/office-canary-go/proto/officecanary"
)

type Client struct {
	client aviary.AviaryClient
}

func NewClient(token string) (*Client, error) {
	addr := "oc1.aviary.services:443"
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(tokenCredentials(token)),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: aviary.NewAviaryClient(conn),
	}, nil
}

func (c *Client) StreamDatapoints(handler func(Datapoint)) error {
	stream, err := c.client.StreamDatapoints(context.Background(), &aviary.StreamDatapointsRequest{})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		d, err := datapointFromProto(resp.Datapoint)
		if err != nil {
			log.Println("Error parsing datapoint:", err)
			continue
		}
		handler(d)
	}
}

type tokenCredentials string

func (c tokenCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"x-api-token": string(c)}, nil
}

func (c tokenCredentials) RequireTransportSecurity() bool { return true }

func datapointFromProto(d *aviary.Datapoint) (Datapoint, error) {
	deveui, err := EUIFromString(d.DeviceEui)
	if err != nil {
		return Datapoint{}, err
	}

	ts, err := ptypes.Timestamp(d.Timestamp)
	if err != nil {
		return Datapoint{}, err
	}

	v := officecanary.Datapoint{}
	if err := ptypes.UnmarshalAny(d.Value, &v); err != nil {
		return Datapoint{}, err
	}

	appeui, err := EUIFromString(v.AppEui)
	if err != nil {
		return Datapoint{}, err
	}

	gatewayeui, err := EUIFromString(v.GatewayEui)
	if err != nil {
		return Datapoint{}, err
	}

	return Datapoint{
		DeviceEUI:  deveui,
		Timestamp:  ts,
		AppEUI:     appeui,
		GatewayEUI: gatewayeui,
		DataRate:   v.DataRate,
		DevAddr:    v.DevAddr,
		Frequency:  v.Frequency,
		RSSI:       v.Rssi,
		SNR:        v.Snr,
		Payload:    v.Payload,
		CO2PPM:     uint16(v.Co2Ppm),
		CO2Status:  CO2Status(v.Co2Status),
		Resistance: v.Resistance,
		TVOCPPB:    uint16(v.TvocPpb),
	}, nil
}
