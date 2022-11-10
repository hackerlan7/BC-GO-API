package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAmount(amount string) (int64, error) {
	gainBal, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0, err
	} else if gainBal < 1 {
		return 0, fmt.Errorf("amount < 1: %v", gainBal)
	}
	return gainBal, nil
}

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

// func CheckAppSecret(operatorID, appSecret string) (*opf.OperatorProfile, error) {
// 	dataOp, err := srvclient.OperatorClient.GetOperatorProfile(context.TODO(), &opf.GetOperatorProfileRequest{
// 		OperatorID: operatorID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	} else if dataOp.OperatorProfile == nil {
// 		return nil, fmt.Errorf("not found:\r\noperatorID:%s\r\nappSecret :%s", operatorID, appSecret)
// 	} else if dataOp.OperatorProfile.AppSecret != appSecret {
// 		return nil, fmt.Errorf("\r\nrequest:%s\r\nserver :%s", appSecret, dataOp.OperatorProfile.AppSecret)
// 	} else if dataOp.OperatorProfile.IsSingleWallet {
// 		return nil, fmt.Errorf("%s IsSingleWallet", operatorID)
// 	}
// 	return dataOp.OperatorProfile, nil
// }

func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func CheckRequestTime(requestTime string) error {
	rtInt, rtErr := strconv.ParseInt(requestTime, 10, 64)
	if rtErr != nil {
		return fmt.Errorf("incorrect format:\r\nrequestTime :%s", requestTime)
	}

	if rtInt-time.Now().Unix() > 120 || time.Now().Unix()-rtInt > 120 {
		return fmt.Errorf("expired:\r\nrequestTime :%s", requestTime)
	}
	return nil
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
