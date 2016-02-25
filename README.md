# gin-react-boilerplate

- [About](#about)
- [Requirements](#requirements)
- [Setup](#setup)
- [License](#license)

## About

This project is a boilerplate for [Facebook React](https://facebook.github.io/react/) with Golang server [Gin](https://gin-gonic.github.io/gin/).

## Requirements

* [Golang](http://golang.org/)
* [Node.js](https://nodejs.org/)

## Setup

1. Install some golang tools by `go get`

 ```bash
go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/elazarl/go-bindata-assetfs/...
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-gonic/contrib/static
go get -u github.com/pilu/fresh
go get github.com/mitchellh/gox
 ```
2. Install JavaScript dependencies

 ```bash
npm install
 ```

## Run with development mode

1. Generate bindata.go with debug mode.

 ```bash
npm run bindata
 ```

2. Start webpack and gin with watch mode.

 ```bash
npm run dev & fresh
 ```

## Release Build

Run webpack with production mode, go-bindata and go build in turn. All you have to do is run `npm run build`.

```bash
npm run build
```

The artifacts are created under `./dist` directory.

```bash
$ ls -l dist
total 9936
-rw-r--r--    1 root root 10175248 Feb 25 10:10 web_linux_amd64
-rwxr-xr-x    1 root root 10172416 Feb 25 10:10 web_windows_amd64.exe
```

## License

Licensed under the [MIT](/LICENSE.txt) license.
