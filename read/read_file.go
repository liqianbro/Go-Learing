package read

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dataFile = "file/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func (f *Feed) ReadFeedByFile() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var feedList []*Feed
	err = json.NewDecoder(file).Decode(&feedList)
	return feedList, err
}

//ReadGiftPackCode 读取文件转SQL
func ReadGiftPackCode(file string, roomId int) string {
	// 检测文件
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(err)
	}
	// 小文件用 ioutil 读取
	open, err := os.Open(file)
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

// RemoveDuplicateElement
// 数组去重
func RemoveDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[strings.ToLower(item)]; !ok {
			temp[strings.ToLower(item)] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
