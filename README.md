# Pinbox

Pinbox is a self-hosted webmail client greatly inspired by Google Inbox.

*Warning: Currently under development*

## Prerequisites

* Angular >= 7.3.3
* Go (https://golang.org/)
* notmuch (https://notmuchmail.org/)

## Getting started

### Hosting your email

Pinbox requires local access to a [Maildir](https://en.wikipedia.org/wiki/Maildir) directory. This can be setup with [OfflineIMAP](http://www.offlineimap.org/) or similar.

### Starting Pinbox

Pinbox consists of a single page app to provide the web interface and [a server](https://github.com/msp301/pinbox-server) for handling access to your email.

1. Clone 'pinbox' and '[pinbox-server](https://github.com/msp301/pinbox-server)'
2. Start `pinbox-server` with `go build && ./pinbox-server <MAILDIR DIRECTORY>`
3. Start `pinbox` client with `ng serve --proxy-config proxy.conf.json`
4. Navigate to `http://localhost:4200`

The proxy config file `proxy.conf.json` is used by Angular to point API requests to `pinbox-server` without hard coding the hostname into the client code. By default this IS setup to use `localhost`.

## License

Licensed under GPL-3.0