#golang as the enviornment
FROM golang:1.8

WORKDIR /go/src/app
COPY . .

#may require "sudo" for the apt-get portion
RUN \
go-wrapper download && \
go-wrapper install && \
apt-get update && \
apt-get install -y mongodb-org=3.2 mongodb-org-server=3.2 mongodb-or-shell=3.2 mongodb-org-mongos=3.2 mongodb-or-tools=3.2 && \
go get gopkg.in/mgo.v2

#defining mountable directories
VOLUME ["/app/data/db"]

#creation of user database through json file
COPY userInfo.json /userInfo.json
CMD mongoimport --host mongodb --db userDb --collection Contacts --type json --file /userInfo.json --jsonArray

CMD ["go-wrapper", "run"]

#default mongo port
# 27017 for process
# 28017 for http
EXPOSE 27017
EXPOSE 28017

#for go use
EXPOSE 8080
