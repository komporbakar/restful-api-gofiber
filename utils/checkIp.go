package utils

var whitelistIp = []string{"127.0.0.1", "127.0.0.1:8080"}

func CheckIp(ip string) bool {

	for _, v := range whitelistIp {
		if v == ip {
			return true
		}
	}
	return false
}
