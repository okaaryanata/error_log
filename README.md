## error_log

#### Installation

```
$ go get -u github.com/okaaryanata/error_log
```

#### Requirement

- **mysql** database is installed
- there is a table named **error_logs** in your database
- **error_logs** table structure:
  <img src="images/table structure.png" height="125">

#### How to use

```go
// Example you want to user the error_log on you main package
package main

import (
	"github.com/okaaryanata/error_log/errorlog"
)

func main() {
  // Function ConnectLog take 3 parameters (database url, environment, repo-name)
  errlog := errorlog.ConnectLog(os.Getenv("LOG_SOURCE"), os.Getenv("GO_ENV"), "master-data-store")

  defer errlog.Close()

  // save error message to your DB
  // Func take 2 parameter (error and payload that triggered error)
  // error = error
  // payload = string
  errlog.Error(error, payload)
}
```
