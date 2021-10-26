package sms

import (
	"context"
	"errors"
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type AliyunSMSService struct{}

func (s *AliyunSMSService) SendSMS(ctx context.Context, phone, code string) error {
	accessKeyId := "key"
	accessKeySecret := "secret"

	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return err
	}

	codeJson := fmt.Sprintf("{\"code\":%s}", code)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(codeJson),
		TemplateCode:  tea.String("SMS_215185275"),
		SignName:      tea.String("好说w"),
	}
	// 复制代码运行请自行打印 API 的返回值
	smsResp, err := client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}
	if *smsResp.Body.Code != "OK" {
		return errors.New(*smsResp.Body.Message)
	}
	fmt.Println(smsResp.String())
	return nil
}
