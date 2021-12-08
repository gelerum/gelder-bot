FROM golang:1.17.4-bullseye

WORKDIR /gelder-bot

COPY ./ ./

RUN go mod download

RUN go build ./cmd/bot/

CMD [ "/gelder-bot/bot" ]