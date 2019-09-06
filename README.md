# go-enc
encode/decode cli tool written in Go.

# Installation
- base64

```sh
$ go get github.com/skanehira/go-enc/cmd/b64
```

- url
```sh
$ go get github.com/skanehira/go-enc/cmd/urlenc
```

# Usage
- base64
```sh
$ b64 -h
Usage of b64:
  -d string
        decode to string
  -e string
        encode to base64
  -f string
        encode file to base64
```

- urlenc
```sh
$ urlenc -h
Usage of urlenc:
  -d string
        decode URL
  -e string
        encode URL
```
