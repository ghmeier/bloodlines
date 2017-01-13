# bloodlines
[![Build Status](https://travis-ci.org/ghmeier/bloodlines.svg?branch=master)](https://travis-ci.org/ghmeier/bloodlines)
[![Coverage Status](https://coveralls.io/repos/github/ghmeier/bloodlines/badge.svg?branch=master)](https://coveralls.io/github/ghmeier/bloodlines?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghmeier/bloodlines)](https://goreportcard.com/report/github.com/ghmeier/bloodlines)

A go service for sending Expresso emails

## Quick Start
Follow this guide if you just want to run an instance of `bloodlines`.

You'll need [docker installed](https://docs.docker.com/engine/installation/)

In `exampleConfig.json`, modify mysql host, port, user, and password to reference your sql instance. Do the same for rabbitmq and sendgrid information.

```bash
$ docker pull ghmeier/bloodlines
$ docker run -p 8080:8080 ghmeier/bloodlines:latest
```

## Development Setup
Follow this guide if you want to develop on bloodlines.

[Install Golang](https://golang.org/doc/install)
[Install Docker](https://docs.docker.com/engine/installation/)

```bash
$ go get github.com/ghmeier/bloodlines
$ cd $GOPATH/src/github.com/ghmeier/bloodlines
$ go get github.com/tools/godep
$ go get github.com/stretchr/testify
$ godep restore
$ make deps
$ make test
```

After following this, you should have a `bloodlines` binary in the root directory. Running `make run` executes that binary, and your output should indicate that all tests are passing.

Now, you need a few other services to properly run bloodlines. We'll use docker to get them up and running.
```bash
$ docker run -d -p 3306:3306 -e MYSQL_PASS="soomepassword" ghmeier/expresso-mysql
$ docker run -d -p 5672:5672 ghmeier/rabbitmq-delayed
```

Modify `exampleConfig.json` to reflect the ip and port of those two instances. (Try `docker ps` if you're unsure).

```bash
$ mv exampleConfig.json config.json
```

Finally, if you want to actually send emails, you'll have to fill out the sendgrid configuration with your own [account](https://sendgrid.com/).

Now start bloodlines:
```
$ make run
```

## API
__All api methods should return `{"success":true}`__

### Content
Object
```javascript
{
	"id": "24073c8c-119e-4d7d-836e-4fff2db98549",
	"contentType": "EMAIL",
	"text": "Hello!",
	"parameters": [],
	"active": true
}
```
* `POST /api/content`
* `GET /api/content`
* `GET /api/content/:contentId`
* `PUT /api/content/:contentId`
* `DELETE /api/content/:contentId`

### Receipt
* `GET /api/receipt`
* `POST /api/receipt/send`
* `GET /api/receipt/:receiptId`

### Job
* `GET /api/job`
* `POST /api/job`
* `GET /api/job/:jobId`
* `PUT /api/job/:jobId`
* `DELETE /api/job/:jobId`

### Trigger
* `POST /api/trigger`
* `GET /api/trigger`
* `GET /api/trigger/:triggerKey`
* `PUT /api/trigger/:triggerKey`
* `DELETE /api/trigger/:triggerKey`
* `POST /api/trigger/:triggerKey/activate`

### Preference
* `POST /api/preference`
* `GET /api/preference/:userId`
* `PATCH /api/preference/:userId`
* `DELETE /api/preference/:userId`
