amisgo is a Go+ [classfile module](https://github.com/goplus/gop/blob/main/doc/classfile.md), so
you can use it by following the classfile specification to reduce much codes.

1. Initialize your amisgo project

```sh
mkdir demo
cd demo
gop mod init demo
gop get github.com/zrcoder/amisgo@latest
```

2. create `main_amis.gox`, with content like:

```c
page().title("Example").body(
	page().title("Inner"),
)

listenAndServe! ":9090"
```

3. run:

```sh
gop mod tidy
gop run .
```
