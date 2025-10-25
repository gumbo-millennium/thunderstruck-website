FROM golang:1.24.6-alpine3.22 AS build
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o /bin/app ./main.go

FROM scratch AS release
COPY --from=build /bin/app /bin/app
CMD [ "/bin/app" ]

FROM golang:1.24.6-alpine3.22 AS debug
WORKDIR /debug
COPY ./ ./
RUN apk add git
RUN go mod download
RUN go install github.com/air-verse/air@v1.61.7
CMD [ "air", "-c", ".air.toml" ]
