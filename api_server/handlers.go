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

import (
	"com.github.com/codeBehindMe/LabAssistant/lab"
	"com.github.com/codeBehindMe/LabAssistant/utils"
	"fmt"
	"log"
	"net/http"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Application: Api Server")
}

func version(w http.ResponseWriter, r *http.Request) {
	version, err := utils.ReadFileToString("VERSION")
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	_, _ = fmt.Fprint(w, version)
}

// newLab creates a new Lab.
func newLab(w http.ResponseWriter, r *http.Request) {
	newLab := &lab.Lab{}
	err := utils.LoadJsonToStruct(r.Body, newLab)
	if err != nil {
		log.Fatalf("Error occured whilst extracting payload: %v", err)
	}
}

func (s *ApiServer) NewLab(w http.ResponseWriter, r *http.Request) {
	newLab := &lab.Lab{}
	if err := utils.LoadJsonToStruct(r.Body, newLab); err != nil {
		log.Fatalf("Error occured whilt extracting payload: %v", err)
		return
	}
}
