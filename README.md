# bloodlines
A go service for sending Expresso emails

## Setup
```go
make deps
make run
```

## API
__All api methods should return `{"success":true}`__

* `POST /api/content`
* `GET /api/content`
* `GET /api/content/:contentId`
* `PUT /api/content/:contentId`
* `DELETE /api/content/:contentId`
