# GoIPLookUP
A tool to lookup IP address from URL

### Installation

Using go get, download the package with the command below

```
go get github.com/heavykenny/goiplookup@latest
```

### Usage of GoIPLookUp:
```
-v, --verbose   [ verbose output { default: false } ]
-h, --help      [ prints help information ]
-u, --url       [ website url ]
-p, --path      [ file path for bulk urls ]
-o, --output    [ output filename ]
```

### Example

```
goiplookup -u=google.com -v
```

For bulk lists of URLs in a file
```
goiplookup -p=iplist.txt -o=file.txt
```

Connect on [Twitter](http://twitter.com/heavykenny)

NOTE: This is my first trail on building CLI tool (to be updated in the future)
