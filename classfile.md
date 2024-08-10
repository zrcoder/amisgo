amisgo is a Go+ classfile module, so you can use it like classfile.

1. Initialize your amisgo project

```sh
mkdir demo
cd demo
gop mod init demo
gop get github.com/zrcoder/amisgo@latest
```

2. create `main_amis.gox`, with content like:

```c
index := page().title("Example").body(
	page().title("Inner"),
)

route "/", index
listenAndServe! ":9090"
```

3. run:

```sh
gop mod tidy
gop run .
```
