# godb-go-driver
A [GoDB](https://github.com/MattLaidlaw/godb) client for the Go language. The *godb-go-driver* module supplies the *driver* package which defines a GoDB client.\
This client can be used in Go programs to interact with a running GoDB server.

# Requirements
* Go 1.17

# Install
```go get github.com/MattLaidlaw/GoDB-Go-Driver@0cc1e1b```

# Usage
## Creating a GoDB client
```
import "github.com/MattLaidlaw/GoDB-Go-Driver/pkg/driver"

...

client, err := driver.NewClient("localhost:6342")
```

## Supported methods

### Client.Set
Inserts a new key-value pair into the database. Returns the count of inserted/replaced items.\
```insertedCount, err := client.Set("key", "val")```\

### Client.Get
Returns the value that matches "key" in the database. This may be an empty string if the key does not exist.\
```value, err := client.Get("key")```

### Client.Del
Deletes the item that matches "key" in the database and returns the count of deleted items. If no such element exists, nothing is deleted, and a count of zero is returned.\
```deletedCount, err := client.Del("key")```
