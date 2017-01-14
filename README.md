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

### Content

**Object**
```javascript
{
	"id":"dd82cc65-d79d-11e6-9d4c-0242ac120004",
	"contentType":"EMAIL",
	"text":"$greeting$ $first_name$,\n Welcome to Expresso.",
	"parameters":["first_name","greeting"],
	"status":"ACTIVE",
	"subject":""
}
```

* `POST /api/content` creates and adds a new content record to the database.

Example:

*Request:*
```
POST localhost:8080/api/content
{
	"contentType": "EMAIL",
	"text": "$greeting$ $first_name$,\n Welcome to Expresso.",
	"parameters": ["first_name","greeting"]
}
```

*Response:*
```
{
  "data": {
    "id": "86c3d82d-da86-11e6-9d4c-0242ac120004",
    "contentType": "EMAIL",
    "text": "Hey $first_name$,\n Welcome to Expresso.",
    "parameters": [
      "first_name"
    ],
    "status": "ACTIVE",
    "subject": ""
  }
}
```

* `GET /api/content?offset=0&limit=20` returns up to `limit` content records starting from `offset` when ordered by contentId

Example:

*Request:*
```
GET localhost:8080/api/content?offset=0&limit=20
```

*Response:*
```
{
  "data": [
    {
      "id": "86c3d82d-da86-11e6-9d4c-0242ac120004",
      "contentType": "EMAIL",
      "text": "Hey $first_name$,\n Welcome to Expresso.",
      "parameters": [
        "first_name"
      ],
      "status": "ACTIVE",
      "subject": ""
    }
  ]
}
```

* `GET /api/content/:contentId` returns the content record with the given contentID

Example:

*Request:*
```
GET localhost:8080/api/content/86c3d82d-da86-11e6-9d4c-0242ac120004
```

*Response:*
```
{
  "data": {
    "id": "86c3d82d-da86-11e6-9d4c-0242ac120004",
    "contentType": "EMAIL",
    "text": "Hey $first_name$,\n Welcome to Expresso.",
    "parameters": [
      "first_name"
    ],
    "status": "ACTIVE",
    "subject": ""
  }
}
```

* `PUT /api/content/:contentId`

Example:

*Request:*
```
PUT localhost:8080/api/content/86c3d82d-da86-11e6-9d4c-0242ac120004
{
    "id": "86c3d82d-da86-11e6-9d4c-0242ac120004",
    "contentType": "EMAIL",
    "text": "Hey $first_name$,\n Welcome to Expresso.",
    "parameters": [
      "first_name"
    ],
    "status": "ACTIVE",
    "subject": "Hello world"
}
```

*Response:*
```
{
  "data": {
    "id": "86c3d82d-da86-11e6-9d4c-0242ac120004",
    "contentType": "EMAIL",
    "text": "Hey $first_name$,\n Welcome to Expresso.",
    "parameters": [
      "first_name"
    ],
    "status": "ACTIVE",
    "subject": "Hello world"
  }
}
```

* `DELETE /api/content/:contentId`

Example:

*Request:*
```
DELETE localhost:8080/api/content/86c3d82d-da86-11e6-9d4c-0242ac120004
```

*Response:*
```
{
  "success": true
}
```

### Receipt
* `POST /api/receipt/send`

Example:

*Request:*
```

```

*Response:*
```

```

* `GET /api/receipt`

Example:

*Request:*
```

```

*Response:*
```

```

* `GET /api/receipt/:receiptId`

Example:

*Request:*
```

```

*Response:*
```

```

### Job
* `GET /api/job`

Example:

*Request:*
```

```

*Response:*
```

```

* `POST /api/job`

Example:

*Request:*
```

```

*Response:*
```

```

* `GET /api/job/:jobId`

Example:

*Request:*
```

```

*Response:*
```

```

* `PUT /api/job/:jobId`

Example:

*Request:*
```

```

*Response:*
```

```

* `DELETE /api/job/:jobId`

Example:

*Request:*
```

```

*Response:*
```

```

### Trigger
* `POST /api/trigger`

Example:

*Request:*
```

```

*Response:*
```

```

* `GET /api/trigger`

Example:

*Request:*
```

```

*Response:*
```

```

* `GET /api/trigger/:triggerKey`

Example:

*Request:*
```

```

*Response:*
```

```

* `PUT /api/trigger/:triggerKey`

Example:

*Request:*
```

```

*Response:*
```

```

* `DELETE /api/trigger/:triggerKey`

Example:

*Request:*
```

```

*Response:*
```

```

* `POST /api/trigger/:triggerKey/activate`

Example:

*Request:*
```

```

*Response:*
```

```

### Preference
* `POST /api/preference`

Example:

*Request:*
```

```

*Response:*
```

```

* `GET /api/preference/:userId`

Example:

*Request:*
```

```

*Response:*
```

```

* `PATCH /api/preference/:userId`

Example:

*Request:*
```

```

*Response:*
```

```

* `DELETE /api/preference/:userId`

Example:

*Request:*
```

```

*Response:*
```

```
