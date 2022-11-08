// Copyright 2019 Martin Pritchard
//
// This file is part of Pinbox.
//
// Pinbox is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Pinbox is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Pinbox.  If not, see <https://www.gnu.org/licenses/>.

package pinbox

// A Label is a tag that can be used to organise messages in the Mailbox.
type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
