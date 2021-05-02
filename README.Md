# Naver Webtoon Alram
- 네이버 신규/완결 웹툰 알람

## 1. WHAT IS NAVER WEBTOON ALRAM
- 네이버에서 신규/완결 웹툰이 생길시 라인/slack으로 알려주는 서비스
- 네이버의 경우 웹툰이 완결난지 약 4개월이 지나면 유료웹툰으로 변함    
  -> 몰아보고싶어 모아두었던 웹툰을 못보는 일 발생
- 매일 데이터를 불러와 기존의 데이터와 분석하여 신규/완결 웹툰을 획득 후 정보 전송


### 2. 구동원리
- POSTGRESQL을 이용하여 데이터베이스 구축
- GO(GIN)를 이용한 웹서버 구축, DB와 연결 및 값가져오기 / 내보내기
- SLACK, LINE open api 이용한 메시지 전송 
- DOCKER를 이용한 서버 실행


### 3. 요구사항
- POSTGRESQl 설치
- GO 설치
- DOCKER for MAC / WINDOW 설치
- SLACK Web Hook 발급 (https://api.slack.com/messaging/webhooks) 
- LINE Notify personal access token 발급
  
    (https://access.line.me/dialog/oauth/weblogin?response_type=code&client_id=1476232700&redirect_uri=https%3A%2F%2Fnotify-bot.line.me%2Flogin%2Fcallback&state=ft1jgkvqk4ZsFPk6Crm4qZ, 라인로그인필요)

### 4. 기본 설정
>####초기 데이터 베이스 설정
><pre><code>
>CREATE monday(
>    id    SERIAL NOT NULL
>    toon  VARCHAR(30),
>    image VARCHAR(150) DEFAULT 'https://image.chosun.com/sitedata/image/202006/24/2020062400801_0.png',
>    PRIMARY KEY(id)
>);
>
> 다른요일도 동일하게 설정
></code></pre>
>
>#### 데이터 입력
>- 초기 데이터입력은 주석처리된 코드를 풀어서 실행해주어야(cron.go)
><pre><code>
>func SetupCron(location string) (*cron.Cron, error) {
>//ReadWriteAllday()  //처음 세팅할때만 돌림
>loc, err := time.LoadLocation(location)
>if err != nil {
>return nil, err
>}
>c := cron.New(cron.WithLocation(loc))
>_,err=c.AddFunc("0 8 * * *",  DailydataCheck)//매일 8시에 확인
>if err != nil {
>return nil, err
>}
>
>	return c, nil
>}           
>
>           
></code></pre>
>#### CONFIG 
> <pre><code>
> {
>"database": {
>"url": "데이터베이스 url : port",
>"user": "데이터베이스 User",
>"password": "데이터베이스  Password",
>"name": "naverwebtoon"
> },
>"naverwebtoon": {
>"url":"네이버 웹툰 메인 홈페이지"
> },
>"slackwebhook" : {
>"url" : "YOUR SLACK WEB HOOK"
> },
>"linetoken" : {
>"token" : "YOUR Notify personal access token"
> }
> }
> </code></pre>

### 5. 구동
                docker build -t <TAG> .
                dockr run -p 8080:8080 <TAG>

### 6. 예시
![EXAMPLE(완결)](./img/IMG_0082.PNG)

### 문의 
- cjm2021401@khu.ac.kr