package app

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

//InitLog 初始化日誌
func InitLog() {
	//設定日誌格式
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //時間格式
	})

	log.SetLevel(log.DebugLevel)
	//宣告日誌切割
	logger := &lumberjack.Logger{
		Filename:   "./log/log.txt",
		MaxSize:    10, // 文件大小，單位是mb
		MaxBackups: 3,  //最大過期日誌保留天數
		MaxAge:     28, //保留紀錄天數
	}
	//同時輸出到文件及控制台
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(fileAndStdoutWriter)

}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		//
		body, _ := ioutil.ReadAll(c.Request.Body)
		//不要重複讀取body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()
		// w := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		// c.Writer = w
		c.Next()

		header := ""
		for k, v := range c.Request.Header {
			header += k + ": " + fmt.Sprint(v) + "\r\n"
		}

		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		errorMsg := c.Keys["ErrorMsg"]

		log.WithField("HTTPResponse", LogResponse{
			Time:   time.Now().Format("2006/01/02 15:04:05"),
			IP:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   path,
			Header: header,
			Body:   string(body),
			Status: c.Writer.Status(),
			// Response: w.body.String(),
			Error: errorMsg,
		}).Info(c.Request.Method, path, c.Writer.Status())
	}
}

type LogResponse struct {
	Time     string
	IP       string
	Method   string
	Path     string
	Header   string
	Body     string
	Status   int
	Response string
	Error    interface{} `json:"FailMsg,omitempty"`
}

// type bodyLogWriter struct {
// 	gin.ResponseWriter
// 	body *bytes.Buffer
// }

// func (w bodyLogWriter) Write(b []byte) (int, error) {
// 	w.body.Write(b)
// 	return w.ResponseWriter.Write(b)
// }
