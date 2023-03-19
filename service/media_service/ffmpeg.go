package media_service

import (
	"fmt"
	"os"
	"os/exec"

	"lucy/models"
	"lucy/pkg/log"
	"lucy/pkg/setting"
)

// ffmpegRecord use ffmpeg to record live video
func ffmpegRecord(streamUrl string, liveName string) {
	go func() {
		owner := ParseUserFromLivePath(liveName)
		ownerDir := fmt.Sprintf("%s/live-record/%s", setting.AppSetting.AppRoot, owner)
		outFileName := func() string {
			i := len(liveName) - 1
			for liveName[i] != '/' {
				i--
			}
			return liveName[i+1:] + ".mp4"
		}()
		os.MkdirAll(ownerDir, os.ModePerm)
		outFilePath := fmt.Sprintf("%s/%s", ownerDir, outFileName)
		cmd := exec.Command("ffmpeg", "-i", streamUrl, outFilePath, "-y")
		err := cmd.Run()
		l := models.Live{}
		if err != nil {
			log.Error("ffmpeg record failed", "err", err.Error())
			err = models.Db().Model(&l).Where(models.Live{Name: liveName}).
				Update("record_status", noRecord).Error
			if err != nil {
				log.Debug("update database failed", "err", err)
			}
			return
		}
		err = models.Db().Model(&l).Where(models.Live{Name: liveName}).
			Updates(models.Live{
				RecordStatus: available,
				RecordPath:   outFilePath,
			}).Error
		if err != nil {
			log.Warn("update database failed", "err", err)
		} else {
			log.Info("record success", "streamUrl", streamUrl,
				"liveName", liveName, "recordPath", outFilePath)
		}
	}()
}
