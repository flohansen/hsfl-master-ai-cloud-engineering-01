FROM golang:1.21-alpine
RUN apk add --update nodejs npm
RUN npm i -g yarn

ENV PRODUCTS_ENDPOINT=

WORKDIR /app
COPY ./src/web-service ./src/web-service

WORKDIR /app/src/web-service
RUN yarn
RUN yarn build
RUN go build -o ./main

EXPOSE 3000
CMD ["/app/src/web-service/main"]
