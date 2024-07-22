## Builder
FROM golang:1.20-alpine as build_base
RUN apk update && apk upgrade && \
    apk --no-cache --update add bash git make gcc g++ libc-dev

WORKDIR /go/src/github.com/rohanchauhan02/food-service
ENV GO111MODULE=on


COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

WORKDIR /go/src/github.com/rohanchauhan02/food-service
COPY . .
RUN ls -lh
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN make engine

## #In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /food-service && mkdir /food-service/app
WORKDIR /food-service/app


EXPOSE 11001

COPY --from=server_builder /go/src/github.com/rohanchauhan02/food-service/engine .
COPY --from=server_builder /go/src/github.com/rohanchauhan02/food-service/app.config* ./

RUN ls -lh
CMD ["/food-service/app/engine"]
