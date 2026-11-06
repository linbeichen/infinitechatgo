package utils

import (
	"crypto/tls"
	"fmt"
	"infinitechatgo/service/user/api/internal/types/constants"
	"net/smtp"
)

// 发送验证码
// targetEmail 目标邮箱
// authCode 验证码
func SendEmailCode(targetEmail, authCode string) error {
	subject := constants.EmailSubject
	//邮箱的内容
	body := fmt.Sprintf("您的验证码为:%s(五分钟内有效)", authCode)

	// 验证码邮件内容
	msg := []byte("From: " + constants.EmailName + " <" + constants.EmailUserName + ">\r\n" +
		"To: " + targetEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		body + "\r\n")

	// 通过 smtp.PlainAuth() 创建认证对象
	auth := smtp.PlainAuth("", constants.EmailUserName, constants.EmailPassword, constants.EmailHostName)

	// tls 配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         constants.EmailHostName,
		MinVersion:         tls.VersionTLS12,
	}

	// 建立 SSL 连接 (465端口)
	// 465 是 smtp 隐式 tls 端口，从连接开始就是加密的。
	conn, err := tls.Dial("tcp", constants.EmailHostName+":465", tlsConfig)
	if err != nil {
		return fmt.Errorf("邮件连接失败: %w", err)
	}
	defer conn.Close()

	// 创建 smtp 客户端
	client, err := smtp.NewClient(conn, constants.EmailHostName)
	if err != nil {
		return fmt.Errorf("创建smtp客户端失败: %w", err)
	}
	defer client.Close()

	// 登录 smtp 服务器
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("邮件认证失败: %w", err)
	}

	// 设置发件人
	if err := client.Mail(constants.EmailUserName); err != nil {
		return fmt.Errorf("设置发件人失败: %w", err)
	}

	// 设置收件人
	if err := client.Rcpt(targetEmail); err != nil {
		return fmt.Errorf("设置收件人失败: %w", err)
	}

	// 获取邮件写入流
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("获取邮件写入流失败: %w", err)
	}

	// 写入邮件内容
	if _, err := w.Write(msg); err != nil {
		return fmt.Errorf("写入邮件内容失败: %w", err)
	}

	// 关闭邮件写入流
	// Close() 结束邮件数据写入，并通知 SMTP 服务器提交邮件
	if err := w.Close(); err != nil {
		return fmt.Errorf("关闭邮件写入流失败: %w", err)
	}

	return client.Quit()
}
