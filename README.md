# go-qrcode #

## Install

    go get -u github.com/skip2/go-qrcode/...

A command-line tool `qrcode` will be built into `$GOPATH/bin`.

## Usage

    import qrcode "github.com/skip2/go-qrcode"

- **Create a 256x256 PNG image:**

    err := qrcode.WriteFile(text, qrcode.Medium, 256, "image.png")

## Maximum capacity
The maximum capacity of a QR Code varies according to the content encoded and the error recovery level. The maximum capacity is 2,953 bytes, 4,296 alphanumeric characters, 7,089 numeric digits, or a combination of these.

## Run 
make swag-gen \n
go run program_file_path/main.go

## Check
http://localhost:8090/swagger/index.html#




