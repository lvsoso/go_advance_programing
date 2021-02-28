package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

var myApp fyne.App

func init() {
	//获取中文字体列表
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//设置字体
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}


func main() {
	myApp = app.New()
	myWindow := myApp.NewWindow("Chat Demo")

	auth := NewAuth(myApp, myWindow)
	myWindow.SetContent(auth.LoginForm())
	myWindow.ShowAndRun()
}