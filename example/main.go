package main

import "github.com/cyy0523xc/go-17monip"

func main() {
	ip.Load("../17monipdb.dat")

	address := ip.Find("8.8.8.8") // output: GOOGLE\tGOOGLE\t
	println(address)

	address = ip.Find("202.106.46.151") // output: 中国\t北京\t北京
	println(address)

	address = ip.Find("202.115.128.64") // output: 中国\t四川\t成都
	println(address)

	address = ip.Find("116.22.198.67") // output: 中国\t广东\t广州
	println(address)
}
