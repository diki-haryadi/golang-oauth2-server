package httputil

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIPAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")
	ipAddress = strings.Replace(ipAddress, " ", "", -1)

	if ipAddress == "" { //get from r.RemoteAddr
		ipFirst := strings.Replace(r.RemoteAddr, " ", "", -1)

		ipAddressArr := strings.Split(ipFirst, ",")
		if len(ipAddressArr) > 0 {
			ipAddress = ipAddressArr[0]
		}

		ipAddressArr = strings.Split(ipAddress, ":")
		if len(ipAddressArr) > 0 {
			ipAddress = ipAddressArr[0]
		}
	} else { //if no empty, looking for sign , and :

		ipAddressArr := strings.Split(ipAddress, ",")
		if len(ipAddressArr) > 0 {
			ipAddress = ipAddressArr[0]
		}

		ipAddressArr = strings.Split(ipAddress, ":")
		if len(ipAddressArr) > 0 {
			ipAddress = ipAddressArr[0]
		}
	}

	addr := net.ParseIP(ipAddress)
	if addr == nil {
		ipAddress = ""
	}

	return ipAddress
}

func GetUserAgent(r *http.Request) string {
	// device := GetOSName(r.FormValue("os_type"))
	// if device == utils.IosString || device == utils.AndroidString {
	// 	return APPSUserAgent
	// }
	return r.Header.Get("User-Agent")
}
