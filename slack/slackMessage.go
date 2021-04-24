package slack

import (
	"NaverWebtoonAlram/config"
	"NaverWebtoonAlram/line"
	"bytes"
	"net/http"
)


func SendSlackEndWebtoonImage(EndwebtoonName string, EndwebtoonImage string,dayofweek string)  {
	reqBody := bytes.NewBufferString("{\"text\":\" "+dayofweek+"웹툰 "+EndwebtoonName+"이 완결났습니다. 유료로 전환될수 있으니 빠른시일내에 다시보는걸 추천드립니다.\"}")
	resp, err := http.Post(config.Env.SlackWebhook.Url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	line.EndwebtoonMessageLine(EndwebtoonName, EndwebtoonImage, dayofweek)
}

func SendSlackStartWebtoonImage(StartwebtoonName string, StartwebtoonImage string, dayofweek string)  {
	reqBody := bytes.NewBufferString("{\"text\":\" "+dayofweek+"웹툰 "+StartwebtoonName+"이 연재를 시작했습니다. 많은 관심 부탁드립니다. \"}")
	resp, err := http.Post(config.Env.SlackWebhook.Url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	line.StartwebtoonSendLine(StartwebtoonName,StartwebtoonImage,dayofweek)
}