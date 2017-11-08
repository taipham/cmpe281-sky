#golang as the enviornment
FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN \
go-wrapper download && \
go-wrapper install 

CMD ["go-wrapper", "run"]