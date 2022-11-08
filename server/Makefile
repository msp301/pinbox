# Copyright 2019 Martin Pritchard
#
# This file is part of Pinbox.
#
# Pinbox is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Pinbox is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Pinbox.  If not, see <https://www.gnu.org/licenses/>.

BINARY=pinbox-server

.PHONY: build clean test

$(BINARY): build test

build:
	go get -d ./...
	go build -o $(BINARY) cmd/pinbox-server/main.go

clean:
	rm -f $(BINARY)

test:
	go test -v -race -cover

