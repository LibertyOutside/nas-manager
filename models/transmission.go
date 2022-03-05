package models

import (
	"github.com/hekmon/transmissionrpc/v2"
	"gorm.io/gorm"
	"time"
)

//see transmission rpc doc here: https://github.com/transmission/transmission/blob/main/docs/rpc-spec.md

type TransmissionInfo struct {
	Host     string                          `json:"host"`
	User     string                          `json:"user"`
	Password string                          `json:"password"`
	Conf     *transmissionrpc.AdvancedConfig `json:"conf"`
}

type TransmissionClient struct {
	gorm.Model
	Alias       string //别名
	Host        string
	Username    string
	Password    string
	Https       bool
	Port        uint16
	RpcUri      string
	HttpTimeout time.Duration
	UserAgent   string
	Debug       bool
	Instance    *transmissionrpc.Client `gorm:"-" json:"-"`
}

type Torrent struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	TorrentSize uint64 `json:"torrent_size"`
	DownloadDir string `json:"download_dir"`
	Status      string `json:"status"`
	//Peers *transmissionrpc.Peer `json:"peers"` //todo:peers
	RateDownload int64  `json:"rate_download"`
	RateUpload   int64  `json:"rate_upload"`
	MagicLink    string `json:"magic_link"`
	LeftTime     int64  `json:"left_time"`
}
