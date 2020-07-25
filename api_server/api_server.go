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

package api_server

// API Server is the main service of Lab Assistant. The API server stands as
// the control plane for Lab Assistant.

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ApiServer struct {
	addr   string
	trackerAddr string
	router *mux.Router
}

func New() ApiServer {
	server := ApiServer{}

	server.router = mux.NewRouter()
	server.router.HandleFunc("/", baseHandler)
	server.router.HandleFunc("/version", version)
	server.router.HandleFunc("/newlab", server.NewLab)

	return server
}

func (s ApiServer) Serve() {
	httpServer := http.Server{Handler: s.router}
	log.Fatal(httpServer.ListenAndServe())
}
