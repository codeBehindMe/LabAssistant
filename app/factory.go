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

package app

import (
	"com.github.com/codeBehindMe/LabAssistant/api_server"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func newHttpServer(r *mux.Router) *http.Server {
	return &http.Server{
		Handler: r,
	}
}

func NewApp(appName string) (*http.Server, error) {
	switch appName {
	case "api_server":
		return newHttpServer(api_server.GetApiServerRouter()), nil
	}
	return nil, errors.New("not a valid application")
}
