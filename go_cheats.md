# Things to do with Go

1.Using modules

- <https://www.digitalocean.com/community/tutorials/how-to-use-go-modules>

```
go mod init chapter-1
go: creating new go.mod: module chapter-1
go: to add module requirements and sums:
go mod tidy
```

2. Compile programas
- To compile for windows OS
```
GOOS=windows GOARCH=amd64 go build -o Chapter_1-windows.exe
```
- To compile for Linux
```
$ GOOS=linux GOARCH=amd64 go build -o Chapter_1-linux
```
