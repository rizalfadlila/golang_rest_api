package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/rest_api/pkg/logger"

	"github.com/gin-gonic/gin"
)

// RequestLoggerMiddleware logs request based on specific configuration
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body interface{}
		if ctx.Request.Body != nil {
			bodyReader := ctx.Request.Body
			bodyBytes, _ := ioutil.ReadAll(bodyReader)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			var bodyMap map[string]interface{}
			json.Unmarshal(bodyBytes, &bodyMap)
			body = bodyMap
		}

		var st time.Time
		var lt time.Duration

		st = time.Now()

		ctx.Next()

		lt = time.Since(st)
		logger.AccessLog(
			logger.SetMessageFormat("%s %s", ctx.Request.URL.Path, ctx.Request.Method),
			logger.Any("tag", "go-access"),
			logger.Any("http.path", ctx.Request.URL.Path),
			logger.Any("http.method", ctx.Request.Method),
			logger.Any("http.body", body),
			logger.Any("http.query-param", ctx.Request.URL.Query()),
			logger.Any("http.agent", ctx.Request.UserAgent()),
			logger.Any("http.referer", ctx.Request.Referer()),
			logger.Any("http.status", ctx.Writer.Status()),
			logger.Any("http.latency", lt.Seconds()),
		)
	}
}
