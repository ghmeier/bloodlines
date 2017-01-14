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

#### `POST /api/content` creates and adds a new content record to the database.

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

#### `GET /api/content?offset=0&limit=20` returns up to `limit` content records starting from `offset` when ordered by contentId

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

#### `GET /api/content/:contentId` returns the content record with the given contentID

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

#### `PUT /api/content/:contentId` updates the content record with the given contentID to match the provided data. This just overrides values, so anything not present in the request will be set to NULL

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

#### `DELETE /api/content/:contentId` sets the contentStatus to INACTIVE so it won't be shown when getting all content. The data can still be accessed by getting the content object by contentId

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
#### `POST /api/receipt/send` creates queues a message based on the given receipt data.

Example:

*Request:*
```
POST localhost:8080/api/receipt/send
{
	"values": {
		"greeting": "Whaddup",
		"first_name": "garret"
	},
	"contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
	"userId": "afee47b7-4eff-4826-a12c-86affd91a2d9"
}
```

*Response:*
```
{
  "data": {
    "id": "78b0444c-da8c-11e6-bc32-0021ccdc3511",
    "ts": "2017-01-14T13:05:53.647317899-06:00",
    "values": {
      "first_name": "garret",
      "greeting": "Whaddup"
    },
    "sendState": "READY",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "userId": "afee47b7-4eff-4826-a12c-86affd91a2d9"
  }
}
```

#### `GET /api/receipt` returns an array of up to `limit` receipt records starting at the `offset` record when ordered by receiptID

Example:

*Request:*
```
GET localhost:8008/api/receipt?offset=0&limit=1
```

*Response:*
```
{
  "data": [
    {
      "id": "7890031d-da8b-11e6-86c2-0021ccdc3511",
      "ts": "2017-01-14T18:58:43Z",
      "values": {
        "first_name": "garret",
        "greeting": "Whaddup"
      },
      "sendState": "FAILURE",
      "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
      "userId": "afee47b7-4eff-4826-a12c-86affd91a2d9"
    }
  ]
}
```

#### `GET /api/receipt/:receiptId` returns the receipt record with the given receiptID

Example:

*Request:*
```
GET localhost:8080/api/receipt/
```

*Response:*
```
{
  "data": {
    "id": "7890031d-da8b-11e6-86c2-0021ccdc3511",
    "ts": "2017-01-14T18:58:43Z",
    "values": {
      "first_name": "garret",
      "greeting": "Whaddup"
    },
    "sendState": "FAILURE",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "userId": "afee47b7-4eff-4826-a12c-86affd91a2d9"
  }
}
```

### Job
*NOT IMPLEMENTED*

### Trigger
#### `POST /api/trigger` creates a new trigger with the given data.

Example:

*Request:*
```
POST localhost:8080/api/trigger
{
  "values": {
    "greeting": "Whaddup"
  },
  "tkey": "welcome",
  "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004"
}
```

*Response:*
```
{
  "data": {{
  "data": {
    "id": "ca6044a6-da8d-11e6-bc32-0021ccdc3511",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "tkey": "welcome",
    "values": {
      "greeting": "Whaddup"
    }
  }
}
    "id": "ca6044a6-da8d-11e6-bc32-0021ccdc3511",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "tkey": "welcome",
    "values": {
      "greeting": "Whaddup"
    }
  }
```

#### `GET /api/trigger` returns a list of `limit` trigger records starting at the `offset` record when they're ordered by triggerID

Example:

*Request:*
```
GET localhost:8080/api/trigger?limit=1&offset=0
```

*Response:*
```
{
  "data": [
    {
      "id": "ca6044a6-da8d-11e6-bc32-0021ccdc3511",
      "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
      "tkey": "welcome",
      "values": {
        "greeting": "Whaddup"
      }
    }
  ]
}
```

#### `GET /api/trigger/:triggerKey` returns the record for the trigger with the given triggerKey

Example:

*Request:*
```
GET localhost:8080/api/trigger/welcome
```

*Response:*
```
{
  "data": {
    "id": "ca6044a6-da8d-11e6-bc32-0021ccdc3511",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "tkey": "welcome",
    "values": {
      "greeting": "Whaddup"
    }
  }
}
```

#### `PUT /api/trigger/:triggerKey` updates the trigger record with new data. All properties will be overwritten.

Example:

*Request:*
```
PUT localhost:8080/api/trigger/welcome
{
  "values": {
    "greeting": "Hello there"
  },
  "tkey": "welcome",
  "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004"
}
```

*Response:*
```
{
  "data": {
    "id": "",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "tkey": "welcome",
    "values": {
      "greeting": "Hello there"
    }
  }
}
```

#### `DELETE /api/trigger/:triggerKey` removes the trigger record

Example:

*Request:*
```
DELETE localhost:8080/api/trigger/welcome
```

*Response:*
```
{
  "success": true
}
```

#### `POST /api/trigger/:triggerKey/activate` activates a trigger, creating a new receipt, and sending content to the user provided. Items in the 'values' property will take precedence over the defaults stored in the trigger.

Example:

*Request:*
```
POST localhost:8080/api/trigger/welcome/activate
{
  "values": {
    "first_name": "garret"
  },
  "userId": "afee47b7-4eff-4826-a12c-86affd91a2d9"
}
```

*Response:*
```
{
  "data": {
    "receiptId": "ce678bf3-da8e-11e6-bc32-0021ccdc3511",
    "contentId": "dd82cc65-d79d-11e6-9d4c-0242ac120004"
  }
}
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
