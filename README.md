# Go Exception

## cannot find package

现象

```go
error.go:4:2: cannot find package "github.com/cockroachdb/errors" in any of:
        /usr/lib/go-1.13/src/github.com/cockroachdb/errors (from $GOROOT)
        /home/leison/go/src/github.com/cockroachdb/errors (from $GOPATH)
```

如何解决

通过go mod init 新建.mod文件，编辑.mod文件，添加require依赖

## proxy.golang.org connect: connection refused

现象

```go
go: github.com/cockroachdb/errors@v1.9.0: Get https://proxy.golang.org/github.com/cockroachdb/errors/@v/v1.9.0.mod: dial tcp 142.251.42.241:443: connect: connection refused
```

如何解决

```go
go env -w GOPROXY="https://mirrors.aliyun.com/goproxy,https://proxy.golang.org,direct"
```

# What is this?

Press the `.` key on any repository or pull request, or swap `.com` with `.dev` in the URL, to go directly to a VS Code environment in your browser.

![github dev](https://user-images.githubusercontent.com/856858/130119109-4769f2d7-9027-4bc4-a38c-10f297499e8f.gif)

# Why?

It’s a quick way to edit and navigate code. It's especially useful if you want to edit multiple files at a time or take advantage of all the powerful code editing features of Visual Studio Code when making a quick change. For more information, see our [documentation](https://github.co/codespaces-editor-help).
