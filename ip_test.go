package ip

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	Load("17monipdb.dat")
}

func TestFind(t *testing.T) {
	ip_addresses := map[string]string{
		"8.8.8.8":        "GOOGLE\tGOOGLE\t",
		"8.8.4.4":        "GOOGLE\tGOOGLE\t",
		"202.106.195.68": "中国\t北京\t",
		"123.118.91.246": "中国\t北京\t",
		"202.115.128.64": "中国\t四川\t成都理工大学",
		"110.106.46.151": "中国\t辽宁\t",
		"11.106.46.151":  "美国\t美国\t",
		"20.106.46.151":  "美国\t美国\t",
		"1.106.46.151":   "韩国\t韩国\t",
		"99.106.46.151":  "加拿大\t加拿大\t",
		"10.106.46.151":  "局域网\t局域网\t",
		"116.22.198.67":  "",
	}

	for ip, address := range ip_addresses {
		found_adress := Find(ip)
		if found_adress != address {
			fmt.Println(ip, found_adress)
		}
	}

	ip := "210.26.48.71"
	println(ip, Find(ip))

	ip = "210.26.48.71"
	println(ip, Find(ip))

	ip = os.Getenv("TEST_IP")
	if len(ip) > 1 {
		println(ip, Find(ip))
	}
}

func BenchmarkFind(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ip := "121.229.193.12"
			address := Find(ip)
			if len(address) < 6 {
				b.Fatalf("ERROR %s", address)
			}
		}
	})
}

func BenchmarkFind2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ip := "202.115.128.64"
			address := Find(ip)
			if len(address) < 6 {
				b.Fatalf("ERROR %s", address)
			}
		}
	})
}
