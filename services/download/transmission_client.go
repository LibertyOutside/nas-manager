package download

import (
	"context"
	"github.com/hekmon/transmissionrpc/v2"
	log "github.com/sirupsen/logrus"
	"nas-manager/db"
	"nas-manager/models"
)

var TrClient []*models.TransmissionClient

func InitTransmissionClient() error {
	var clients []models.TransmissionClient
	db.DB.Find(&clients)
	for _, client := range clients {
		trInstance, err := transmissionrpc.New(client.Host, client.Username, client.Password,
			&transmissionrpc.AdvancedConfig{
				HTTPS:       client.Https,
				Port:        client.Port,
				RPCURI:      client.RpcUri,
				HTTPTimeout: client.HttpTimeout,
				UserAgent:   client.UserAgent,
				Debug:       client.Debug,
			})
		if err != nil {
			return err
		}
		client.Instance = trInstance
		TrClient = append(TrClient, &client)
	}
	return nil
}

func GetTrTorrents() ([]models.Torrent, error) {
	var torrents []transmissionrpc.Torrent
	for _, client := range TrClient {
		clientTorrents, err := client.Instance.TorrentGetAll(context.TODO())
		if err != nil {
			return nil, err
		}
		torrents = append(torrents, clientTorrents...)
	}
	return transferTorrentEntity(torrents), nil
}

func transferTorrentEntity(tOrigins []transmissionrpc.Torrent) []models.Torrent {
	var result []models.Torrent
	for _, t := range tOrigins {
		result = append(result, models.Torrent{
			ID:           *t.ID,
			Name:         *t.Name,
			TorrentSize:  uint64(*t.TotalSize),
			DownloadDir:  *t.DownloadDir,
			Status:       t.Status.String(),
			RateDownload: *t.RateDownload,
			RateUpload:   *t.RateUpload,
			MagicLink:    *t.MagnetLink,
			LeftTime:     *t.LeftUntilDone,
		})
	}
	return result
}

func init() {
	err := InitTransmissionClient()
	if err != nil {
		log.Errorf("Failed to Init Transmission Clients")
	}
}
