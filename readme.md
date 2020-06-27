# go-mail-test

## Gmail Setting
1. setup IMAP

    [Gmail Help](https://support.google.com/mail/answer/7126229)

2. turn on "Less secure app access"

    [Google Account Help](https://support.google.com/accounts/answer/6010255)

## Project Setting

```shell
$ cp .envrc.example .envrc
$ direnv edit . # set your Gmail account infomations
$ direnv allow
```

## Run

```shell
$ go run index.go
```
