package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

//產生requestTime 並返回字串型態的資料
func Time() string {

	return strconv.FormatInt(time.Now().Unix(), 10)

}

//加密生成signature 並返回字串型態的資料

func GetSignature(singSource string) string {

	// singSource為需要加密的字串當成參數送進來
	singCode := fmt.Sprintf("%x", md5.Sum([]byte(singSource)))
	//return singcode 是我需要的驗簽
	return singCode

}
