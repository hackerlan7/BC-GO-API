# FROM golang

# WORKDIR /app 

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY . .
# RUN go build -o /go-api

# CMD [ "/go-api" ]

FROM golang:1.18 as builder

ARG MKEY
ENV MKEY=${MKEY}
ENV GO111MODULE=on
ENV GIT_PATH="https://${MKEY}:x-oauth-basic@github.com/"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN git config --global url.${GIT_PATH}.insteadOf "https://github.com/"
RUN git --version
RUN GOSUMDB=off go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=direct GOSUMDB=off go build cmd/server/main.go

# # final stage
# FROM alpine:3.6 as production
# RUN apk add -U --no-cache ca-certificates
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY --from=builder /app/internal /app/internal
# COPY --from=builder /app/doc /app/doc
# COPY --from=builder /app/assets /app/assets
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["/app/main"]