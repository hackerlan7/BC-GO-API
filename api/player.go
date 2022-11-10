package api

import (
	"encoding/json"
	"fmt"
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

	singStr := fmt.Sprint(appSecret + nickname + operatorID + playerID + requestTime)
	singCode := utils.GetSignature(singStr)

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

	// Step 2: Check length
	if len(operatorID) > 20 || len(playerID) > 50 || len(playerID) > 200 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
		return
	}
	req, _ := http.NewRequest("POST", link+"/create", strings.NewReader(form.Encode()))
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

}
func Login(c *gin.Context) {
	//

	appSecret := c.PostForm("appSecret")
	operatorID := c.PostForm("operatorID")
	playerID := c.PostForm("playerID")
	//自動產生時間
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	if errorResponseCode == "200" {
		errorResponseCode = "200"
	}

	//待加密字串自動生成簽名,注意需轉換格式
	singStr := fmt.Sprint(appSecret + operatorID + playerID + requestTime)
	//將字串利用參數的方式帶入加密的函數
	singCode := utils.GetSignature(singStr)

	form := url.Values{

		"operatorID":  {operatorID},
		"playerID":    {playerID},
		"requestTime": {requestTime},
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/player/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("signature", singCode)
	clt := http.Client{}
	r, _ := clt.Do(req)
	//客戶端完成之後要關閉請求
	if r != nil {

		defer r.Body.Close()
	}
	//讀取請求的資料
	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	c.JSON(200, data)

}

func Logout(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	requestTime := utils.Time()

	//待加密字串自動生成簽名,注意需轉換格式
	singStr := fmt.Sprint(appSecret + operatorID + playerID + requestTime)
	//將字串利用參數的方式帶入加密的函數
	singCode := utils.GetSignature(singStr)

	form := url.Values{

		"operatorID":  {operatorID},
		"playerID":    {playerID},
		"requestTime": {requestTime},
	}

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length
	if len(operatorID) > 20 || len(playerID) > 50 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
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

	c.JSON(200, (body))
}

func Balance(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	requestTime := utils.Time()
	singStr := fmt.Sprint(appSecret + operatorID + playerID + requestTime)
	singCode := utils.GetSignature(singStr)
	form := url.Values{

		"operatorID":  {operatorID},
		"playerID":    {playerID},
		"requestTime": {requestTime},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length
	if len(operatorID) > 20 || len(playerID) > 50 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
		return
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

	c.JSON(200, string(body))
}

func Deposit(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	uid := c.PostForm("uid")
	amount := c.PostForm("amount")
	requestTime := utils.Time()
	singStr := fmt.Sprint(amount + appSecret + operatorID + playerID + requestTime + uid)
	singCode := utils.GetSignature(singStr)
	form := url.Values{
		"appSecret":   {appSecret},
		"operatorID":  {operatorID},
		"playerID":    {playerID},
		"uid":         {uid},
		"amount":      {amount},
		"requestTime": {requestTime},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length
	if len(operatorID) > 20 || len(playerID) > 50 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
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

	c.JSON(200, string(body))
}

func Withdraw(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	playerID := c.PostForm("playerID")
	uid := c.PostForm("uid")
	amount := c.PostForm("amount")
	requestTime := utils.Time()
	singStr := fmt.Sprint(amount + appSecret + operatorID + playerID + requestTime + uid)
	singCode := utils.GetSignature(singStr)
	form := url.Values{
		"appSecret":   {appSecret},
		"operatorID":  {operatorID},
		"playerID":    {playerID},
		"uid":         {uid},
		"amount":      {amount},
		"requestTime": {requestTime},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	// Step 2: Check length
	if len(operatorID) > 20 || len(playerID) > 50 {
		utils.ErrorResponse(c, 400, "Maximum request length exceeded", nil)
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

	c.JSON(200, string(body))
}
