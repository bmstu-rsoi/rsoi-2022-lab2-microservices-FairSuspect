# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY . .
RUN go mod download

RUN mkdir -p logs
RUN mkdir -p temp

RUN go build -o ./app

EXPOSE 8070
ENV PORT=8070

CMD [ "./app" ]
