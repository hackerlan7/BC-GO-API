package api

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Param struct {
	OperatorID        string `json:"operatorID"`
	AppSecret         string `json:"appSecret"`
	PlayerID          string `json:"playerID"`
	Nickname          string `json:"nickname"`
	Uid               string `json:"uid"`
	Amount            int    `json:"amount"`
	Limit             int    `json:"limit"`
	Offset            int    `json:"offset"`
	BetID             string `json:"betID"`
	Category          string `json:"category"`
	BetStatus         string `json:"betStatus"`
	IsSettlementTime  int    `json:"isSettlementTime"`
	StartTime         int    `json:"startTime"`
	EndTime           int    `json:"endTime"`
	RequestTime       int    `json:"requestTime"`
	ErrorResponseCode string `json:"errorResponseCode"`
	Signature         string `json:"signature"`
}

func Check(c *gin.Context) {
	urlstr := "https://uat-op-api.bpweg.com/check"
	//多用於查詢參數和表單的值
	appSecret := c.PostForm("appSecret")
	operatorID := c.PostForm("operatorID")

	form := url.Values{}
	form.Add("operatorID", operatorID)
	form.Add("appSecret", appSecret)
	r, _ := http.PostForm(urlstr, form)
	//check不需驗簽不用設請求頭

	if r != nil {

		defer r.Body.Close()
	}
	body, _ := ioutil.ReadAll(r.Body)
	//http.statusok=200
	c.JSON(200, string(body))
}

//  req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/deposit", strings.NewReader(form.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("signature", signature)
// 	clt := http.Client{}
// 	r, _ := clt.Do(req)
