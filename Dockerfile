FROM golang:1.6.2

# for gateway
ENV SERVICE_NAME datahub_custom

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

ENV SERVICE_PORT 8080
EXPOSE 8080

ENV SERVICE_SOURCE_URL github.com/asiainfoLDP/datahub_custom

WORKDIR $GOPATH/src/$SERVICE_SOURCE_URL

ADD . .

RUN go build

CMD ["sh", "-c", "./datahub_custom"]