package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

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
		timer := time.Minute*5 + time.Duration(int64(time.Second)*rand.Int63n(5*60))
		//timer := time.Second * 10
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

func launch(exe string) {
	cmd := exec.Command(exe)
	cmd.Stdout = os.Stdout
	cmd.Start()
}
