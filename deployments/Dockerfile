FROM alpine

WORKDIR /app

COPY configs configs
COPY web web
COPY build/potato .

EXPOSE 8080

ENTRYPOINT ["./potato"]
