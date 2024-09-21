package sendto

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"go.uber.org/zap"
)

const (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = "587"
	from     = "kaitok57a01@gmail.com"
	password = "erjc rwah nonq ffpr"
)

// kzil sryc bdec ykvt
type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mai Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mai.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mai.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mai.Subject)
	msg += fmt.Sprintf("\r\n%s", mai.Body)
	return msg
}

// func SendEmailOTP(to []string, from string, otp string) error {
// 	contentEmail := Mail{
// 		From:    EmailAddress{Address: from, Name: "Admin"},
// 		To:      to,
// 		Subject: "OTP",
// 		Body:    "Your OTP is " + "otp. Please use it within 10 minutes.",
// 	}
// 	messageMail := BuildMessage(contentEmail)
// 	auth := smtp.PlainAuth("", from, password, SMTPHost)
// 	fmt.Println("to: ", to[0])
// 	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(messageMail))
// 	if err != nil {
// 		global.Logger.Error("Error sending email: ", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

func SendEmailOTP(to []string,from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "Admin"},
		To:      to,
		Subject: "OTP",
		Body:    "Your OTP is " + otp + ". Please use it within 10 minutes.",
	}
	messageMail := BuildMessage(contentEmail)

	// Establish a connection to the SMTP server without TLS first
	serverName := SMTPHost + ":" + SMTPPort
	conn, err := net.Dial("tcp", serverName)
	if err != nil {
		fmt.Println("Error establishing TCP connection:", err)
		global.Logger.Error("Error establishing TCP connection: ", zap.Error(err))
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, SMTPHost)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		global.Logger.Error("Error creating SMTP client: ", zap.Error(err))
		return err
	}
	defer client.Quit()

	// Upgrade the connection to TLS using STARTTLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Only use in development, not in production
		ServerName:         SMTPHost,
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		fmt.Println("Error starting TLS:", err)
		global.Logger.Error("Error starting TLS: ", zap.Error(err))
		return err
	}

	// Authentication
	auth := smtp.PlainAuth("", from, password, SMTPHost)
	if err = client.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		global.Logger.Error("Error authenticating: ", zap.Error(err))
		return err
	}

	// Set the sender and recipients
	if err = client.Mail(from); err != nil {
		fmt.Println("Error setting sender:", err)
		global.Logger.Error("Error setting sender: ", zap.Error(err))
		return err
	}

	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			fmt.Println("Error setting recipient:", err)
			global.Logger.Error("Error setting recipient: ", zap.Error(err))
			return err
		}
	}

	// Send the email body
	wc, err := client.Data()
	if err != nil {
		fmt.Println("Error starting data transfer:", err)
		global.Logger.Error("Error starting data transfer: ", zap.Error(err))
		return err
	}

	_, err = wc.Write([]byte(messageMail))
	if err != nil {
		fmt.Println("Error writing message:", err)
		global.Logger.Error("Error writing message: ", zap.Error(err))
		return err
	}

	err = wc.Close()
	if err != nil {
		fmt.Println("Error closing data writer:", err)
		global.Logger.Error("Error closing data writer: ", zap.Error(err))
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}


