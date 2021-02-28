package main

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

const (
	LoginStatus = iota
	RegisterStatus
)

var (
	genderVal string
)

type Auth struct {
	app fyne.App
	window fyne.Window
	status int
	loginUrl string
	registerUrl string
}

func NewAuth(app fyne.App, window fyne.Window) *Auth {
	return &Auth{
		app: app,
		window:window,
		status:LoginStatus,
		loginUrl: serverURL+"/user/login",
		registerUrl: serverURL+"/user/register",
	}
}

func (auth *Auth) ShowRegisterSuccess()  {
	dialog.NewInformation(
		"","register success !",
		auth.window,
		).Show()
}

func (auth *Auth) ToHome()  {

}

func (auth *Auth) Toggle()  {
	switch auth.status {
	case LoginStatus:
		auth.status = RegisterStatus
		auth.window.SetContent(auth.RegisterForm())
	case RegisterStatus:
		auth.status = LoginStatus
		auth.window.SetContent(auth.LoginForm())
	}
}

func (auth *Auth) LoginForm() *widget.Form {
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()

	login := widget.NewButton("Login", func() {
		auth.Login(username.Text, password.Text)
	})

	register := widget.NewButton("To Register", func() {
		auth.Toggle()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username", Widget: username},
			{Text: "Password", Widget: password},
			&widget.FormItem{Widget: login},
			&widget.FormItem{Widget: register},
		},
	}

	return form
}

func (auth *Auth) RegisterForm() *widget.Form {
	username := widget.NewEntry()
	nickname := widget.NewEntry()
	gender := widget.NewSelect([]string{"male", "female"}, func(value string) {
		genderVal = value
	})
	mobile := widget.NewEntry()
	password := widget.NewPasswordEntry()
	repassword := widget.NewPasswordEntry()

	register := widget.NewButton("register", func() {
		auth.Register(username.Text, nickname.Text, genderVal, mobile.Text, password.Text)
	})
	login := widget.NewButton("To Login", func() {
		auth.Toggle()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Username", Widget: username},
			{Text: "Nickname", Widget: nickname},
			{Text: "Gender", Widget: gender},
			{Text: "Mobile", Widget: mobile},
			{Text: "Password", Widget: password},
			{Text: "Password Confirm", Widget: repassword},
			&widget.FormItem{Widget: register},
			&widget.FormItem{Widget: login},
		},
	}

	return form
}

func (auth *Auth) Register(username, nickname, gender, mobile, password string) {
	registerReq := RegisterReq{
		Username: username,
		Nickname: nickname,
		Gender:   gender,
		Mobile:   mobile,
		Password: password,
	}

	data, err := json.Marshal(registerReq)
	if err != nil {
		log.Fatal(err)
	}
	_, err = Post(auth.registerUrl, data)
	if err != nil && err.Error()!="EOF" {
		log.Fatal(err)
	}
	auth.ShowRegisterSuccess()
	auth.Toggle()
}

func (auth *Auth) Login(username, password string) {
	loginReq := LoginReq{
		Username: username,
		Password: password,
	}
	data, err := json.Marshal(loginReq)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := Post(auth.registerUrl, data)
	if err != nil && err.Error()!="EOF" {
		log.Fatal(err)
	}

	userReply := new(UserReply)
	err = json.Unmarshal(*resp, userReply)
	if err != nil{
		log.Fatal(err)
	}

	auth.Toggle()
}