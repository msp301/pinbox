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

// A Thread represents a Message chain stored in the Mailbox.
type Thread struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Subject    string    `json:"subject"`
	NewestDate int64     `json:"newestDate"`
	OldestDate int64     `json:"oldestDate"`
	Authors    []string  `json:"authors"`
	Messages   []Message `json:"messages"`
}
