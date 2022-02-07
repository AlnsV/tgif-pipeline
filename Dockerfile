FROM golang:1.17-alpine as builder

# Installing git
RUN apk add --no-cache curl git

# Add Maintainer Info
LABEL maintainer="Jesus Valdez <alonso6230@gmail.com>"
WORKDIR /go/src/tgif-pipeline

ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o  pipeline tgif-pipeline/cmd/pipeline

# Final stage that will run the executable binary resulting from previous stage
FROM alpine


WORKDIR /pipeline
COPY --from=builder /go/src/tgif-pipeline /pipeline
RUN chmod +x ./pipeline
RUN apk add ca-certificates && rm -rf /var/cache/apk/*
CMD ./pipeline
