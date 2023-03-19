package media_service

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"lucy/pkg/log"

	"lucy/pkg/setting"
)

var httpApiPath string

func monitorLog(stdout *bufio.Reader, stderr *bufio.Reader) {
	logFilePath := "./srs/srs.log"
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Info("Srs log file open failed", "err", err)
		return
	}
	defer logFile.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	getOutput := func(r *bufio.Reader) {
		outputBytes := make([]byte, 512)
		for {
			n, err := r.Read(outputBytes)
			if err != nil {
				if err == io.EOF {
					break
				}
				e := fmt.Sprintf("srs monitorLog error reading stdout: %s", err.Error())
				io.WriteString(logFile, e)
			} else {
				info := string(outputBytes[:n])
				io.WriteString(logFile, info)
			}
		}
		wg.Done()
		return
	}

	go getOutput(stdout)
	go getOutput(stderr)
	wg.Wait()
}

func init() {
	httpApiPath = fmt.Sprintf("http://%s:%s",
		setting.SrsSetting.Ip,
		setting.SrsSetting.HttpApiPort,
	)
	if !setting.SrsSetting.Run {
		return
	}
	cmd := exec.Command("bash", "./srs/run.sh",
		setting.SrsSetting.RtmpPort,
		setting.SrsSetting.NginxHttpPort,
		setting.SrsSetting.NginxHttpsPort,
		setting.SrsSetting.HttpApiPort,
		setting.SrsSetting.RtcServerPort)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Srs stdout pipe open failed", "err", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Srs stderr pipe open failed", "err", err)
	}

	stdoutReader := bufio.NewReader(stdout)
	stderrReader := bufio.NewReader(stderr)

	if err = cmd.Start(); err != nil {
		log.Fatal("Srs start failed", "err", err)
	}

	go monitorLog(stdoutReader, stderrReader)
}
