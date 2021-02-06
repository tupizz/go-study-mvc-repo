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

