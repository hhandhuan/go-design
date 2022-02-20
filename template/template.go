package template

import "fmt"

type ISMS interface {
	Send(content string, phone int) error
}

type Sms struct {
	ISMS
}

// Valid 校验短信字数
func (s *Sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

// Send 发送短信
func (s *Sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}
	// 调用子类的方法发送短信
	return s.Send(content, phone)
}
