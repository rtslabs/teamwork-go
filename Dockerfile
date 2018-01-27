# build stage
FROM golang:alpine AS build-env
ADD . /src
ENV GOPATH /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp