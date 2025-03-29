package utils

import (
	"backend/internal/config"
	"fmt"
	"net/smtp"
)

// SendEmail отправляет HTML-письмо с кнопкой подтверждения регистрации
func SendEmail(to string, url string, cfg *config.SMTPConfig) error {
	if cfg == nil {
		return fmt.Errorf("SMTP configuration is nil")
	}

	subject := "Подтверждение регистрации"

	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Подтверждение регистрации</title>
	</head>
	<body style="font-family: 'Helvetica Neue', Arial, sans-serif; margin: 0; padding: 0; background-color: #f4f7fa; color: #333;">
		<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 10px rgba(0,0,0,0.1); margin-top: 20px; margin-bottom: 20px;">
			<!-- Header -->
			<div style="background-color: #1a73e8; padding: 24px; text-align: center;">
				<h1 style="color: #ffffff; margin: 0; font-size: 24px; font-weight: 500;">Подтверждение регистрации</h1>
			</div>
			
			<!-- Content -->
			<div style="padding: 30px 40px;">
				<h2 style="color: #1a73e8; font-weight: 500; margin-top: 0;">Добро пожаловать!</h2>
				<p style="font-size: 16px; line-height: 1.6; color: #555; margin-bottom: 25px;">Благодарим вас за регистрацию. Пожалуйста, нажмите на кнопку ниже, чтобы подтвердить вашу учетную запись.</p>
				
				<!-- Button -->
				<div style="text-align: center; margin: 30px 0;">
					<a href="%s" style="display: inline-block; padding: 12px 28px; font-size: 16px; color: #ffffff; background-color: #1a73e8; text-decoration: none; border-radius: 4px; font-weight: 500; box-shadow: 0 2px 5px rgba(0,0,0,0.1); transition: all 0.2s ease-in-out;">Подтвердить регистрацию</a>
				</div>
				
				<p style="font-size: 14px; line-height: 1.6; color: #777; margin-top: 25px;">Если кнопка не работает, вы можете скопировать и вставить следующую ссылку в адресную строку вашего браузера:</p>
				<p style="font-size: 14px; background-color: #f5f5f5; padding: 10px; border-radius: 4px; word-break: break-all; margin-bottom: 0;"><a href="%s" style="color: #1a73e8; text-decoration: none;">%s</a></p>
			</div>
			
			<!-- Footer -->
			<div style="background-color: #f5f7fa; padding: 20px; text-align: center; border-top: 1px solid #e0e0e0;">
				<p style="font-size: 13px; color: #999; margin: 0;">Это автоматическое сообщение, пожалуйста, не отвечайте на него.</p>
				<p style="font-size: 13px; color: #999; margin: 10px 0 0 0;">&copy; 2025 VSRS-RS. Все права защищены.</p>
			</div>
		</div>
	</body>
	</html>`, url, url, url)

	// Формируем сообщение в формате MIME
	message := []byte("From: " + cfg.From + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=UTF-8\n\n" +
		body)

	// Аутентификация
	auth := smtp.PlainAuth("", cfg.From, cfg.Password, cfg.Host)

	// Отправка письма
	return smtp.SendMail(cfg.Host+":"+cfg.Port, auth, cfg.From, []string{to}, message)
}
