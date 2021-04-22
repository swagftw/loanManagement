package service

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(to string, name string, status string) {
	from := "swapnilnakade87@gmail.com"
	password := os.Getenv("APP_PASSWORD")
	toEmail := []string{
		to,
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	var body string
	fmt.Println(status)
	switch status {
	case "new":
		body = fmt.Sprintf("Hello,\n%s\nYour loan process has been initiated successfully.\n\nThank you.\nRegards.", name)
		break
	case "cancelled":
		body = fmt.Sprintf("Hello,\n%s\nYour loan process has been cancelled.\n\nThank you.\nRegards.", name)
		break
	case "approve":
		body = fmt.Sprintf("Hello,\n%s\nCongratulations your loan has been approved.\n\nThank you.\nRegards.", name)
		break
	case "reject":
		body = fmt.Sprintf("Hello,\n%s\nWe are sorry to say that your loan has been rejected.\n\nThank you.\nRegards.", name)
		break
	default:
	}

	message := []byte(body)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toEmail, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
