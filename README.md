# go-weasel

[![GoDoc](https://godoc.org/github.com/svetlana-rezvaya/go-weasel?status.svg)](https://godoc.org/github.com/svetlana-rezvaya/go-weasel)
[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-weasel)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-weasel)
[![Build Status](https://app.travis-ci.com/svetlana-rezvaya/go-weasel.svg?branch=master)](https://app.travis-ci.com/svetlana-rezvaya/go-weasel)
[![codecov](https://codecov.io/gh/svetlana-rezvaya/go-weasel/branch/master/graph/badge.svg)](https://codecov.io/gh/svetlana-rezvaya/go-weasel)

[Weasel program](https://en.wikipedia.org/wiki/Weasel_program).

## Installation

```
$ go get github.com/svetlana-rezvaya/go-weasel/...
```

## Usage

```
$ go-weasel -h | -help | --help
$ go-weasel [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-sample STRING` &mdash; target string (default: `METHINKS IT IS LIKE A WEASEL`);
- `-rate FLOAT` &mdash; character mutate rate (default: `0.05`);
- `-count INTEGER` &mdash; population size (default: `100`).

## Output Example

```
0 GFILWFXLYHNKDULEAEZWQJ VN II
10 GEIHINEN HNKWSLLIEZ AJ FOCHL
20 GETHINKN INKISLLIKE A WEUCHL
30 GETHINKN ITKISLLIKE A WEASEL
40  ETHINKS IT IS LIKE A WEASEL
45 METHINKS IT IS LIKE A WEASEL
```

## License

The MIT License (MIT)

Copyright &copy; 2021 svetlana-rezvaya
