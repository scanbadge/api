# ScanBadge API
REST API for initiating, maintaining and authenticating [ScanBadge](https://scanbadge.xyz/discover), written in Go.

## Setup

1. `go get github.com/scanbadge/api`
- `go install github.com/scanbadge/api`
- Create a new key for creating/verifying JWT, e.g. `openssl rand -out $GOPATH/bin/scanbadge.key -base64 256`
- Add and edit the `config.json` to `$GOPATH/bin/config.json`.
- Use `cd $GOPATH/bin && ./api` to run ScanBadge API.
- Gorp will automatically create empty tables in the selected database.
- Add [the first API user](#how-do-i-create-the-first-api-user).

## Sample configuration

#### config.json
```
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

- `/auth`
- `/devices`
- `/logging`
- `/users`

See the [API documentation](https://scanbadge.xyz/documentation/api#endpoints) for more detailed information about our endpoints.

### How do I create the first API user?

Until we have implemented a proper way of configuring a first user, you must create the first user manually. This can be done using the following steps:

1. Generate a bcrypt hash of your desired password:
	```
	package main

	import (
		"fmt"
		"golang.org/x/crypto/bcrypt"
	)

	func main() {
		hash, err := bcrypt.GenerateFromPassword([]byte("yourpassword"), bcrypt.DefaultCost)

		if(err != nil) {
			panic(err)
		}
		fmt.Println(string(hash))
	}
	```
- Connect to your MySQL database.
- Insert a new user by using this query (replace appropriate values):

  ```
insert into users (`username`,`email`,`password`,`firstname`,`lastname`,`admin`)
values ('john','john.doe@example.org','the generated bcrypt hash, see step 1','John','Doe',1);
```

## Links
- [Project website](https://scanbadge.xyz/)
- [Full documentation](https://scanbadge.xyz/documentation)
- [Coding style rules](https://golang.org/doc/effective_go.html#formatting)
