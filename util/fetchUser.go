package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	ACCESS_FROM_USER = 0
	COLON            = ":"
)

func GetUserName(ak string, resourceOwner uint64) string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(ACCESS_FROM_USER))
	buffer.WriteString(COLON)
	buffer.WriteString(strconv.FormatUint(resourceOwner, 10))
	buffer.WriteString(COLON)
	buffer.WriteString(ak)
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func GetUserNameBySTSToken(ak string, resourceOwnerId uint64, stsToken string) string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(ACCESS_FROM_USER))
	buffer.WriteString(COLON)
	buffer.WriteString(strconv.FormatUint(resourceOwnerId, 10))
	buffer.WriteString(COLON)
	buffer.WriteString(ak)
	buffer.WriteString(COLON)
	buffer.WriteString(stsToken)
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func GetPassword(sk string) string {
	now := time.Now()
	currentMillis := strconv.FormatInt(now.UnixNano()/1000000, 10)
	var buffer bytes.Buffer
	buffer.WriteString(strings.ToUpper(HmacSha1(currentMillis, sk)))
	buffer.WriteString(COLON)
	buffer.WriteString(currentMillis)
	fmt.Println(currentMillis)
	fmt.Println(HmacSha1(sk, currentMillis))
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func HmacSha1(keyStr string, message string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
