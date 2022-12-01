# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY . .
RUN go mod download

RUN mkdir -p logs
RUN mkdir -p temp

RUN go build -o ./app

EXPOSE 8060
ENV PORT=8060

CMD [ "./app" ]
