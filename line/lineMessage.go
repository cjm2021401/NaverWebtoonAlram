package line

import (
	"NaverWebtoonAlram/config"
	"log"
	"os/exec"
)

func EndwebtoonMessageLine(endwebtoon string, dayofweek string)  {
	message:=dayofweek+"웹툰 "+endwebtoon+"이 완결났습니다. 유료로 전환될수 있으니 빠른시일내에 다시보는걸 추천드립니다."
	cmd := exec.Command("sh","linemessage.sh", config.Env.LineToken.Token, message)
	output , err :=cmd.Output()
	if err!=nil{
		panic(err)
	}
	log.Println(string(output))
}

func StartwebtoonSendLine(startwebtoon string, dayofweek string)  {
	message:=dayofweek+"웹툰 "+startwebtoon+"이 연재를 시작했습니다. 많은 관심 부탁드립니다."
	cmd := exec.Command("sh","linemessage.sh", config.Env.LineToken.Token, message)
	output , err :=cmd.Output()
	if err!=nil{
		panic(err)
	}
	log.Println(string(output))
}

