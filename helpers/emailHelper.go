package helpers

import (
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"net/smtp"
)

var (
	verifier = emailverifier.NewVerifier()
)

func EmailIsValid(email string) bool {

	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return false
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return false
	}

	fmt.Println("email validation result", ret)

	return true

}

//Code will be sent to email.
func SendEmail(code, mail string) {
	//put ur e-mail address that you want to sent e-mail by.
	from := "emailverifiy8@gmail.com"
	//put your email' password!!!
	pass := "emailonayla"

	to := []string{
		mail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("To: " + mail + "\r\n" +
		"Subject: Verification Code\r\n" +
		"\r\n" +
		"Hello dear,\r\n" + "Your code is\n" +
		code)

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Is Successfully sent.")

}
