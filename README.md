# Pinbox

Pinbox is a self-hosted webmail client greatly inspired by Google Inbox.

*Warning: Pinbox is in early stages of development and lacks most functionality expected in a webmail client -- Contributions welcome :-)*

## Prerequisites

* Docker and `docker-compose`
* notmuch (https://notmuchmail.org/)

## Getting started

### Hosting your email

Pinbox requires local access to a [Maildir](https://en.wikipedia.org/wiki/Maildir) directory. A mailbox can be synchronized to a local Maildir using [OfflineIMAP](http://www.offlineimap.org/) or similar. If you wish to try Pinbox with an exported mailbox and have a `.mbox` file, this file can be converted to a Maildir directory using [mb2md](https://github.com/dovecot/tools/blob/main/mb2md.pl).

### Before starting Pinbox

Pinbox uses the [Notmuch](https://notmuchmail.org/) mail indexer to provide fast email access and managing mailbox labels.

Before Pinbox is able to read your Maildir directory `notmuch` needs to be installed on your system. Run `notmuch setup && notmuch new <MAILDIR DIRECTORY>` to initialise the `notmuch` database ready for Pinbox.

### Starting Pinbox

Pinbox consists of a single page app to provide the web interface and a server component for handling access to your email.

1. Edit `docker-compose.yml` to insert the full path of your `maildir` directory
2. Start all services using: `docker-compose up -d`
3. Navigate to `http://localhost:4200`

## License

Licensed under AGPL-3.0-or-later
