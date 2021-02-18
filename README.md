# go-tar

This a warper of `archive/tar` & `compress/gzip` standard libary.

It's only create `tar` and `tar.gz` currently.

Compression is depends on filename of the archive.

## How to use this module?

To use this module,

```sh
go get github.com/mrinjamul/go-tar/tarwarper
```

and then,

```go
import (
    ...
    "github.com/mrinjamul/go-tar/tarwarper"
)
```

## Feature

- `tarwarper.CreateTar`
- `tarwarper.ExtractTar`

## CLI Usage

    Usage: go-tar [options]
    Options:
    -c, --create string    Create Tar
    -x, --extract string   Extract Tar
    -h, --help             help message
    -o, --output string    Output name

## License

This go module is licensed under MIT Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>
