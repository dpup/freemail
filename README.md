# go freemail

A database of free and [disposable](http://en.wikipedia.org/wiki/Disposable_email_address) email domains and a handy Go module for querying it.

Use the Go module or access the files in the `./data` directory and parse with your language of choice. For Node.js see the [original implementation](https://github.com/willwhite/freemail).

## Database

There are three key data files in this project:

- [free.txt](https://github.com/dpup/freemail/blob/master/data/free.txt) contains a list of domains that are known to provide free email service.
- [disposable.txt](https://github.com/dpup/freemail/blob/master/data/disposable.txt) contains a list of domains that are known to provide disposable email service.
- [spammy.txt](https://github.com/dpup/freemail/blob/master/data/spammy.txt) contains a list of domains that Stop Forum Spam report as toxic and abusive.

Domains may only be a member of one list.

## Updating the database

Run `./update` to pull in the latest domains from the free sources listed in `sources.txt`. These domains will be placed in `free.txt`. Disposable emails from [ivolo/disposable-email-domains](https://github.com/ivolo/disposable-email-domains) will be added to `disposable.txt`. Spammy emails from [Stop Forum Spam](https://www.stopforumspam.com/downloads) will be added to `spammy.txt`.

Then run `go generate`.

## Go

```
go get github.com/dpup/freemail
```

```go
import "github.com/dpup/freemail"

freemail.IsFree("gmail.com")
> true
freemail.IsFree("mailinater.com")
> true
freemail.IsDisposable("gmail.com")
> false
freemail.IsDisposable("mailinater.com")
> true
```

