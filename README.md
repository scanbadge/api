# ScanBadge API
REST API for initiating, maintaining and authenticating [ScanBadge](https://scanbadge.xyz/discover), written in Go.

[![GoDoc](https://godoc.org/github.com/ScanBadge/api/endpoints?status.svg)](https://godoc.org/github.com/ScanBadge/api/endpoints)

## Setup

1. `$ go get -u github.com/scanbadge/api`
- `$ go install github.com/scanbadge/api`
- Create a new key for creating/verifying JWT, e.g. `$ openssl rand -out $GOPATH/bin/scanbadge.key -base64 256`
- Add and edit the `config.json` to `$GOPATH/bin/config.json`.
- Use `$ cd $GOPATH/bin && ./api` to run ScanBadge API.
- Gorp will automatically create empty tables in the selected database.
- Add [the first API user](#how-do-i-create-the-first-api-user).

## Sample configuration

#### config.json
```json
{
  "ServerHost": "localhost",
  "ServerPort": 8080,
  "Key": "scanbadge.key",
  "Database": {
    "Username": "username",
    "Password": "password",
    "DatabaseName": "scanbadge",
    "Protocol": "tcp",
    "Host": "localhost",
    "Port": "3306",
    "Charset": "utf8mb4,utf8",
    "Engine": "InnoDB",
    "Encoding": "UTF8"
  }
}
```

## FAQ

### Do you have a list of your API endpoints?

Yes, we have. Our API is RESTful, so endpoints support `GET`,`PUT`,`POST`,`DELETE` requests. The following endpoints are currently implemented:

- `/auth`<sup>1</sup>
- `/actions`
- `/conditions`
- `/devices`
- `/logs`
- `/users`

*<sup>1</sup> only used for [authentication](https://github.com/scanbadge/api#do-you-have-a-list-of-your-api-endpoints)*

See the [API documentation](https://scanbadge.xyz/documentation/api#endpoints) for more detailed information about our endpoints.

### How do I create the first API user?

Run API with flag `-add-user`, e.g. `$ ./api -add-user` and follow the on-screen instructions.

If the user is successfully added, you can obtain an authentication token by sending a `POST` request to `/auth` using `multipart/form-data`: `username=foo&password=bar` or use cURL:

    $ curl --form "username=foo" --form "password=bar" https://api.example.org/auth

If the authentication is successful, a JSON-encoded result with the authentication token will be returned, like so:

    {"token":"eyJhbGciOiJIUzI1NiIsImtpZCI6ImxvZ2luIiwidHlwIjoiSldUIn0.eyJleHAiOjQyOTQ5NjcyOTUsImlkIjoxLCJuYW1lIjoiRm9vIEJhciJ9.5wPuGctuwTb0EqD_ER1dGQQeK2RyIGq64w552_zW-sw"}

## Links
- [Project website](https://scanbadge.xyz/)
- [Full documentation](https://scanbadge.xyz/documentation)
- [Coding style rules](https://golang.org/doc/effective_go.html#formatting)
