package email_notification

import (
    "fmt"
    "log"
    "net/smtp"
    "strings"
)

type Mail struct {
    Sender  string
    To      []string
    Cc      []string
    Bcc     []string
    Subject string
    Body    string
}

func SendEmailWIthBCCExample() {

    sender := "john.doe@example.com"

    to := []string{
        "roger.roe@example.com",
        "adam.smith@example.com",
        "thomas.wayne@example.com",
        "oliver.holmes@example.com",
    }

    cc := []string{
        "adam.smith@example.com",
        "thomas.wayne@example.com",
    }

    // not used
    bcc := []string{
        "oliver.holmes@example.com",
    }

    user := "9c1d45eaf7af5b"
    password := "ad62926fa75d0f"

    subject := "simple testing mail"
    body := "email body message"

    request := Mail{
        Sender:  sender,
        To:      to,
        Cc:      cc,
        Subject: subject,
        Body:    body,
    }

    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    msg := BuildMessage(request)
    auth := smtp.PlainAuth("", user, password, host)

    err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Emails sent successfully")
}

func BuildMessage(mail Mail) string {

    msg := ""
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)

    if len(mail.To) > 0 {
        msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
    }

    if len(mail.Cc) > 0 {
        msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
    }

    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}
