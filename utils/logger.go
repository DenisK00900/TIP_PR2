package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

type contextKey string

const RequestIDKey contextKey = "request-id"

func NewID16() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func LogRequest(r *http.Request) {
	reqID := "no-id"
	if id, ok := r.Context().Value(RequestIDKey).(string); ok {
		reqID = id
	}
	
	fmt.Printf("[%s] [%s] %s %s %s\n",
		time.Now().Format(time.RFC3339),
		reqID,
		r.RemoteAddr,
		r.Method,
		r.URL.Path,
	)
}

func LogInfo(msg string) {
	fmt.Printf("[INFO] %s %s\n", time.Now().Format(time.RFC3339), msg)
}

func LogError(msg string) {
	fmt.Printf("[ERROR] %s %s\n", time.Now().Format(time.RFC3339), msg)
}