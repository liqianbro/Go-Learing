package read

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFeed_ReadFeedByFile(t *testing.T) {
	type fields struct {
		Name string
		URI  string
		Type string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*Feed
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Feed{
				Name: tt.fields.Name,
				URI:  tt.fields.URI,
				Type: tt.fields.Type,
			}
			got, err := f.ReadFeedByFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFeedByFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFeedByFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{
			name: "读取文件测试",
			args: args{file: "file/code.txt", roomId: 1},
			want: "insert into `gift_pack` ( `room_id`,`activity_code` ) values (1, 'UP1237LD7MF39121');",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadGiftPackCode(tt.args.file, tt.args.roomId); got != tt.want {
				t.Errorf("ReadGiftPackCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateElement(t *testing.T) {
	//s := []string{"hs1MZQIO", "hs1mZQIO", "hello", "golang", "hello", "ruby", "php", "java", "PHP"}

	// 小文件用 ioutil 读取
	open, err := os.Open("/Users/cwcw/Desktop/Learn/Go-Learing/file/ss.txt")
	if err != nil {
		panic(err)
	}
	defer open.Close()

	var codes []string
	// 每行读取
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(len(RemoveDuplicateElement(codes))) //output: hello world golang ruby php java

}
