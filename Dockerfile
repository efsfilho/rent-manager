# syntax=docker/dockerfile:1

FROM alpine:3.21
# FROM golang:alpine3.20
# AS builder

# ARG PORT=":3000"
# ARG ADDRESS="10.0.0.10"

RUN apk update && apk upgrade
RUN apk add --no-cache nodejs=22.11.0-r1
RUN apk add --no-cache npm=10.9.1-r0
RUN apk add --no-cache gcc=14.2.0-r4
RUN apk add --no-cache go=1.23.4-r0

COPY ./server /tmp/server
COPY ./web /tmp/web

WORKDIR /tmp/web
# creates /tmp/web/dist
RUN rm .env
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


# EXPOSE $SERVER_PORT

# VOLUME ["${EZNVR_STORAGE}"]


ENV PORT=5000

# ENV MONGODB_URI=""

# EXPOSE ${PORT}

# RUN touch .env

# RUN cat > .env <<EOF
# VITE_APP_ADDRESS=$VITE_APP_ADDRESS
# EOF

# ENTRYPOINT /app/rent_manager -debug
# ENTRYPOINT ["./rent_manager", "-debug"]
CMD ["./rent_manager", "-debug"]
#CMD ["tail","-f","/dev/null"]
