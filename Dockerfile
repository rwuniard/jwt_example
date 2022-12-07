## Build
FROM golang:1.19-buster AS Build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go get github.com/gin-gonic/gin
RUN go get github.com/joho/godotenv
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres
RUN go get -u github.com/golang-jwt/jwt/v4

## Copying all the code and env files
COPY *.go ./
COPY ./config ./config
COPY ./controllers ./controllers
COPY ./middleware ./middleware
COPY ./models ./models
COPY .env ./

## Compile the code
RUN go build -o ./jwt_example

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/jwt_example ./jwt_example
COPY --from=build /app/.env ./

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT [ "./jwt_example" ]