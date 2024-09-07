package model

type SpotifyAddItemToPlaybackQueueBodyReq struct {
	URI      string `json:"uri"`
	DeviceId string `json:"device_id"`
}
