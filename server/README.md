# Pinbox webmail backend

Pinbox is a self-hosted webmail client greatly inspired by Google Inbox.

*Warning: Currently under development*

## Prerequsites

* Go (https://golang.org/)
* notmuch (https://notmuchmail.org/)

## Configuration

Pinbox server requires a [TOML](https://github.com/toml-lang/toml) configuration file to define the Maildir location, port etc. See [example/config.toml](https://github.com/msp301/pinbox-server/blob/master/example/config.toml).

## Starting the server

`go build && ./pinbox-server config.toml`

## Architecture

Client App -> Backend API -> Notmuch -> OfflineIMAP

## License

Licensed under GPL-3.0
