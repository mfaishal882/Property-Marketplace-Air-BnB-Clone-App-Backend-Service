FROM golang:1.17.7-alpine as build

# buat direktori app
RUN mkdir /app

# set working dir
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o api-airbnb-alta

EXPOSE 8001

CMD [ "./api-airbnb-alta" ]