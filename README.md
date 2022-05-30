# go-middleware
I am creating middleware for web server development.

The currently implemented functions are as follows.
- Access-Log
- Panic Recovery
- Basic-Authentication

## Usage
```
import mid "github.com/pyotarou/go-middleware"

func hogehoge() {
    http.Handle("/", mid.AccessLogger(<http.Handler>))
}
```

## Installation
```
$ go get github.com/pyotarou/go-middleware
```

## Author
Seitaro Fujigaki (a.k.a pyotarou)