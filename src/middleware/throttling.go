package middleware

import (
	"JuneGoBlog/src/util"
	"errors"
	"fmt"
	juneGin "github.com/520MianXiangDuiXiang520/GinTools/gin"
	utils "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/520MianXiangDuiXiang520/agingMap"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ThrottledFunc func(ctx *gin.Context) (interface{}, bool)

func BaseThrottled(f ThrottledFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		if resp, ok := f(context); !ok {
			context.Abort()
			context.JSON(http.StatusTooManyRequests, resp)
		}
	}
}

// 节流规则
type ThrottledRule int

const (
	// 只根据 IP 节流
	ThrottledRuleByIP ThrottledRule = iota + 1

	// 只根据 UA 节流
	ThrottledRuleByUserAgent

	// 根据 UA 和 IP 节流
	ThrottledRuleByUserAgentAndIP

	// 根据自定义字段节流, 若要使用自定义字段
	// 请使用 ThrottledCustom() 方法
	throttledRuleByCustomField
)

func getKey(rule ThrottledRule, ctx *gin.Context, customFields []string) string {
	switch rule {
	case ThrottledRuleByIP:
		return util.HashByMD5([]string{ctx.ClientIP()})
	case ThrottledRuleByUserAgent:
		return util.HashByMD5([]string{ctx.Request.UserAgent()})
	case ThrottledRuleByUserAgentAndIP:
		return util.HashByMD5([]string{ctx.ClientIP(), ctx.Request.UserAgent()})
	case throttledRuleByCustomField:
		fields := make([]string, len(customFields))
		for i, v := range customFields {
			fields[i] = ctx.Request.Header.Get(v)
		}
		return util.HashByMD5(fields)
	}
	return ""
}

func parseRate(rate string) (int, time.Duration) {
	r := strings.Split(rate, "/")
	if len(r) != 2 {
		msg := fmt.Sprintf("The rate string does not comply with the rules," +
			" please use a style similar to 1s")
		utils.ExceptionLog(errors.New("ParseFail"), msg)
		return -1, 0
	}
	n, d := r[0], r[1]
	frequency, err := strconv.Atoi(n)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse from string(%s) to int", n)
		utils.ExceptionLog(err, msg)
		return -1, 0
	}
	duration := map[string]time.Duration{"s": time.Second, "m": time.Minute,
		"h": time.Hour, "d": time.Hour * 24}[strings.ToLower(d)]
	return frequency, duration
}

var cache = agingMap.NewBaseAgingMap(time.Second*5, 0.5)

// 简单的节流中间件
func SimpleRateThrottle(rule ThrottledRule, rate string) ThrottledFunc {
	return throttled(rule, rate, nil)
}

func SimpleRateCustomFieldsThrottled(rate string, fields []string) ThrottledFunc {
	return throttled(throttledRuleByCustomField, rate, fields)
}

func throttled(rule ThrottledRule, rate string, fields []string) ThrottledFunc {
	return func(ctx *gin.Context) (interface{}, bool) {
		key := getKey(rule, ctx, fields)
		if key == "" {
			return nil, true
		}
		history, deadline, ok := cache.LoadWithDeadline(key)
		frequency, duration := parseRate(rate)
		if !ok {
			cache.Store(key, 1, duration)
			return nil, true
		}
		his := history.(int)

		if his >= frequency {
			// 直接拦截，不计入 cache
			resp := juneGin.BaseRespHeader{Code: http.StatusTooManyRequests,
				Msg: fmt.Sprintf("您的请求太快了，休息一下吧 ^_^ (%ds)", int64(deadline))}
			return resp, false
		}
		cache.Store(key, his+1, duration)
		return nil, true
	}
}
