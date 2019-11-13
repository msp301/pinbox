# Pinbox

Pinbox is a self-hosted webmail client greatly inspired by Google Inbox.

*Warning: Currently under development*

## Prerequisites

* Docker (tested on Debian 10 with v18.09.1)
* Go (https://golang.org/)
* notmuch (https://notmuchmail.org/)

## Getting started

### Hosting your email

Pinbox requires local access to a [Maildir](https://en.wikipedia.org/wiki/Maildir) directory. This can be setup with [OfflineIMAP](http://www.offlineimap.org/) or similar.

[Notmuch](https://notmuchmail.org/) mail indexer is used to provide fast email access and managing mailbox labels.
Once you have `notmuch` installed on your system, run `notmuch setup && notmuch new <MAILDIR DIRECTORY>` to initialise the `notmuch` database ready for Pinbox.

### Starting Pinbox

Pinbox consists of a single page app to provide the web interface and [a server](https://github.com/msp301/pinbox-server) for handling access to your email.

1. Clone 'pinbox' and '[pinbox-server](https://github.com/msp301/pinbox-server)'
2. Start `pinbox-server` with `go build && ./pinbox-server <CONFIG FILE>`
3. Build `pinbox` client with `make`
4. Start `pinbox` client with `make start`
5. Navigate to `http://localhost:4200`

The Angular config files `proxy.conf.json` and `server.conf` point API requests to `pinbox-server` without hard coding the hostname into the client code. The setup expects `pinbox-server` and `pinbox` to be available on the same host.

## License

Licensed under GPL-3.0
