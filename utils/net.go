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

package utils

import (
	"encoding/json"
	"io"
)

// FIXME: Requires thoughtful test or function redefinition.
// LoadJsonToStruct loads json data into a structure. Pointer to a structure
// should be passed in as parameter t.
func LoadJsonToStruct(r io.ReadCloser, t interface{}) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(t)
}