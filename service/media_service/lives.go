package media_service

import (
	"errors"
	"fmt"
	"time"

	"lucy/models"
	"lucy/pkg/log"
	"lucy/utils"
)

const liveTokenLength = 32
const app = "lucy"

const noRecord = "none"
const recording = "recording"
const available = "available"

type Live struct {
	Id         string `json:"id"`
	Owner      string `json:"owner"`
	WebrtcLink string `json:"webrtc link"`
	FlvLink    string `json:"flv link"`
}

func GetLives() []Live {
	streams := GetStreams()
	lives := make([]Live, 0)
	for i, _ := range streams {
		if !streams[i].Publish.Active || !VerifyLiveName(streams[i].Url) {
			continue
		}
		r := Live{
			Id:         streams[i].Id,
			Owner:      ParseUserFromLivePath(streams[i].Url),
			WebrtcLink: fmt.Sprintf("/play/webrtc?live_id=%s", streams[i].Id),
			FlvLink:    fmt.Sprintf("/play/flv?live_id=%s", streams[i].Id),
		}
		lives = append(lives, r)
	}
	return lives
}

func Record(streamUrl string, username string) error {
	i := len(streamUrl) - 1
	formatErr := errors.New("rtmp url format error")
	log.Debug("Record", "streamUrl", streamUrl, "username", username)
	if len(streamUrl) < 4 || streamUrl[len(streamUrl)-4:] != ".flv" {
		return formatErr
	}
	for j := 0; j < 3; j++ {
		for i >= 0 && streamUrl[i] != '/' {
			i--
		}
		if i < 0 {
			return formatErr
		}
		i--
	}
	liveName := streamUrl[i+1 : len(streamUrl)-4]
	log.Debug("Get liveName", "liveName", liveName)

	if !VerifyLiveName(liveName) {
		return formatErr
	}
	if username != ParseUserFromLivePath(liveName) {
		return errors.New("requester are not owner")
	}

	l := models.Live{}
	err := models.Db().Where(models.Live{Name: liveName}).First(&l).Error
	if err != nil {
		log.Warn("Can not find live", "liveName", liveName)
		return err
	}
	if l.RecordStatus != noRecord {
		return errors.New("recording has been started")
	}
	err = models.Db().Model(&l).Where(models.Live{Name: liveName}).
		Update("record_status", recording).Error
	if err != nil {
		log.Warn("update live failed", "liveName", liveName)
		return err
	}
	ffmpegRecord(streamUrl, liveName)
	return nil
}

func GenerateLive(username string) (string, error) {
	lToA := func(buf *[]byte, i int64, wid int) {
		var b [liveTokenLength + 1]byte
		bp := len(b) - 1
		for i >= 10 || wid > 1 {
			wid--
			q := i / 10
			b[bp] = byte('0' + i - q*10)
			bp--
			i = q
		}
		b[bp] = byte('0' + i)
		*buf = append(*buf, b[bp:]...)
	}
	buf := make([]byte, 0)
	lToA(&buf, time.Now().UnixNano(), liveTokenLength)
	for i := 0; i < len(buf) && buf[i] == '0'; i++ {
		buf[i] = utils.RandStr(1)[0]
	}
	name := fmt.Sprintf("/%s/%s/%s", app, username, string(buf))
	l := &models.Live{
		Name:         name,
		Owner:        username,
		RecordStatus: noRecord,
		RecordPath:   "",
	}
	err := models.Db().Create(l).Error
	if err != nil {
		return "", err
	}
	return name, nil
}

func ParseUserFromLivePath(path string) string {
	return path[len(app)+2 : len(path)-liveTokenLength-1]
}

func VerifyLiveName(name string) bool {
	if len(name) < liveTokenLength+len(app)+3 {
		return false
	}
	if name[1:len(app)+1] != app {
		return false
	}
	if name[len(name)-liveTokenLength-1] != '/' {
		return false
	}
	return true
}
