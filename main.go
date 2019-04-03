package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//Response format.
type Response struct {
	TrackingCode string      `json:"tracking_code"`
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	Env          string      `json:"env"`
	Data         interface{} `json:"data"`
}

var pool = "abcdefgthgklmnopqrstuvxyz1234567890"

// Generate a random string of A-Z chars with len = l
func randomString(l int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{
			TrackingCode: randomString(16),
			Status:       `success`,
			Message:      `Test k8s with golang resposne json.`,
			Env:          os.Getenv("SLACK_TOKEN"),
		}
		bytes, _ := json.Marshal(resp)
		w.Write(bytes)
		fmt.Println(resp)
	})
	http.ListenAndServe(":80", nil)
}
