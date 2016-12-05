# bloodlines
![Build Status](https://travis-ci.org/ghmeier/bloodlines.svg?branch=master)

A go service for sending Expresso emails

## Quick Start
```bash
$ docker run ghmeier/bloodlines:latest
```

## Setup
```bash
$ go get github.com/ghmeier/bloodlines
$ cd $GOPATH/src/github.com/ghmeier/bloodlines
$ make deps
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
