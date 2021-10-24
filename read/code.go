package read

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ReadGiftPackCode 读取文件转SQL
func ReadGiftPackCode(file string, roomId int) string {
	// 小文件用 ioutil 读取
	open, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer open.Close()

	b, err := ioutil.ReadAll(open)
	if err != nil {
		panic(err)
	}
	codes := strings.Split(string(b), "\n")
	sql := "insert into `gift_pack` ( `room_id`,`activity_code` ) values "
	for _, c := range codes {
		sql += fmt.Sprintf("(%v, '%s'),", roomId, c)
	}
	sql = strings.TrimRight(sql, ",")
	sql += ";"
	return sql
}

// WriteFile 写入文件
func WriteFile(fileName string, content string) error {
	filePath := fmt.Sprintf("file/%s.sql", fileName)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	write := bufio.NewWriter(file)
	_, _ = write.WriteString(content)
	// 处理错误
	err = write.Flush()
	if err != nil {
		return err
	}
	return nil
}
