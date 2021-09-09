# Building the binary of the App
FROM golang:1.15 AS build

# `boilerplate` should be replaced with your project name
WORKDIR /go/src/boilerplate

RUN git config --global url."https://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:x-oauth-basic@github.com/ProjectAthenaa".insteadOf "https://github.com/ProjectAthenaa"

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

# Create the `public` dir and copy all the assets into it
RUN mkdir ./static
COPY ./static ./static

# `boilerplate` should be replaced here as well
COPY --from=build /go/src/boilerplate/app .

ENV REDIS_URL=rediss://default:kulqkv6en3um9u09@athena-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061
ENV PG_URL=postgresql://doadmin:rh3rc0vgg1f706kz@athenadb-do-user-9223163-0.b.db.ondigitalocean.com:25060/defaultdb


# Exposes port 3000 because our program listens on that port
EXPOSE 8080 8080

CMD ["./app"]