package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func init() {
	osCheck()
	directoryExists(configPath)
	directoryExists(keysPath)

	if fileExists(buddyListPath) {
		GlobalBuddyList, _, _ = loadBuddylist()
	} else {
		createFile(buddyListPath)
		addFriendFromEntry("donuthandler@3ck0.com::b19d495c2dcf9f92b0b420ba4d2a975d44df9b70014930f3b883b8ddbd253cc4")
		GlobalBuddyList, _, _ = loadBuddylist()
	}

	// if fileExists(requestListPath) {
	// 	GlobalRequestList, _, _ = loadRequestlist()
	// } else {
	// 	createFile(requestListPath)
	// 	GlobalRequestList, _, _ = loadRequestlist()
	// }

}

func main() {
	// keyTest()

	initKeys()

	a := app.New()

	w := a.NewWindow("buddylist")
	_, names, _ := loadBuddylist()

	usercount := strconv.Itoa(len(names))

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Chat Client",
			fyne.NewMenuItem("Chattter 1.0.0", func() {
			}),
			fyne.NewMenuItem(usercount+" Friends", func() {
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Show Buddy List", func() {
				w.Show()
			}),
			fyne.NewMenuItem("Hide Buddy List", func() {
				w.Hide()
			}))
		desk.SetSystemTrayMenu(m)
	}

	list := widget.NewList(
		func() int {
			return len(names)
		},
		func() fyne.CanvasObject {
			return widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {

			loadBuddylist()
			o.(*widget.Label).SetText(names[i])

		})

	yourFriendsLabel := widget.NewLabel("Friends")
	yourFriendsLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Monospace: false, Symbol: false, TabWidth: 2}

	versionText := widget.NewLabel("v0.1")

	mainWindowBottomButtonBar := container.New(
		layout.NewHBoxLayout(),
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			wAddFriend := a.NewWindow("add a friend")
			wAddFriend.Resize(fyne.NewSize(640, 140))
			entry := widget.NewEntry()

			form := &widget.Form{
				Items: []*widget.FormItem{
					{Text: "Name ", Widget: entry}},
				OnSubmit: func() {
					if !addRequestFromEntry(entry.Text) {
						log.Println("Something went wrong")
					}
					list.Refresh()
					wAddFriend.Hide()
					entry.Text = ""
				},
				SubmitText: "Add Friend",
			}

			// we can also append items
			addFriendContent := container.New(layout.NewVBoxLayout(), widget.NewLabel("Enter your friend's full username in the bow below, then press the Add Friend button."), widget.NewLabel("Example:\tdonuthandler@3ck0.com"), form)

			wAddFriend.SetContent(addFriendContent)
			wAddFriend.CenterOnScreen()
			wAddFriend.Show()
		}),

		layout.NewSpacer(),
		versionText,
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
			wSettings := a.NewWindow("settings")
			wSettings.Resize(fyne.NewSize(760, 420))
			settingsContent := container.New(layout.NewHBoxLayout(), widget.NewLabel("this"), widget.NewLabel("that"))

			wSettings.SetContent(settingsContent)
			wSettings.CenterOnScreen()
			wSettings.Show()
		}),
	)

	w.Resize(fyne.NewSize(240, 480))
	w.SetContent(container.NewBorder(nil, mainWindowBottomButtonBar, nil, nil, list))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}
