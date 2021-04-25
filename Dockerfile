### Builder
FROM golang:1.15-alpine as builder
RUN apk update && apk add git && apk add ca-certificates && apk add tzdata

WORKDIR /usr/src/app
COPY . .
ARG ENV
ENV env $ENV
# CGO_ENABLED=0 : cgo를 사용하지 않습니다. Scratch 이미지에는 C 바이너리조차 없기 때문에, 반드시 cgo를 비활성화 후 빌드해야합니다.
# GOOS=linux GOARCH=amd64 : OS와 아키텍쳐 설정입니다.
# -a : 모든(all) 의존 패키지를 cgo를 사용하지 않도록 재빌드합니다.
# -ldflags '-s' : 바이너리를 조금 더 경량화하는 Linker 옵션입니다.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .

### Make executable image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/main /main
COPY --from=builder /usr/src/app/config /config
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 50000


CMD [ "/main" ]