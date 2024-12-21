package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func LoginWidget(manager ScreenMan) fyne.CanvasObject {
	//function
	var login = func()(bool,error){
		return true,nil
	}

	//email
	emailInput := widget.NewEntry()
	emailInput.SetPlaceHolder("Email")
	emailAt := container.New(layout.NewHBoxLayout(),canvas.NewText("@",color.White))
	emailDomain := widget.NewSelectEntry([]string{"gmail.com","ymail.com"})
	emailDomain.SetPlaceHolder("Domain")

	//password
	passwordInput := widget.NewPasswordEntry()
	passwordInput.SetPlaceHolder("Password")
	
	//remember-me checkbox
	rememberMeCheckbox:= widget.NewCheck("Remember me", func(val bool){})
	rememberMeCheckbox.Checked = true

	//forgot-pswd button
	forgotPasswordButton:= widget.NewHyperlinkWithStyle("Forgot password?",nil,fyne.TextAlignLeading,fyne.TextStyle{})
	forgotPasswordButton.OnTapped = func() {
		fmt.Println("!")
	}

	//register-new button


	//login button
	loginButton := widget.NewButton("Login",func(){
		_, err := login()
		if err != nil {
			panic(err)
		} else {
			manager.SwitchScreen("home")
		}
	})

	//composed layout
	fixedInput:= container.New(layout.NewGridWrapLayout(fyne.NewSize(200,36)),emailInput)
	fixedDomain:=container.New(layout.NewGridWrapLayout(fyne.NewSize(200,36)),emailDomain)
	emailComposed := container.New(layout.NewHBoxLayout(),fixedInput,emailAt,fixedDomain,)
	composed := container.New(layout.NewVBoxLayout(),emailComposed,passwordInput,rememberMeCheckbox)
	centered := container.New(layout.NewVBoxLayout(),layout.NewSpacer(),composed,loginButton,forgotPasswordButton,layout.NewSpacer(),)
	ctx:= container.New(layout.NewHBoxLayout(),layout.NewSpacer(),centered,layout.NewSpacer())

	return ctx
}