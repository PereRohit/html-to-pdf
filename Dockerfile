FROM golang:1.18

RUN apt-get update -y

WORKDIR /app

RUN echo "Downloading wkhtmltopdf package" && \
    wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6.1-2/wkhtmltox_0.12.6.1-2.bullseye_arm64.deb -O wkhtmltopdf.deb

RUN apt-get install -f ./wkhtmltopdf.deb -y

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-wkhtmltopdf main.go

RUN cp go-wkhtmltopdf /bin/.

RUN rm -rf *

#ENTRYPOINT ["wkhtmltopdf"]
ENTRYPOINT ["go-wkhtmltopdf"]



