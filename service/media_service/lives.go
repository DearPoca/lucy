package media_service

import (
	"errors"
	"fmt"
	"time"

	"lucy/models"
	"lucy/pkg/log"
	"lucy/pkg/setting"
	"lucy/utils"
)

const liveTokenLength = 32
const app = "lucy"

const noRecord = "none"
const recording = "recording"
const available = "available"

type Live struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Owner     string `json:"owner,omitempty"`
	WebrtcUrl string `json:"webrtc_link,omitempty"`
	RtmpUrl   string `json:"rtmp_link,omitempty"`
	FlvUrl    string `json:"flv_link,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	RecordUrl string `json:"record_url,omitempty"`
}

func GetLiveByStream(stream *Stream) (*Live, error) {
	if !stream.Publish.Active {
		return nil, errors.New("stream inactive")
	}
	if !VerifyLiveName(stream.Url) {
		return nil, errors.New("stream format error")
	}
	l := models.Live{}
	err := models.Db().Where(models.Live{Name: stream.Url}).First(&l).Error

	if err != nil {
		return nil, err
	} else {
		ret := Live{
			Id:        stream.Id,
			Name:      l.Name,
			Owner:     l.Owner,
			WebrtcUrl: l.WebrtcUrl,
			RtmpUrl:   l.RtmpUrl,
			FlvUrl:    l.HttpFlvUrl,
			StartTime: l.CreatedAt.String(),
		}
		if l.RecordStatus == available {
			ret.RecordUrl = l.RecordPath
		}
		return &ret, nil
	}
}

func GetActiveLives() []Live {
	streams := GetStreams()
	lives := make([]Live, 0)
	for i, _ := range streams {
		r, err := GetLiveByStream(&streams[i])
		if err != nil {
			log.Warn("stream not valid", "stream", streams[i])
		} else {
			lives = append(lives, *r)
		}

	}
	return lives
}

func GetLivesByUser(username string) ([]Live, error) {
	var ls []models.Live
	err := models.Db().Where(models.Live{Owner: username}).Find(&ls).Error
	var ret []Live
	if err != nil {
		log.Info("user have no live", "username", username)
		return ret, errors.New("user have no live")
	}
	for i, _ := range ls {
		live := Live{
			Name:      ls[i].Name,
			Owner:     ls[i].Owner,
			WebrtcUrl: ls[i].WebrtcUrl,
			RtmpUrl:   ls[i].RtmpUrl,
			FlvUrl:    ls[i].HttpFlvUrl,
			StartTime: ls[i].CreatedAt.String(),
		}
		if ls[i].RecordStatus == available {
			live.RecordUrl = ls[i].RecordPath
		}
		ret = append(ret, live)
	}
	return ret, nil
}

func LiveRecord(liveName string, username string) error {
	log.Debug("LiveRecord", "liveName", liveName, "username", username)
	if !VerifyLiveName(liveName) {
		return errors.New("live name format error")
	}
	if username != ParseUserFromLivePath(liveName) {
		return errors.New("requester are not owner")
	}

	var l models.Live
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
	ffmpegRecord(l.HttpFlvUrl, liveName)
	return nil
}

func GenerateLive(username string) (*Live, error) {
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
	livaName := fmt.Sprintf("/%s/%s/%s", app, username, string(buf))
	l := &models.Live{
		Name:  livaName,
		Owner: username,
		WebrtcUrl: fmt.Sprintf("webrtc://%s:%s%s",
			setting.SrsSetting.Ip, setting.SrsSetting.HttpApiPort,
			livaName),
		RtmpUrl: fmt.Sprintf("rtmp://%s:%s%s",
			setting.SrsSetting.Ip, setting.SrsSetting.RtmpPort,
			livaName),
		HttpFlvUrl: fmt.Sprintf("http://%s:%s%s.flv",
			setting.SrsSetting.Ip, setting.SrsSetting.NginxHttpPort,
			livaName),
		RecordStatus: noRecord,
		RecordPath:   "",
	}
	ret := &Live{
		Owner:     username,
		Name:      livaName,
		WebrtcUrl: l.WebrtcUrl,
		RtmpUrl:   l.RtmpUrl,
		FlvUrl:    l.HttpFlvUrl,
	}
	err := models.Db().Create(&l).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
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
