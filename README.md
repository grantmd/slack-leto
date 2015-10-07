slack-leto
=======

A bot for [Slack](https://slack.com) written in Go (golang) that randomly once per days interrupts, Leto-style, to let you know its thoughts on Slack.

Usage
-----

* Build the code with `go build`

* Start the bot with `./slack-leto` on an internet-accessible server. (Check the output of `./slack-leto -h` for configuration options)

* Configure an [Outgoing Webhook](https://my.slack.com/services/new/outgoing-webhook) in your Slack and point it to the place where your bot is running. For example: `http://example.com:8002/`

