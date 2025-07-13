package api

import (
	"american-empire/backend/database"
	"american-empire/backend/models"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func saveUploadedPhoto(c *gin.Context, file *multipart.FileHeader, eventID uint) (string, error) {
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	if os.Getenv("GIN_MODE") == "release" {
		path := fmt.Sprintf("photos/%d/%s", eventID, filename)

		src, err := file.Open()
		if err != nil {
			return "", err
		}
		defer src.Close()

		_, err = database.SupabaseClient.Storage.UploadFile("photos", path, src)
		if err != nil {
			return "", err
		}

		return path, nil
	} else {
		dir := fmt.Sprintf("data/photos/%d", eventID)
		os.MkdirAll(dir, 0755)
		fullPath := fmt.Sprintf("%s/%s", dir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			return "", err
		}

		return fmt.Sprintf("%d/%s", eventID, filename), nil
	}
}

func notifyAdmin(email string, event *models.Event) {
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := os.Getenv("GMAIL_USER")
	smtpPassword := os.Getenv("GMAIL_PASS")

	from := smtpUsername
	to := smtpUsername
	subject := fmt.Sprintf("%d: New Event Submission", event.ID)
	body := fmt.Sprintf("Title: %s\nSubmitted by: %s", event.Title, email)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s",
		email, subject, body)

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		from,
		[]string{to},
		[]byte(message),
	)

	if err != nil {
		fmt.Println(err)
	}
}
