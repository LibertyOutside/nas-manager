package download

import (
	"context"
	"fmt"
	"github.com/hekmon/transmissionrpc/v2"
	"testing"
)

func TestGetTrClient(t *testing.T) {
	/*	trconf := models.TransmissionInfo{
		Host:     "192.168.1.12",
		User:     "",
		Password: "",
		Conf:     nil,
	}*/
	err := InitTransmissionClient()

	if err != nil {
		_ = fmt.Errorf("failed to create Transmission Client:%v", err)
	}
	c := TrClient[0].Instance

	ok, serverVersion, serverMinimumVersion, err := c.RPCVersion(context.TODO())
	torrents, _ := c.TorrentGetAll(context.TODO())
	for _, torrent := range torrents {
		fmt.Printf("%#v\n", torrent)
	}
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(fmt.Sprintf("Remote transmission RPC version (v%d) is incompatible with the transmission library (v%d): remote needs at least v%d",
			serverVersion, transmissionrpc.RPCVersion, serverMinimumVersion))
	}
	fmt.Printf("Remote transmission RPC version (v%d) is compatible with our transmissionrpc library (v%d)\n",
		serverVersion, transmissionrpc.RPCVersion)
}
