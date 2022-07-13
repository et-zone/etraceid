# etraceid

### install
```go get github.com/et-zone/etraceid```

### example
```	
trace:=etraceid.MakeTraceID()
fmt.Println(trace.GetIP())
ctx:=etraceid.ToContext(context.TODO(),trace)
s:=etraceid.ContexFormat(ctx)
fmt.Println(s)
```