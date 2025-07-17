package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"upload-service/pkg/common"
)

type responseWriter struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

type LogEntry struct {
	Timestamp    time.Time `json:"timestamp"`
	Method       string    `json:"method"`
	URL          string    `json:"url"`
	ResponseBody string    `json:"response_body"`
	UserID       string    `json:"user_id,omitempty"`
	StatusCode   int       `json:"status_code"`
	IPAddress    string    `json:"ip_address"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			body:           bytes.NewBufferString(""),
			statusCode:     200,
		}

		next.ServeHTTP(rw, r)

		userID := ""
		if userIDCtx := r.Context().Value(common.UserIDContextKey); userIDCtx != nil {
			if uid, ok := userIDCtx.(string); ok {
				userID = uid
			}
		}

		ipAddress := r.Header.Get("X-Forwarded-For")
		if ipAddress == "" {
			ipAddress = r.Header.Get("X-Real-IP")
		}
		if ipAddress == "" {
			ipAddress = r.RemoteAddr
		}

		logEntry := LogEntry{
			Timestamp:    start,
			Method:       r.Method,
			URL:          r.URL.String(),
			ResponseBody: rw.body.String(),
			UserID:       userID,
			StatusCode:   rw.statusCode,
			IPAddress:    ipAddress,
		}

		logJSON, err := json.MarshalIndent(logEntry, "", "  ")
		if err != nil {
			log.Printf("Error marshaling log entry: %v", err)
			return
		}

		fmt.Fprintln(os.Stdout, string(logJSON))
	})
}
