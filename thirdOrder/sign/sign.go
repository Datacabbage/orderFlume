package sign

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

func GetSign(signMsg map[string]string, key string) string {
	var lists []string
	var sign string
	for k, _ := range signMsg {
		lists = append(lists, k)
	}
	sort.Strings(lists)
	for i := 0; i < len(lists); i++ {
		key := lists[i]
		sign += key + "=" + signMsg[key] + "&"
	}
	sign += "key=" + key
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(sign))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetMeituanSign(signMsg map[string]string, key, url string) string {
	var lists []string
	var sign string
	for k, _ := range signMsg {
		lists = append(lists, k)
	}
	sort.Strings(lists)
	for i := 0; i < len(lists); i++ {
		key := lists[i]
		sign += key + "=" + signMsg[key] + "&"
	}
	sign = url + "?" + sign + key
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(sign))
	cipherStr := md5Ctx.Sum(nil)
	return string(cipherStr)
}
