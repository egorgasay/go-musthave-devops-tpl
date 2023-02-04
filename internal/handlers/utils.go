package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"devtool/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func NewCookie(key []byte, src []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(src)

	return fmt.Sprintf("%x", h.Sum(nil))
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
		src := []byte(fmt.Sprintf("%s:counter:%f", metric.ID, metric.Value))
		cookie = NewCookie(key, src)
	} else if metric.MType == "counter" {
		src := []byte(fmt.Sprintf("%s:gauge:%d", metric.ID, metric.Delta))
		cookie = NewCookie(key, src)
	}
	c.Header("hash", cookie)

	return cookie
}

func getSrc(mt storage.Metrics, key []byte) string {
	if mt.MType == "gauge" {
		src := []byte(fmt.Sprintf("%s:gauge:%f", mt.ID, mt.Value))
		nc := NewCookie(key, src)
		log.Printf("Generated!!!!:%s:%s:new %s\n", src, mt.Hash, nc)
		return nc
	} else if mt.MType == "counter" {
		src := []byte(fmt.Sprintf("%s:counter:%d", mt.ID, mt.Delta))
		return NewCookie(key, src)
	}

	return ""
}

func checkCookies(cookie string, metric storage.Metrics, key []byte) bool {
	cookie2 := getSrc(metric, key)
	return cookie2 == cookie
}
