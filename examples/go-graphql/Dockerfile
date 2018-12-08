FROM library/centos:6
RUN yum -y install wget git
RUN rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
RUN yum -y install epel-release
RUN yum -y install golang

WORKDIR /root/go/app

RUN go get -v github.com/graph-gophers/graphql-go
RUN go get -v github.com/graph-gophers/graphql-go/relay

COPY ./main.go /root/go/app/main.go

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o main.exe main.go