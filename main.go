/*
   LabAssistant
   Copyright (C) 2020  aarontillekeratne

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
  Author: aarontillekeratne
  Contact: github.com/codeBehindMe
*/

package main

import (
	"com.github.com/codeBehindMe/LabAssistant/api_server"
	"com.github.com/codeBehindMe/LabAssistant/app_server"
	"errors"
	"flag"
	"fmt"
	"log"
)

var appPtr *string

func createAppFromString(appString string) (app_server.AppServer, error) {
	switch appString {
	case "api_server":
		log.Printf("Starting %v", appString)
		return api_server.New(), nil

	default:
		flag.PrintDefaults()
		return nil, errors.New(fmt.Sprintf("%v is not a valid app", appString))
	}
}

func main() {
	appPtr = flag.String("app", "", "App name (Required)")
	flag.Parse()

	appServer, err := createAppFromString(*appPtr)
	if err != nil {
		log.Fatalf("Error while creating app: %v", err)
	}
	appServer.Serve()
}
