package verified

import "testing"

func Test_realNameVerified_UserRealNameVerified(t *testing.T) {
	type args struct {
		realName string
		idCard   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"测试用例1",
			args{realName: "1111111", idCard: "11111111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &realNameVerified{}
			r.UserRealNameVerified(tt.args.realName, tt.args.idCard)
		})
	}
}
