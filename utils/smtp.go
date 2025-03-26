package utils

import (
	"net/smtp"
)

// SendEmail отправляет электронное письмо с сообщением о регистрации
func SendEmail(to string, subject string, body string) error {
	from := "vsrs-rs@yandex.com"   // Убедитесь, что вы установили переменную окружения SMTP_FROM
	password := "hecshrwpjigzeokn" // Убедитесь, что вы установили переменную окружения SMTP_PASSWORD

	smtpHost := "smtp.yandex.ru" // Замените на ваш SMTP сервер
	smtpPort := "587"            // Замените на ваш SMTP порт

	// Формируем сообщение
	message := []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	// Аутентификация
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Отправка письма
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}
