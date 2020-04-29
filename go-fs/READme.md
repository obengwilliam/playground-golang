learnt about `go build`

Install executable in the bin directory of the gopath.
`go install`

Create an excutable in your current directory
```
go build -o some_name
```


Build for linux
```
GOOS=linux GOARCH=arm go build
```


Nb.
You can ask for help for any executable 
by providing the "-help" flag

like in our case

```
fs -help
```
