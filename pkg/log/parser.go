package log

import (
	"fmt"
	"runtime"
	"time"
)

func iToA(buf *[]byte, i int, wid int) {
	var b [20]byte
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

func serialize(callDepth int, level int, msg string, v ...any) string {
	buf := make([]byte, 0)
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	}
	{
		buf = append(buf, '[')
		buf = append(buf, logLevelStr[level]...)
		buf = append(buf, ']')
	}
	{
		iToA(&buf, year, 4)
		buf = append(buf, '/')
		iToA(&buf, int(month), 2)
		buf = append(buf, '/')
		iToA(&buf, day, 2)
		buf = append(buf, ' ')
		iToA(&buf, hour, 2)
		buf = append(buf, ':')
		iToA(&buf, min, 2)
		buf = append(buf, ':')
		iToA(&buf, sec, 2)
		buf = append(buf, '.')
		iToA(&buf, now.Nanosecond()/1e6, 3)
		buf = append(buf, ' ')
	}
	{
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		buf = append(buf, short...)
		buf = append(buf, ':')
		iToA(&buf, line, -1)
		buf = append(buf, ": "...)
	}
	{
		buf = append(buf, msg...)
	}
	var a []any
	for i := 0; i < len(v); i += 2 {
		if i+1 < len(v) {
			buf = append(buf, ", "...)
			buf = append(buf, v[i].(string)...)
			buf = append(buf, ": %+v"...)
			a = append(a, v[i+1])
		} else {
			buf = append(buf, v[i].(string)...)
		}
	}
	return fmt.Sprintf(string(buf), a...)
}
