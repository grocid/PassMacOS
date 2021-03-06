/*
Copyright (c) 2018 Carl Löndahl. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Pass Desktop nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
    "github.com/murlokswarm/app"
    _ "github.com/murlokswarm/mac"
    "log"

    "pass/util"
)

type (
    Application struct {
        Locked bool
        Icons  map[string]bool
    }
)

var (
    win  app.Contexter
    pass Application
)

const (
    MinimumPasswordLength = 8
)

var config util.Configuration

func GetImageName(filename string) string {
    if !pass.Icons[filename] {
        filename = "default"
    }
    return filename
}

func main() {
    // Pass is locked by default.
    pass.Locked = true

    // Load the config.
    config, _ = util.GetConfig(app.Resources())

    log.Println(config)

    pass.Icons = util.ListAvailableIcons(app.Resources())

    log.Println(pass.Icons)

    app.OnLaunch = func() {
        // Creates the AppMainMenu component.
        appMenu := &AppMainMenu{}

        // Mounts the AppMainMenu component into the application menu bar.
        if menuBar, ok := app.MenuBar(); ok {
            menuBar.Mount(appMenu)
        }

        // Create the main window
        win = newMainWindow()
    }

    app.OnReopen = func() {
        if win != nil {
            return
        }
        win = newMainWindow()
    }

    app.Run()
}

func newMainWindow() app.Contexter {
    // Creates a window context.
    win := app.NewWindow(app.Window{
        Title:          "Pass",
        Width:          300,
        Height:         548,
        Vibrancy:       app.VibeDark,
        TitlebarHidden: true,
        OnClose: func() bool {
            win = nil
            return true
        },
    })

    log.Println("NewWindow")

    // Create component...
    ps := &UnlockScreen{}

    // ...and mount to window
    win.Mount(ps)

    // Return to context
    return win
}
