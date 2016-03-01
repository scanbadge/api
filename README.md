# Introduction
API for initiating, maintaining and authenticating [ScanBadge](https://scanbadge.xyz/discover).

This API is written in Go.

## Setup

1. `go get github.com/scanbadge/api`
- `go install github.com/scanbadge/api`
- Create a new key for creating/verifying JWT, e.g. `openssl rand -out $GOPATH/bin/scanbadge.key -base64 256`
- Edit the `config.example.json` file as following:
  - Add the missing MySQL details (`Username`,`Password`,`DatabaseName`)
  - Edit applicable values to match your current MySQL configuration
  - Add the **relative path** to the key file, e.g. if the key is stored in `/etc/ssl/private/scanbadge.key` and your application runs from `/home/user/go/bin/`, use `../../../etc/ssl/private/scanbadge.key`.
  - Rename `config.example.json` to `config.json`.
  - Move `config.json` to the directory of your current ScanBadge application (e.g. `$GOPATH/bin/`).
- Use `cd $GOPATH/bin && ./api` to run ScanBadge API.
- Gorp will automatically create empty tables in the selected database.

## FAQ

### Do you have a list of your API endpoints?

Yes, we have. Our API is RESTful, so endpoints support `GET`,`PUT`,`POST`,`DELETE` requests. The following endpoints are currently implemented:

- `/users`
- `/devices`

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
