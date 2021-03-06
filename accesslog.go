package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ua "github.com/mileusna/useragent"
)

type AccessLog struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func NewAccessLog(timeStamp time.Time, latency int64, path, os string) *AccessLog {
	return &AccessLog{
		Timestamp: timeStamp,
		Latency:   latency,
		Path:      path,
		OS:        os,
	}
}

// ログを取得するミドルウェア
// 引数：http.Handler型
// 戻り値：http.Handler型
func AccessLogger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		accessTimeBefore := time.Now()
		defer func() {
			ua := ua.Parse(r.UserAgent())
			accessTimeAfter := time.Now()
			accessTimeDiff := accessTimeAfter.Sub(accessTimeBefore).Microseconds()
			accessLog := NewAccessLog(accessTimeBefore, accessTimeDiff, r.URL.Path, ua.OS)
			accessLog.PrintJson()
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// http.HandlerFunc型のログを取得するミドルウェア
// 引数：http.HandlerFunc型
// 戻り値：http.HandlerFunc型
func AccessLoggerFunc(h http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		accessTimeBefore := time.Now()
		defer func() {
			ua := ua.Parse(r.UserAgent())
			accessTimeAfter := time.Now()
			accessTimeDiff := accessTimeAfter.Sub(accessTimeBefore).Microseconds()
			accessLog := NewAccessLog(accessTimeBefore, accessTimeDiff, r.URL.Path, ua.OS)
			accessLog.PrintJson()
		}()
		h(w, r)
	}
	return fn
}

func (a *AccessLog) PrintJson() {
	accessLogJson, err := json.Marshal(a)
	if err != nil {
		return
	}
	fmt.Println(string(accessLogJson))
}
