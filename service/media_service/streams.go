package media_service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"lucy/pkg/log"
)

type Kbps struct {
	Recv int `json:"recv_30s"`
	Send int `json:"send_30s"`
}

type Publish struct {
	Active bool   `json:"active"`
	Cid    string `json:"cid"`
}

type Video struct {
	Codec   string `json:"codec"`
	Profile string `json:"profile"`
	Level   string `json:"level"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type Audio struct {
	Codec      string `json:"codec"`
	SampleRate int    `json:"sample_rate"`
	Channel    int    `json:"channel"`
	Profile    string `json:"profile"`
}

type Stream struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Vhost     string   `json:"vhost"`
	App       string   `json:"app"`
	TcUrl     string   `json:"tcUrl"`
	Url       string   `json:"url"`
	LiveMs    int64    `json:"live_ms"`
	Clients   int      `json:"clients"`
	Frames    int      `json:"frames"`
	SendBytes int      `json:"send_bytes"`
	RecvBytes int      `json:"recv_bytes"`
	Kbps      *Kbps    `json:"kbps"`
	Publish   *Publish `json:"publish"`
	Video     *Video   `json:"video"`
	Audio     *Audio   `json:"audio"`
}

type StreamsResponse struct {
	Code    int      `json:"code"`
	Server  string   `json:"server"`
	Service string   `json:"service"`
	Pid     string   `json:"pid"`
	Streams []Stream `json:"streams"`
}

func GetStreams() []Stream {
	var myClient = &http.Client{}
	resp, err := myClient.Get(fmt.Sprintf("%s/api/v1/streams/", httpApiPath))
	if err != nil {
		log.Warn("Get streams failed", "err", err.Error())
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var streamsResp StreamsResponse
	err = json.Unmarshal(body, &streamsResp)
	if err != nil {
		log.Warn("Get streams failed", "err", err.Error())
		return nil
	}
	log.Info("Get response", "streamsResp", streamsResp)
	return streamsResp.Streams
}
