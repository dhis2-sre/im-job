FROM golang:1.17-alpine AS build
RUN apk -U upgrade && \
    apk add gcc musl-dev
WORKDIR /src
RUN go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/im-job ./cmd/serve

FROM alpine:3.13
RUN apk --no-cache -U upgrade \
    && apk add --no-cache postgresql-client
WORKDIR /app
COPY --from=build /app/im-job .
COPY --from=build /src/swagger/swagger.yaml ./swagger/
USER guest
CMD ["/app/im-job"]
