package read_test

import (
	"Go-Learing/read"
	"testing"
)

func TestReadGiftPackCode(t *testing.T) {
	type args struct {
		file   string
		roomId int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "读取文件测试",
			args: args{file: "file/code.txt", roomId: 1},
			want: "insert into `gift_pack` ( `room_id`,`activity_code` ) values (1, 'UP1237LD7MF39121');",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := read.ReadGiftPackCode(tt.args.file, tt.args.roomId); got != tt.want {
				t.Errorf("ReadGiftPackCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
