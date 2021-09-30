FROM golang:latest as builder

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

COPY ./src .

RUN go mod download

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main ./app/main.go  

FROM xushikuan/alpine-build:1.0

ENV BUILDER_WORK_DIR=/app
ENV WORK_DIR=/app
ENV TIME_ZONE=Asia/Thailand
WORKDIR $WORK_DIR
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY --from=builder $BUILDER_WORK_DIR/main $WORK_DIR

#Copy sqlite database in to container
COPY ./src/transaction.db transaction.db

EXPOSE 25671

ENTRYPOINT ./main -c config.conf