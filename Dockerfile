FROM surnet/alpine-wkhtmltopdf:3.16.0-0.12.6-small as builder

FROM golang:1.18.7-alpine

WORKDIR /

COPY . .

RUN go mod download
RUN go mod verify
RUN apk add --no-cache \
    libstdc++ \
    libx11 \
    libxrender \
    libxext \
    libssl1.1 \
    ca-certificates \
    fontconfig \
    freetype \
    ttf-dejavu \
    ttf-droid \
    ttf-freefont \
    ttf-liberation \
    && apk add --no-cache --virtual .build-deps \
    \
    # Clean up when done
    && rm -rf /tmp/* \
    && apk del .build-deps

RUN go build -o hello main.go

COPY --from=builder /bin/wkhtmltopdf /bin/wkhtmltopdf


EXPOSE 3000

CMD [ "./hello" ]
