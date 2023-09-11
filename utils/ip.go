package utils

import (
	"strconv"
	"strings"
)

func IsLocalIp(ip string) bool {
	/*
		局域网（intranet）的IP地址范围包括：

		10．0．0．0／8－－10．0．0．0～10．255．255．255（A类）

		172．16．0．0／12－172．16．0．0－172．31．255．255（B类）

		192．168．0．0／16－－192．168．0．0～192．168．255．255（C类）
	*/
	ipAddr := strings.Split(ip, ".")

	if strings.EqualFold(ipAddr[0], "10") {
		return true
	} else if strings.EqualFold(ipAddr[0], "172") {
		addr, _ := strconv.Atoi(ipAddr[1])
		if addr >= 16 && addr < 31 {
			return true
		}
	} else if strings.EqualFold(ipAddr[0], "192") && strings.EqualFold(ipAddr[1], "168") {
		return true
	}
	return false
}
