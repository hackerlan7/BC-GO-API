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

func Transfer(c *gin.Context) {

	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	playerID := c.PostForm("playerID")
	uid := c.PostForm("uid")
	limit := c.PostForm("limit")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	singStr := fmt.Sprint(appSecret + endTime + limit + operatorID + playerID + requestTime + startTime)
	singCode := utils.GetSignature(singStr)

	form := url.Values{
		"appSecret":         {appSecret},
		"operatorID":        {operatorID},
		"startTime":         {startTime},
		"endTime":           {endTime},
		"playerID":          {playerID},
		"uid":               {uid},
		"limit":             {limit},
		"errorResponseCode": {errorResponseCode},
		"requestTime":       {requestTime},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/history/transfer", strings.NewReader(form.Encode()))
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

//定義一個當伺服器啟動接收到客戶端發出的post請求就會執行的函數
func HistoryBet(c *gin.Context) {
	//宣告要從表單獲取的資料
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	playerID := c.PostForm("playerID")
	betID := c.PostForm("betID")
	category := c.PostForm("category")
	betStatus := c.PostForm("betstatus")
	offset := c.PostForm("offset")
	limit := c.PostForm("limit")
	isSettlementTime := c.PostForm("isSettlementTime")
	errorResponseCode := c.PostForm("errorResponseCode")
	requestTime := utils.Time()

	singStr := fmt.Sprint(appSecret, betID, betStatus, category, endTime, isSettlementTime, limit, offset, operatorID, playerID, requestTime, startTime)
	singCode := utils.GetSignature(singStr)

	form := url.Values{
		// "appSecret":          {appSecret},
		"operatorID":         {operatorID},
		"startTime":          {startTime},
		"endTime":            {endTime},
		"playerID":           {playerID},
		"offset":             {offset},
		"limit":              {limit},
		"betID":              {betID},
		"category":           {category},
		"betstatus":          {betStatus},
		"isSettlementTime":   {isSettlementTime},
		"errorResponseCode ": {errorResponseCode},
		"requestTime":        {requestTime},
	}

	fmt.Println("appSecret:", appSecret)
	fmt.Println("startTime:", startTime)
	fmt.Println("endTime:", endTime)
	fmt.Println("playerID:", playerID)
	fmt.Println("betID", betID)
	fmt.Println("category", category)
	fmt.Println("betStatus", betStatus)
	fmt.Println("limit", limit)
	fmt.Println("offset", offset)
	fmt.Println("isSettlementTime", isSettlementTime)
	fmt.Println("requestTime", requestTime)
	fmt.Println("signature", singCode)

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/history/bet", strings.NewReader(form.Encode()))
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

func HistorySummary(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	isSettlementTime := c.PostForm("isSettlementTime")
	requestTime := utils.Time()
	errorResponseCode := c.PostForm("errorResponseCode")
	singStr := fmt.Sprint(appSecret + endTime + isSettlementTime + operatorID + requestTime + startTime)
	singCode := utils.GetSignature(singStr)

	form := url.Values{
		"appSecret":         {appSecret},
		"operatorID":        {operatorID},
		"startTime":         {startTime},
		"endTime":           {endTime},
		"isSettlementTime":  {isSettlementTime},
		"requestTime":       {requestTime},
		"errorResponseCode": {errorResponseCode},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/history/summary", strings.NewReader(form.Encode()))
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
func ReportBet(c *gin.Context) {
	//宣告要從表單獲取的資料
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	playerID := c.PostForm("playerID")
	betID := c.PostForm("betID")
	category := c.PostForm("category")
	betStatus := c.PostForm("betstatus")
	offset := c.PostForm("offset")
	limit := c.PostForm("limit")
	isSettlementTime := c.PostForm("isSettlementTime")
	requestTime := utils.Time()
	singStr := fmt.Sprint(appSecret + betID + betStatus + category + endTime + isSettlementTime + c.PostForm(limit) + c.PostForm(offset) + operatorID + playerID + requestTime + startTime)
	singCode := utils.GetSignature(singStr)

	form := url.Values{
		"appSecret":        {appSecret},
		"operatorID":       {operatorID},
		"startTime":        {startTime},
		"endTime":          {endTime},
		"playerID":         {playerID},
		"offset":           {offset},
		"limit":            {limit},
		"betID":            {betID},
		"category":         {category},
		"betstatus":        {betStatus},
		"isSettlementTime": {isSettlementTime},
		"requestTime":      {requestTime},
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/report/bet", strings.NewReader(form.Encode()))
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
func ReportSummary(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")

	requestTime := utils.Time()
	singStr := fmt.Sprint(appSecret + endTime + operatorID + requestTime + startTime)
	singCode := utils.GetSignature(singStr)

	form := url.Values{
		"appSecret":   {appSecret},
		"operatorID":  {operatorID},
		"startTime":   {startTime},
		"endTime":     {endTime},
		"requestTime": {requestTime},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/report/summary", strings.NewReader(form.Encode()))
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

func Betrecord(c *gin.Context) {
	operatorID := c.PostForm("operatorID")
	appSecret := c.PostForm("appSecret")
	betID := c.PostForm("betID")
	requestTime := utils.Time()
	singCode := utils.GetSignature(appSecret + betID + operatorID + requestTime)

	form := url.Values{
		"appSecret":   {appSecret},
		"operatorID":  {operatorID},
		"betID":       {betID},
		"requestTime": {requestTime},
	}
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "operatorID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	req, _ := http.NewRequest("POST", "https://uat-op-api.bpweg.com/report/betrecord", strings.NewReader(form.Encode()))
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
