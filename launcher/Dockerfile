FROM library/centos:6
RUN yum -y install wget git
RUN rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
RUN yum -y install epel-release
RUN yum -y install golang
RUN yum -y install xz

WORKDIR /root/go/app

RUN wget https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz -O upx.tar.xz
RUN mkdir upx && tar xvf upx.tar.xz -C upx --strip-components=1

RUN go get -v github.com/aws/aws-lambda-go/lambda
RUN go get -v github.com/aws/aws-lambda-go/events
RUN go get -v github.com/cenkalti/backoff

COPY ./launcher.go /root/go/app/launcher.go
# COPY ./test.go /root/go/app/test.go

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o launcher launcher.go
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o test test.go

RUN strip launcher
RUN ./upx/upx --ultra-brute --best launcher