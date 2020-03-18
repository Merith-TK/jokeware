package main

import "github.com/zserge/webview"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Annoying You")
	w.SetSize(100, 50, webview.HintNone)
	w.Navigate("file:///home/merith/Workspace/jokeware/popup/popup.html")
	w.Run()
}
