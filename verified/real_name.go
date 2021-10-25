package verified

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	verifiedUrl = "http://checkone.market.alicloudapi.com/chinadatapay/1882"
	appCode     = "123456"
)

var RealNameVerified = new(realNameVerified)

type realNameVerified struct{}

type VerifiedResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Result string `json:"result"`
	} `json:"data"`
	SeqNo string `json:"seqNo"`
}

type UserVerified struct {
	Id        int
	UserId    int
	RealName  string
	IdCard    int
	CreatedAt time.Time
}

//UserRealNameVerified
//  @Description: 调用阿里云实名认证
//  @receiver r
//  @param realName
//  @param idCard
//
func (r *realNameVerified) UserRealNameVerified(realName, idCard string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", verifiedUrl, nil)
	if err != nil {
		return
	}
	// 组装Query参数
	params := make(url.Values)
	params.Add("name", realName)
	params.Add("idcard", idCard)
	req.URL.RawQuery = params.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "APPCODE", appCode))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return
	}
	var result VerifiedResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(result)

	switch result.Data.Result {
	case "1":
		fmt.Println("实名验证成功")
	case "2":
		fmt.Println("实名验证失败")
	case "3":
		fmt.Println("实名验证出异常")
	}

}
