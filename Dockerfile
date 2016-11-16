FROM        golang:onbuild

COPY        . /app
WORKDIR     /app
RUN         go build -o main .
EXPOSE      8080

ENTRYPOINT  ["/app/main"]
