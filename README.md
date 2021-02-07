# Learning about go

Building application gallery with go

```bash
go run ./main.go
```

# Add hot reload in our application

Firstly you must install to your packages

```bash
go get github.com/pilu/fresh
```

After installing you'll be able to type the following command on your terminal and kinda watch for the changes

```bash
fresh
```

Also, is important to create a config file to tell fresh what kind of files he needs to pay attention to.


New file: `runner.conf`

```
root:              .
tmp_path:          ./tmp
build_name:        runner-build
build_log:         runner-build-errors.log
valid_ext:         .go, .tpl, .tmpl, .html, .gohtml
ignored:           assets, tmp
build_delay:       300
colors:            1
log_color_main:    cyan
log_color_build:   yellow
log_color_runner:  green
log_color_watcher: magenta
log_color_app:     bold_white
```

## Add hot reload watching your changes with Air

- https://github.com/cosmtrek/air
- https://techinscribed.com/5-ways-to-live-reloading-go-applications/
- Create the `.air.toml`
- Run `air` on the command line inside the folder

## Installing a router for our api/application

Here with the `-u` flag, we're telling for gomod to get the last updated version of the library

```bash
go get -u github.com/gorilla/mux
```

After this, we must import it to our `main.go`

```go
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // here 😍
)
```
## How to run our playground files? easy peace

```bash
go run playground/main.go
```

## Working with Gorilla Schema

```bash
go get -u github.com/gorilla/schema
```

```go
values := map[string]string{
	"name": {"John"}, // a string to a slice
	"Phone": {"999-999-999"} // a string to a slice
}
```

It will be mapped for a structure like:

```go
type Person struct {
	Name 		string
	Phone 	string
}
```

Also we'll use struct tags

```
Struct tags are a form of metadata that can be added to to the fields of any struct that you create. When coding them in they will be in the format ‘key:"value"‘, and are found directly after the field’s type (see Listing 7.23 for an example).

Later, when your program is running, other packages can use the reflect package to look up each struct tag and then use that data to determine how it should proceed. For example, the encoding/json package uses this data to determine what key you want used for each field when converting a struct into JSON.

While it isn’t a requirement, most packages will use their package name as the key for all of their struct tags. This makes it easier to determine what package each struct tag is used for when reading your source code.

In Listing 7.23 each struct tag starts with the key schema which denotes that it is used by a package named schema. Similarly, all struct tags for the encoding/json package use the key json.
```

```go
type SignupForm struct {
  Email    string `schema:"email"`
  Password string `schema:"password"`
}
```