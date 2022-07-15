# `golicenses`

[![Go Reference](https://pkg.go.dev/badge/github.com/imjasonh/golicenses.svg)](https://pkg.go.dev/github.com/imjasonh/golicenses)

This is an **experimental** package to lookup the license for a Go package.

For example:

```golang
lic, _ := golicenses.Get("github.com/google/go-containerregistry")
fmt.Println(lic)
```

prints

```
Apache-2.0
```

This is based on the public BigQuery dataset provided by https://deps.dev/

This repo periodically queries the public dataset and regenerates `licenses.csv`, which is gzipped and `//go:embed`ed into the package.

The result is a ~3MB dependency that can be loaded and queried in ~200ms.
There are almost certainly more optimizations that could improve both size and query time.
