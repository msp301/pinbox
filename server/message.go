// Copyright 2019-2022 Martin Pritchard
//
// This file is part of Pinbox.
//
// Pinbox is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
//
// Pinbox is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public
// License along with Pinbox.  If not, see <https://www.gnu.org/licenses/>.

package pinbox

// A Message represents an individual email.
type Message struct {
	ID     string   `json:"id"`
	Epoch  int64    `json:"epoch"`
	Author string   `json:"author"`
	Files  []string `json:"files"`
}

// MessageContent is a container for the body of a Message.
type MessageContent struct {
	ID      string `json:"id"`
	Epoch   int64  `json:"epoch"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
