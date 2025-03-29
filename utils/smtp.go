package utils

import (
	"fmt"
	"net/smtp"
)

// SendEmail отправляет HTML-письмо с кнопкой подтверждения регистрации
func SendEmail(to string, url string) error {
	subject := "Подтверждение регистрации"

	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Подтверждение регистрации</title>
	</head>
	<body style="font-family: Arial, sans-serif; text-align: center;">
		<h2 style="color: #333;">Добро пожаловать!</h2>
		<p style="font-size: 16px; color: #555;">Пожалуйста, подтвердите вашу регистрацию, нажав на кнопку ниже.</p>
		<a href="%s" style="display: inline-block; padding: 10px 20px; font-size: 18px; color: #fff; background-color: #007BFF; text-decoration: none; border-radius: 5px;">Подтвердить регистрацию</a>
		<p style="margin-top: 20px; font-size: 14px; color: #999;">Если кнопка не работает, скопируйте и вставьте следующую ссылку в браузер:<br>%s</p>
	</body>
	</html>`, url, url)

	from := "vsrs-rs@mail.ru"
	password := "VwuD7ynRjauwecAN2PHs"

	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	// Формируем сообщение в формате MIME
	message := []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=UTF-8\n\n" +
		body)

	// Аутентификация
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Отправка письма
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
}
