FROM golang:1.17.7-alpine as build

# buat direktori app
RUN mkdir /app

# set working dir
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o api-alta-dashboard2

EXPOSE 8000

CMD [ "./api-alta-dashboard2" ]