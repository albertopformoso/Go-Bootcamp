FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY . .

# RUN go mod download
RUN go mod verify    && \
    go mod tidy -v

RUN go build -o bin/app ./cmd/...

#### Alpine Image ####
FROM alpine:3.16

ARG USER="resume-be"

RUN adduser \
        -g "log daemon user" \
        --disabled-password \
        ${USER}

RUN mkdir -p bin/

COPY --from=builder /app/bin/app /bin/app

EXPOSE 8080

ENTRYPOINT ["/bin/app"]
