package utils

import (
	"fmt"
	"math/rand"
	// "net/smtp"
	"strconv"

)

func SendCode(email string) string{
	code := strconv.Itoa(rand.Intn(10000))
	// smtpHost := "smtp.gmail.com"
	// smtpPort := "587"
	// fromEmail := "simplefinnaceplanner@gmail.com"
	// password := "phwn dpbx ivmu laju" // Use App Passwords, NOT your real password

	// // Message body
	// msg := "From: " + fromEmail + "\n" +
	// 	"To: " + email + "\n" +
	// 	"Subject: " + "Your Verification Code" + "\n\n" +
	// 	"Your verification code is: " + code

	// // Authentication
	// auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	// // Send email
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, []string{email}, []byte(msg))
	// if err != nil {
	// 	fmt.Println("Error sending email:", err)
	// 	return err
	// }

	// fmt.Println("Email sent successfully!")
	fmt.Println(code)
	return code
}