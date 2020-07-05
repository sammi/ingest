FROM golang AS build

RUN apt-get update \
	&& apt-get -y install git unzip build-essential autoconf libtool

RUN git clone https://github.com/google/protobuf.git --recursive \
    && cd protobuf \
    && ./autogen.sh\
    && ./configure\
    && make\
    && make install \
    && ldconfig\
    && make clean\
    && cd ..\
    rm -r protobuf

RUN go get google.golang.org/grpc

RUN go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/app
ADD . /go/src/app

ENV GOPATH="$GOPATH:/go/src/app"

RUN protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative src/contract/*.proto

RUN cd src \
    && go get -d -v ./...\
    && go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /

