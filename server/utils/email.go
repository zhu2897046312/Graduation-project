package utils

import (
	"fmt"
	"log"
	"server/config" // 替换为你的实际项目路径

	"gopkg.in/gomail.v2"
)

// SendEmail 发送邮件函数
func SendEmail(to string, subject string, message string) error {
	emailConfig := config.GlobalConfig.Email
	
	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message) // 使用HTML格式，如果需要纯文本可改为"text/plain"

	// 创建邮件拨号器
	d := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.Username,
		emailConfig.Password,
	)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Printf("邮件发送失败: %v", err)
		return fmt.Errorf("邮件发送失败: %v", err)
	}

	log.Printf("邮件发送成功: %s -> %s", emailConfig.From, to)
	return nil
}

// SendTextEmail 发送纯文本邮件
func SendTextEmail(to string, subject string, message string) error {
	emailConfig := config.GlobalConfig.Email
	
	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message) // 纯文本格式

	d := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.Username,
		emailConfig.Password,
	)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("邮件发送失败: %v", err)
		return fmt.Errorf("邮件发送失败: %v", err)
	}

	log.Printf("纯文本邮件发送成功: %s -> %s", emailConfig.From, to)
	return nil
}

// SendEmailWithAttachment 发送带附件的邮件
func SendEmailWithAttachment(to string, subject string, message string, attachmentPath string) error {
	emailConfig := config.GlobalConfig.Email
	
	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	
	// 添加附件
	m.Attach(attachmentPath)

	d := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.Username,
		emailConfig.Password,
	)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("带附件邮件发送失败: %v", err)
		return fmt.Errorf("带附件邮件发送失败: %v", err)
	}

	log.Printf("带附件邮件发送成功: %s -> %s", emailConfig.From, to)
	return nil
}