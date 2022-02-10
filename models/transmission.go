package models

import "github.com/hekmon/transmissionrpc/v2"

type TransmissionInfo struct {
	Host     string                          `json:"host"`
	User     string                          `json:"user"`
	Password string                          `json:"password"`
	Conf     *transmissionrpc.AdvancedConfig `json:"conf"`
}

type Torrent struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	TorrentSize string `json:"torrent_size"`
	DownloadDir string `json:"download_dir"`
	IsFinished  bool   `json:"is_finished"`
	IsStalled   bool   `json:"is_stalled"` //是否暂停
	//Peers *transmissionrpc.Peer `json:"peers"` //todo:peers
	RateDownload int64 `json:"rate_download"`
	RateUpload   int64 `json:"rate_upload"`
}
