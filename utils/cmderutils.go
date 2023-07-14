package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"unsafe"
)

/**
执行系统脚本工具类
参照这个工具类，写的很棒 https://github.com/andeya/goutil/blob/master/cmder/cmd.go
*/

var cmdArg = make([]string, 2)

// 根据操作系统环境初始化cmd关闭命令行的参数
func init() {
	if runtime.GOOS == "windows" {
		cmdArg[0] = "cmd"
		cmdArg[1] = "/c"
	} else {
		cmdArg[0] = "/bin/sh"
		cmdArg[1] = "-c"
	}
}

type Result struct {
	buf bytes.Buffer
	err error
	str *string
}

func RunCmd(cmdLine string, timeout ...time.Duration) *Result {
	cmd := exec.Command(cmdArg[0], cmdArg[1], cmdLine)

	var re = new(Result)
	cmd.Stdout = &re.buf
	cmd.Stderr = &re.buf
	cmd.Env = os.Environ()
	re.err = cmd.Start()
	if re.err != nil {
		return re
	}
	if len(timeout) == 0 || timeout[0] <= 0 {
		re.err = cmd.Wait()
		return re
	}
	timer := time.NewTimer(timeout[0])
	done := make(chan error)
	//启动协程去执行命令
	go func() { done <- cmd.Wait() }()
	select {
	case re.err = <-done:
		//正常执行完，或者出现异常
		timer.Stop()
	case <-timer.C:
		//时间到期
		if err := cmd.Process.Kill(); err != nil {
			re.err = fmt.Errorf("command timed out and killing process fail: %s", err.Error())
		} else {
			// wait for the command to return after killing it
			<-done
			re.err = errors.New("command timed out")
		}
	}
	return re
}

func (r *Result) Err() error {
	return r.err
}

// String returns the exec log.
func (r *Result) String() string {
	if r.str == nil {
		b := bytes.TrimSpace(r.buf.Bytes())
		r.str = (*string)(unsafe.Pointer(&b))
	}
	return *r.str
}
