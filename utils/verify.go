package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckSignature(signature, singSource string) error {
	has := md5.Sum([]byte(singSource))
	singMD5 := fmt.Sprintf("%x", has)
	fmt.Println("signature:", signature)
	fmt.Println("singMD5:", singMD5)
	if !strings.EqualFold(signature, singMD5) {
		return fmt.Errorf("\r\nrequest :%s\r\ngenerate:%s", signature, singMD5)
	}
	return nil
}

func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func ErrorResponse(c *gin.Context, code int, msg string, err error) {

	statusCode := code

	// check if dedicated status code
	errorResponseCode := c.PostForm("errorResponseCode")
	if errorResponseCode != "" && errorResponseCode == "200" {
		s, e := strconv.Atoi(errorResponseCode)
		if e == nil {
			statusCode = s
		}
	}

	errorMsg := msg
	if err != nil {
		errorMsg = fmt.Sprintf("%s: %v", msg, err)
	}
	c.Set("ErrorMsg", errorMsg)
	c.JSON(statusCode, gin.H{"error": msg, "code": code})
}
