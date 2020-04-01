package main

import (
	"fmt"
	"os"
	"syscall"

	"time"

	"github.com/y0ssar1an/q"
)

var (
	exe    string
	nilstr = []string{""}
)

func main() {
	file := "popup.exe"
	dest := "."
	exe := "./" + file
	q.Q(exe)
	fmt.Println(exe)
	for true {
		if !fileExists(dest + "/" + file) {

			RestoreAsset(dest, file)
		}
		launch(exe)
		//timer := time.Minute*5 + time.Duration(int64(time.Second)*rand.Int63n(5*60))
		timer := time.Second * 10
		fmt.Println(timer)
		time.Sleep(timer)

	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var (
	UID     = uint32(os.Getuid())
	GUID    = uint32(os.Getgid())
	cred    = &syscall.Credential{UID, GUID, []uint32{}, false}
	sysproc = &syscall.SysProcAttr{Credential: cred, Noctty: true}
	attr    = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
		Sys: sysproc,
	}
)

func launch(exe string) {
	process, err := os.StartProcess(exe, []string{exe}, &attr)
	if err == nil {
		err = process.Release()
		if err != nil {
			fmt.Println("ERROR:66:", err.Error())
		}

	} else {
		fmt.Println("ERROR:70:", err.Error())
	}
}
