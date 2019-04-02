FROM golang

COPY ./ /go/src/app

WORKDIR /go/src/app

RUN go build -o app

EXPOSE 80

ENTRYPOINT [ "./app" ]
