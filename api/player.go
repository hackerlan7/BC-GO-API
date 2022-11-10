package api

import (
	"encoding/json"
	"io/ioutil"
	"myproject/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var link string = "https://uat-op-api.bpweg.com/player"

func CreatePlayer(c *gin.Context) {

	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	nickname := c.PostForm("nickname")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	//驗簽的順序按照英文字母順序ABCD~Z

	singCode := utils.GetSignature(appSecret + nickname + operatorID + playerID + requestTime)

	form := url.Values{

		"operatorID":        {operatorID},
		"playerID":          {playerID},
		"nickname":          {nickname},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID", "playerID", "nickname"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", link+"/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)

	body, _ := ioutil.ReadAll(r.Body)
	if r != nil {

		defer r.Body.Close()
	}

	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, (data))

}
func Login(c *gin.Context) {
	//

	appSecret := c.PostForm("appSecret")
	operatorID := c.PostForm("operatorID")
	playerID := c.PostForm("playerID")
	//自動產生時間
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")

	//待加密字串自動生成簽名,注意需轉換格式
	singCode := utils.GetSignature(appSecret + operatorID + playerID + requestTime)

	form := url.Values{

		"operatorID":        {operatorID},
		"playerID":          {playerID},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)
	//客戶端完成之後要關閉請求

	//讀取請求的資料
	body, _ := ioutil.ReadAll(r.Body)
	if r != nil {

		defer r.Body.Close()
	}
	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, (data))

}

func Logout(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	singCode := utils.GetSignature(appSecret + operatorID + playerID + requestTime)

	form := url.Values{

		"operatorID":        {operatorID},
		"playerID":          {playerID},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", link+"/logout", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)
	if r != nil {

		defer r.Body.Close()
	}
	body, _ := ioutil.ReadAll(r.Body)

	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, (data))
}

func Balance(c *gin.Context) {

	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	singCode := utils.GetSignature(appSecret + operatorID + playerID + requestTime)
	form := url.Values{

		"operatorID":        {operatorID},
		"playerID":          {playerID},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}

	req, _ := http.NewRequest("POST", link+"/balance", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)
	if r != nil {

		defer r.Body.Close()
	}

	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)

	c.JSON(200, data)
}

func Deposit(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	uid := c.PostForm("uid")
	amount := c.PostForm("amount")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	singCode := utils.GetSignature(amount + appSecret + operatorID + playerID + requestTime + uid)
	form := url.Values{
		"appSecret":         {appSecret},
		"operatorID":        {operatorID},
		"playerID":          {playerID},
		"uid":               {uid},
		"amount":            {amount},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", link+"/deposit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)
	if r != nil {

		defer r.Body.Close()
	}
	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, (data))
}

func Withdraw(c *gin.Context) {

	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	uid := c.PostForm("uid")
	amount := c.PostForm("amount")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	singCode := utils.GetSignature(amount + appSecret + operatorID + playerID + requestTime + uid)
	form := url.Values{
		"appSecret":         {appSecret},
		"operatorID":        {operatorID},
		"playerID":          {playerID},
		"uid":               {uid},
		"amount":            {amount},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", link+"/withdraw", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)
	if r != nil {

		defer r.Body.Close()
	}
	body, _ := ioutil.ReadAll(r.Body)

	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, (data))
}
