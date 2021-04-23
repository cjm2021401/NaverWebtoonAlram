package slack

import (
	"NaverWebtoonAlram/config"
	"bytes"
	"net/http"
)

func SendSlackEndWebtoon(endwebtoon string, dayofweek string)  {
	reqBody := bytes.NewBufferString("{\"text\":\" 종료웹툰 : "+endwebtoon +"(" + dayofweek + ") \n약 4개월 뒤 유료화 됨으로 지금 보세요!\"}")
	resp, err := http.Post(config.Env.SlackWebhook.Url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func SendSlackStartWebtoon(startwebtoon string, dayofweek string)  {
	reqBody := bytes.NewBufferString("{\"text\":\" 새로운 웹툰 : "+startwebtoon +"(" + dayofweek + ") \n새로운 웹툰 많은 관심 부탁드립니다! \"}")
	resp, err := http.Post(config.Env.SlackWebhook.Url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}