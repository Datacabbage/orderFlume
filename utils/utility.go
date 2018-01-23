package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"order_flume/lib/Error"
	"order_flume/lib/beelog"

	"github.com/satori/go.uuid"
	"gopkg.in/redis.v5"
)

var (
	ErrorParameter = errors.New("param is disabled ")
	Cache          *redis.Client
)

func SendToSZ(content []byte, url string) error {
	var (
		err  error
		body []byte
	)
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	body, err = HttpPost(url, string(content), header)
	if err != nil {
		//读取失败
		beelog.Error(string(content), string(body), err.Error())
		return err
	}

	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		beelog.Error(fmt.Sprintf("send: %v and response:%v convert to json fail,error:%v", string(content), string(body), err.Error()))
		return Error.GetDefError(131075)
	}
	if res.Is_ok == true {
		//发送成功
		beelog.Info(fmt.Sprintf("send:%v success", string(content)))
		return nil
	} else {
		//失败
		beelog.Error(fmt.Sprintf("send: %v and response:%v error:%v", string(content), string(body), res.Error))
		return err
	}
}

func HttpPost(url, params string, header http.Header) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(params))

	if err != nil {
		// handle error
		fmt.Println((err))
		return nil, err
	}
	req.Header = header

	beelog.Debug("req:", req)

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println((err))
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}

func HttpPostNo(url string) (error, []byte) {
	var bytes []byte
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(""))
	if err != nil {
		// handle error
		return err, bytes
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, bytes
	}
	return nil, body

}

func HttpGet(url string) (error, []byte) {
	var bytes []byte
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return err, bytes
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//	test := base64.StdEncoding.EncodeToString(body)
	//	bytes = []byte(test)
	//	if err != nil {
	//		// handle error
	//		return err, bytes
	//	}
	return nil, body
}

//CreateId create one id
func CreateId() string {
	return uuid.NewV4().String()
}
