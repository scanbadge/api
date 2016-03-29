# ScanBadge API
REST API for initiating, maintaining and authenticating [ScanBadge](https://scanbadge.xyz/discover), written in Go.

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
- `/devices`
- `/logging`
- `/users`

*<sup>1</sup> only `POST` requests supported*

See the [API documentation](https://scanbadge.xyz/documentation/api#endpoints) for more detailed information about our endpoints.

### How do I create the first API user?

Run API with flag `-add-user`, e.g. `$ ./api -add-user`. The following information is required:

1. Username
- Password
- Email
- First name
- Last name

If the user is successfully added, you can obtain an authentication token by sending a `POST` request to `/auth` with the following information: `username=yourusername&password=yourpassword`.

## Links
- [Project website](https://scanbadge.xyz/)
- [Full documentation](https://scanbadge.xyz/documentation)
- [Coding style rules](https://golang.org/doc/effective_go.html#formatting)
