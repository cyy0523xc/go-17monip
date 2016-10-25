package ip

import "testing"

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
			t.Error(ip, found_adress)
		}
	}

	ip := "210.26.48.71"
	println(ip, Find(ip))

	ip = "14.28.107.174"
	println(ip, Find(ip))
}
