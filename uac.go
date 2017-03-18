/*
    Hardentools
    Copyright (C) 2017  Claudio Guarnieri, Mariano Graziano

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
    "golang.org/x/sys/windows/registry"
)

func trigger_uac(enable bool) {
    key, _ := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", registry.WRITE)

    if enable {
        events.AppendText("Setting UAC to not ask for changing Windows settings (default)\n")
        err := key.SetDWordValue("ConsentPromptBehaviorAdmin", 5)
        if err != nil {
            events.AppendText("!! Setting UAC failed.\n")
        }
    } else {
        events.AppendText("Setting UAC to always ask on secure desktop (secure)\n")
        err := key.SetDWordValue("ConsentPromptBehaviorAdmin", 2)
        if err != nil {
            events.AppendText("!! Setting UAC failed.\n")
        }
    }

    key.Close()
}