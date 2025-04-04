# syntax=docker/dockerfile:1

FROM alpine:3.21
# FROM golang:alpine3.20
# AS builder

RUN apk update && apk upgrade
RUN apk add --no-cache nodejs=22.13.1-r0
RUN apk add --no-cache npm=10.9.1-r0
RUN apk add --no-cache gcc=14.2.0-r4
RUN apk add --no-cache go=1.23.7-r0

COPY ./server /tmp/server
COPY ./web /tmp/web

WORKDIR /tmp/web
# creates /tmp/web/dist
RUN rm -f .env
RUN npm install && npm run build

WORKDIR /tmp/server
ENV CGO_ENABLED=1
RUN go mod download
RUN go build -o /tmp/server/rent_manager

# FROM alpine:edge
# FROM busybox

WORKDIR /app
# COPY --from=builder /tmp/server/rent_manager .
# COPY --from=builder /tmp/web/dist ./web
RUN cp /tmp/server/rent_manager .
RUN cp -r /tmp/web/dist/ ./web
RUN mkdir /storage

ENV PORT=3000

CMD ["./rent_manager", "-debug"]
