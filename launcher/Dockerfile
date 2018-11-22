FROM library/centos:6
RUN yum -y install wget git
RUN rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
RUN yum -y install epel-release
RUN yum -y install golang

WORKDIR /root/go/app

RUN go get -v github.com/aws/aws-lambda-go/lambda
RUN go get -v github.com/aws/aws-lambda-go/events
RUN go get -v github.com/cenkalti/backoff

COPY ./launcher.go /root/go/app/launcher.go

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o launcher launcher.go
RUN strip launcher