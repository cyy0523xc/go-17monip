# 17MonIP Golang Lib

Base on https://github.com/axgle/ip

IP search based on 17monipdb, the IP database parser for china with golang.


## install

```sh
go get github.com/cyy0523xc/go-17monip
```

## example

```go
package main

import "github.com/cyy0523xc/go-17monip"

func main() {
    ip.Load("../17monipdb.dat")

    address := ip.Find("8.8.8.8") //output: GOOGLE\tGOOGLE\t
    println(address)

    address = ip.Find("202.106.46.151") // output: 中国\t北京\t
    println(address)

    address = ip.Find("202.115.128.64") // output: 中国\t四川\t成都理工大学
    println(address)
}
```

## 性能

在4核4G的个人电脑上测试，如下：

```
// go test -test.bench=".*"
BenchmarkFind-4      2000000          841 ns/op

// 2016-11-19
BenchmarkFind-4   	 2000000	       783 ns/op
BenchmarkFind2-4      100000         14106 ns/op

// 将原来的循环查找，改为二分法查找
BenchmarkFind-4    	10000000	       136 ns/op
BenchmarkFind2-4   	10000000	       133 ns/op
```

## License

BSD

## Version

## Thanks

* Paul Gao: for his 17monipdb.dat data.
