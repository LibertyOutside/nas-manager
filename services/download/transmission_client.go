package download

import (
	"github.com/hekmon/transmissionrpc/v2"
	"nas-manager/models"
)

var TrClient *transmissionrpc.Client

func InitTransmissionClient(info models.TransmissionInfo) error {

	client, err := transmissionrpc.New(info.Host, info.User, info.Password, info.Conf)
	if err != nil {
		return err
	}
	TrClient = client
	return nil
}
