package models

import (
	"fmt"

	database "github.com/biseshbhattarai/talwar/db"
	"github.com/go-mail/mail"
)

type Email struct {
	Id       int64 `gorm:"primaryKey"`
	Email    string
	ScanType ScanType
}

func (email *Email) Save() (*Email, error) {
	err := database.DbConn.Create(&email).Error
	if err != nil {
		return &Email{}, err
	}
	return email, nil
}

func ListEmails() ([]Email, error) {
	var emails []Email
	err := database.DbConn.Find(&emails).Error
	if err != nil {
		return []Email{}, err
	}
	return emails, nil
}

func SendGridSendEmail(email Email, content string) {
	m := mail.NewMessage()
	m.SetHeader("From", "biseshbhattaraiii@gmail.com")
	m.SetHeader("To", email.Email)
	m.SetHeader("Subject", "Scan Completed")
	m.SetBody("text/html", fmt.Sprintf("Hello ! <br/> <h3>%s</h3>", content))
	d := mail.NewDialer("smtp.gmail.com", 587, "biseshbhattaraiii@gmail.com", "jeajnzxjzebrfduc")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func SendEmails(emails []Email, content string) {
	for _, email := range emails {
		go func(email Email) {
			SendGridSendEmail(email, content)
		}(email)
	}
}

func StartSendEmails(content string) {
	targets, err := ListEmails()
	if err != nil {
		fmt.Println(err)
	}
	SendEmails(targets, content)
}
