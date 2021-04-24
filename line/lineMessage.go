package line

import (
	"NaverWebtoonAlram/config"
	"log"
	"os/exec"
)

func EndwebtoonMessageLine(EndwebtoonName string, EndwebtoonImage string, dayofweek string)  {
	message:=dayofweek+"웹툰 "+EndwebtoonName+"이 완결났습니다. 유료로 전환될수 있으니 빠른시일내에 다시보는걸 추천드립니다."
	cmd := exec.Command("sh","linemessage.sh", config.Env.LineToken.Token, message, EndwebtoonImage)
	output , err :=cmd.Output()
	if err!=nil{
		panic(err)
	}
	log.Println(string(output))
}

func StartwebtoonSendLine(StartwebtoonName string, StartwebtoonImage string, dayofweek string)  {
	message:=dayofweek+"웹툰 "+StartwebtoonName+"이 연재를 시작했습니다. 많은 관심 부탁드립니다."
	cmd := exec.Command("sh","linemessage.sh", config.Env.LineToken.Token, message, StartwebtoonImage)
	output , err :=cmd.Output()
	if err!=nil{
		panic(err)
	}
	log.Println(string(output))
}