package media_service

import (
	"fmt"
	"time"

	"lucy/pkg/errors"

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
	Title     string `json:"title,omitempty"`
	Owner     string `json:"owner,omitempty"`
	WebrtcUrl string `json:"webrtc_link,omitempty"`
	RtmpUrl   string `json:"rtmp_link,omitempty"`
	FlvUrl    string `json:"flv_link,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	RecordUrl string `json:"record_url,omitempty"`
}

func GetLiveByStream(stream *Stream) (*Live, error) {
	if !stream.Publish.Active {
		return nil, errors.ErrStreamInactive
	}
	if !VerifyLiveName(stream.Url) {
		return nil, errors.ErrStreamFormatError
	}
	l := models.Live{}
	err := models.Db().Where(models.Live{Name: stream.Url}).First(&l).Error

	if err != nil {
		return nil, err
	} else {
		ret := Live{
			Id:        stream.Id,
			Name:      l.Name,
			Title:     l.Title,
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

func GetLiveById(liveId string) (*Live, error) {
	streams := GetStreams()
	for i, _ := range streams {
		if streams[i].Id == liveId {
			var l models.Live
			err := models.Db().Where(models.Live{Name: streams[i].Url}).First(&l).Error
			if err != nil {
				return nil, errors.ErrLiveNotFound
			}
			ret := Live{
				Id:        liveId,
				Name:      l.Name,
				Title:     l.Title,
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
	return nil, errors.ErrLiveNotFound
}

func GetLivesByUser(username string) ([]Live, error) {
	var ls []models.Live
	err := models.Db().Where(models.Live{Owner: username}).Find(&ls).Error
	var ret []Live
	if err != nil {
		log.Info("user have no live", "username", username)
		return ret, errors.ErrUserHaveNoLive
	}
	for i, _ := range ls {
		live := Live{
			Name:      ls[i].Name,
			Title:     ls[i].Title,
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
		return errors.ErrLiveFormatError
	}
	if owner, _, ok := ParseLiveName(liveName); !ok || owner != username {
		log.Debug("LiveRecord", "owner", owner, "username", username, "ok", ok)
		return errors.ErrRequesterNotOwner
	}

	var l models.Live
	err := models.Db().Where(models.Live{Name: liveName}).First(&l).Error
	if err != nil {
		log.Warn("Can not find live", "liveName", liveName)
		return err
	}
	if l.RecordStatus != noRecord {
		return errors.ErrRecordingStarted
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

func GenerateLive(username string, title string) (*Live, error) {
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
		Title: title,
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
		Name:      livaName,
		Owner:     username,
		Title:     title,
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

func ParseLiveName(liveName string) (owner, token string, ok bool) {
	if !VerifyLiveName(liveName) {
		return "", "", false
	}
	flag := 1
	var ownerBuf []byte
	var tokenBuf []byte
	for i := len(app) + 2; i < len(liveName); i++ {
		if liveName[i] == '/' {
			flag++
		} else {
			if flag == 1 {
				ownerBuf = append(ownerBuf, liveName[i])
			} else {
				tokenBuf = append(tokenBuf, liveName[i])
			}
		}
	}
	owner = string(ownerBuf)
	token = string(tokenBuf)
	ok = true
	return
}

func VerifyLiveName(liveName string) bool {
	if len(liveName) < liveTokenLength+len(app)+3 {
		return false
	}
	if liveName[0] != '/' || liveName[1:len(app)+1] != app {
		return false
	}
	if liveName[len(liveName)-liveTokenLength-1] != '/' {
		return false
	}
	count := 0
	for i := 0; i < len(liveName); i++ {
		if liveName[i] == '/' {
			count++
		}
	}
	return count == 3
}
