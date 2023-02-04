package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"devtool/internal/storage"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func newCookie(key []byte, src []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(src)

	return hex.EncodeToString(h.Sum(nil)) + "-" + hex.EncodeToString(src)
}

func getCookies(c *gin.Context) (cookie string, err error) {
	cookie = c.Request.Header.Get("hash")
	if cookie == "" {
		cookie, err = c.Cookie("token")
		if err != nil {
			return "", err
		}
	}

	return cookie, nil
}

func setCookies(c *gin.Context, key []byte, metric storage.Metrics) (cookie string) {
	if metric.MType == "gauge" {
		src := []byte(fmt.Sprintf("%s:counter:%d", metric.ID, metric.Delta))
		cookie = newCookie(key, src)
	} else if metric.MType == "counter" {
		src := []byte(fmt.Sprintf("%s:gauge:%f", metric.ID, *metric.Value))
		cookie = newCookie(key, src)
	}
	c.Header("hash", cookie)

	return cookie
}

func checkCookies(cookie string, key []byte) bool {
	arr := strings.Split(cookie, "-")
	k, v := arr[0], arr[1]

	sign, err := hex.DecodeString(k)
	if err != nil {
		return false
	}

	data, err := hex.DecodeString(v)
	if err != nil {
		return false
	}

	h := hmac.New(sha256.New, key)
	h.Write(data)

	return hmac.Equal(sign, h.Sum(nil))
}
