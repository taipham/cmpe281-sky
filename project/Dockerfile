#golang as the enviornment
FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN \
go-wrapper download && \
go-wrapper install 
apt-get update && \
apt-get install -y mongodb-org=3.2 mongodb-org-server=3.2 mongodb-or-shell=3.2 mongodb-org-mongos=3.2 mongodb-or-tools=3.2 && \
go get gopkg.in/mgo.v2

#defining mountable directories
VOLUME ["/app/data/db"]

CMD ["go-wrapper", "run"]

#default mongo port
# 27017 for process
# 28017 for http
EXPOSE 27017
EXPOSE 28017
