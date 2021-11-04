FROM golang:1.16-alpine as build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN  go build -o /cmtapp github.com/bburaksseyhan/ctmapp/src/cmd/api   

FROM alpine:3.14

RUN apk update \
    && apk upgrade\
    && apk add --no-cache tzdata curl

#RUN apk --no-cache add bash
ENV TZ Europe/Istanbul

WORKDIR /app
COPY --from=build-env /cmtapp .

EXPOSE 80
CMD [ "./cmtapp" ]