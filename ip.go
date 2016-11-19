package ip

import (
	//"fmt"
	"io/ioutil"
	"net"
)

var (
	buffer   []byte
	ipIndex  []byte
	ipOffset uint32
)

func Load(ipBinaryFilePath string) {
	// 避免多次初始化
	if buffer == nil {
		var err error
		if buffer, err = ioutil.ReadFile(ipBinaryFilePath); err != nil {
			panic(err)
		}

		ipOffset = bytesBigEndianToUint32(buffer[:4])
		ipIndex = buffer[4:ipOffset]
	}
}

func Find(ipAddress string) string {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return ""
	}
	ip = ip.To4()

	var (
		tmp         = make([]byte, 4)
		indexOffset uint32
		indexLength uint32
	)
	nip := bytesBigEndianToUint32(ip)
	tmpOffset := int(ip[0]) << 2
	startLen := bytesLittleEndianToUint32(ipIndex[tmpOffset : tmpOffset+4])
	for start := startLen*8 + 1024; start < ipOffset-1028; start += 8 {
		if bytesBigEndianToUint32(ipIndex[start:start+4]) >= nip {
			indexLength = 0xFF & uint32(ipIndex[start+7])
			tmp[0] = ipIndex[start+4]
			tmp[1] = ipIndex[start+5]
			tmp[2] = ipIndex[start+6]
			indexOffset = bytesLittleEndianToUint32(tmp)
			break
		}
	}
	if indexOffset == 0 {
		return ""
	}

	pos := indexOffset + ipOffset - 1024
	return string(ipIndex[pos-4 : pos+indexLength-4])
}

//binary.BigEndian.Uint32(b)
func bytesBigEndianToUint32(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func bytesLittleEndianToUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
