package sms_test

import (
	"Go-Learing/sms"
	"context"
	"testing"
)

func TestAliyunSMSService_SendSMS(t *testing.T) {
	type args struct {
		ctx   context.Context
		phone string
		code  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			args: args{ctx: context.TODO(), phone: "16621056741", code: "123456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sms.AliyunSMSService{}
			if err := s.SendSMS(tt.args.ctx, tt.args.phone, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("SendSMS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
