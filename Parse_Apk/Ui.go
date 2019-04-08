package main

import (
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	label := theme.CreateLabel()
	label.SetText("Are you sure?")

	yesButton := theme.CreateButton()
	yesButton.SetText("Yes")
	yesButton.OnClick(func(ev gxui.MouseEvent) {
		fmt.Println("Yes")
	})
	noButton := theme.CreateButton()
	noButton.SetText("No")
	noButton.OnClick(func(ev gxui.MouseEvent) {
		fmt.Println("No")
	})

	layout := theme.CreateLinearLayout()
	layout.AddChild(label)

	btnLayout := theme.CreateLinearLayout()
	btnLayout.AddChild(yesButton)
	btnLayout.AddChild(noButton)
	//btnLayout.SetOrientation(gxui.Horizontal)

	layout.AddChild(btnLayout)
	layout.SetHorizontalAlignment(gxui.AlignCenter)
	layout.SetVerticalAlignment(gxui.AlignMiddle)

	window := theme.CreateWindow(120, 60, "Message Box")
	window.AddChild(layout)
	window.OnClose(driver.Terminate)

	//gxui.EventLoop(driver)
}

func main() {
	gl.StartDriver(appMain)
}
