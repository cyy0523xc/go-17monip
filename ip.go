package ip

import (
	//"fmt"
	"io/ioutil"
	"net"
	"strings"
)

var (
	ipIndex  []byte
	ipOffset uint32
	ipLen    int
)

func Load(ipBinaryFilePath string) {
	// 避免多次初始化
	if ipIndex == nil {
		var err error
		var buffer []byte
		if buffer, err = ioutil.ReadFile(ipBinaryFilePath); err != nil {
			panic(err)
		}

		ipOffset = bytesBigEndianToUint32(buffer[:4])
		ipIndex = buffer[4:ipOffset]
		ipLen = len(ipIndex)
	}
}

// Find TODO IPv6的地址暂时不做处理
func Find(ipAddress string) string {
	if tmp := strings.Split(ipAddress, ":"); len(tmp) > 1 {
		return ""
	}

	ip := net.ParseIP(ipAddress)
	if ip.IsMulticast() || ip.IsLoopback() {
		return ""
	}

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

	start := startLen*8 + 1024
	end := ipOffset - 1028
	if start < end {
		if index := fastFind(start, end, nip); index > 0 {
			//println(index)
			indexLength = 0xFF & uint32(ipIndex[index+7])
			tmp[0] = ipIndex[index+4]
			tmp[1] = ipIndex[index+5]
			tmp[2] = ipIndex[index+6]
			indexOffset = bytesLittleEndianToUint32(tmp)
		}
	}
	/*for index := startLen*8 + 1024; index < ipOffset-1028; index += 8 {
		if bytesBigEndianToUint32(ipIndex[index:index+4]) >= nip {
			indexLength = 0xFF & uint32(ipIndex[index+7])
			tmp[0] = ipIndex[index+4]
			tmp[1] = ipIndex[index+5]
			tmp[2] = ipIndex[index+6]
			indexOffset = bytesLittleEndianToUint32(tmp)
			break
		}
	}*/
	if indexOffset == 0 {
		return ""
	}

	pos := indexOffset + ipOffset - 1024
	return string(ipIndex[pos-4 : pos+indexLength-4])
}

// fastFind 使用二分法查找
func fastFind(start, end, nip uint32) uint32 {
	//println(start, end)
	if start+8 >= end {
		if bytesBigEndianToUint32(ipIndex[end:end+4]) >= nip {
			return end
		}
		return 0
	}

	// 计算新的开始值
	tmp := (end - start) >> 1
	newStart := start + tmp - tmp&7

	// 二分
	tmpIp := bytesBigEndianToUint32(ipIndex[newStart : newStart+4])
	if tmpIp > nip {
		return fastFind(start, newStart, nip)
	} else if tmpIp < nip {
		return fastFind(newStart, end, nip)
	}

	return newStart
}

//binary.BigEndian.Uint32(b)
func bytesBigEndianToUint32(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func bytesLittleEndianToUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
