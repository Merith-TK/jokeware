package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	. "github.com/y0ssar1an/q"
)

var (
	nilstr = []string{""}
)

func main() {
	file := "popup"
	dest := "./"
	exe := dest + "/" + file
	Q(exe)
	fmt.Println(dest)
	for true {
		if !fileExists(dest + "/" + file) {

			RestoreAsset(dest, file)
		}
		//cmd := exec.Command(exe)
		//cmd.Stdout = os.Stdout
		//cmd.Start()
		time.Sleep(5 * time.Second)
		err := syscall.Exec(exe, nilstr, nilstr)
		if err != nil {
			fmt.Println("PRINT: ", err)
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
