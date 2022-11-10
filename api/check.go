package api

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// type Param struct {
// 	OperatorID        string `json:"operatorID"`
// 	AppSecret         string `json:"appSecret"`
// 	PlayerID          string `json:"playerID"`
// 	Nickname          string `json:"nickname"`
// 	Uid               string `json:"uid"`
// 	Amount            int    `json:"amount"`
// 	Limit             int    `json:"limit"`
// 	Offset            int    `json:"offset"`
// 	BetID             string `json:"betID"`
// 	Category          string `json:"category"`
// 	BetStatus         string `json:"betStatus"`
// 	IsSettlementTime  int    `json:"isSettlementTime"`
// 	StartTime         int    `json:"startTime"`
// 	EndTime           int    `json:"endTime"`
// 	RequestTime       int    `json:"requestTime"`
// 	ErrorResponseCode string `json:"errorResponseCode"`
// 	Signature         string `json:"signature"`
// }

func Check(c *gin.Context) {
	urlstr := "https://uat-op-api.bpweg.com/check"
	//多用於查詢參數和表單的值
	appSecret := c.PostForm("appSecret")
	operatorID := c.PostForm("operatorID")
	errorResponseCode := c.PostForm("errorResponseCode")
	form := url.Values{

		"operatorID":        {operatorID},
		"appSecret":         {appSecret},
		"errorResponseCode": {errorResponseCode},
	}

	r, _ := http.PostForm(urlstr, form)

	if r != nil {

		defer r.Body.Close()
	}

	//讀取得到的請求體訊息是byte
	body, _ := ioutil.ReadAll(r.Body)

	//收到的byte轉成字串用JSON格式呈現

	c.String(200, string(body))
	//確認是否有必填的參數

}
