package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/DILNATHRK/GoWithAzureCommunications/SendEmail/helper"
	"github.com/DILNATHRK/GoWithAzureCommunications/SendEmail/model"
	"github.com/joho/godotenv"
)

func SendEmailWithACS(email model.EmailNotification) error {
	godotenv.Load(".env")

	host := os.Getenv("Host")
	pathAndQuery := os.Getenv("PathAndQuery")
	secret := os.Getenv("Secret")
	resourceEndpoint := "https://" + host
	senderAddress := os.Getenv("SenderAddress")
	requestURI := resourceEndpoint + pathAndQuery

	body := map[string]interface{}{
		"senderAddress": senderAddress,
		"recipients": map[string]interface{}{
			"to": []map[string]string{
				{"address": email.Recipient},
			},
		},
		"content": map[string]string{
			"subject":   email.Subject,
			"plainText": email.Payload,
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling body: %w", err)
	}

	date := helper.FormatDate(time.Now().UTC())
	contentHash := helper.ComputeContentHash(jsonData)

	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", pathAndQuery, date, host, contentHash)
	signature, err := helper.ComputeSignature(stringToSign, secret)
	if err != nil {
		return fmt.Errorf("error computing signature: %w", err)
	}
	authHeader := fmt.Sprintf("HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s", signature)

	req, err := http.NewRequest("POST", requestURI, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-content-sha256", contentHash)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email, status: %d", resp.StatusCode)
	}

	return nil
}
