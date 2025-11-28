package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Request struct {
	WebhookURL string  `json:"webhook_url" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

type WebhookPayload struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}

var paymentStatuses = []string{
	"APPROVED",
	"CANCELED",
	"REJECTED",
}

func main() {
	apiPort := os.Getenv("API_PORT")

	app := gin.Default()

	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	app.POST("/payments", func(ctx *gin.Context) {
		var request Request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paymentID := generatePaymentID()
		status := generatePaymentStatus()

		go triggerWebhook(request.WebhookURL, paymentID, status)

		ctx.JSON(http.StatusOK, gin.H{
			"payment_id": paymentID,
		})
	})

	if apiPort == "" {
		apiPort = "8080"
	}

	if err := app.Run(":" + apiPort); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}

func triggerWebhook(url, paymentID, status string) {
	time.Sleep(5 * time.Second)

	payload := WebhookPayload{
		PaymentID: paymentID,
		Status:    status,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("erro ao enviar webhook:", err)
		return
	}

	rawResponse, _ := io.ReadAll(resp.Body)

	log.Printf("webhook enviado. HTTP Status Code: %d Response: %s", resp.StatusCode, string(rawResponse))
}

func generatePaymentID() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return time.Now().Format("20060102150405")
}

func generatePaymentStatus() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(paymentStatuses))
	return paymentStatuses[randomIndex]
}
