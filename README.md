# ms
[![Build Status](https://travis-ci.org/dolegi/ms.svg?branch=master)](https://travis-ci.org/dolegi/ms)

Golang has support for this built in. Check out [ParseDuration](https://golang.org/pkg/time/#ParseDuration)

Converts time strings to and from number of milliseconds

This is a rewrite of [zeit/ms](https://github.com/zeit/ms) in golang

Documentation [godoc](https://godoc.org/github.com/dolegi/ms)

## Examples

```go
ms.Parse('2 days')  // 172800000
ms.Parse('1d')      // 86400000
ms.Parse('10h')     // 36000000
ms.Parse('2.5 hrs') // 9000000
ms.Parse('2h')      // 7200000
ms.Parse('1m')      // 60000
ms.Parse('5s')      // 5000
ms.Parse('1y')      // 31557600000
ms.Parse('100')     // 100
ms.Parse('-3 days') // -259200000
ms.Parse('-1h')     // -3600000
ms.Parse('-200')    // -200
```

### Convert from Milliseconds

```go
ms.Fmt(60000)             // "1m"
ms.Fmt(2 * 60000)         // "2m"
ms.Fmt(-3 * 60000)        // "-3m"
ms.Fmt(ms('10 hours'))    // "10h"
```

### Time Format Written-Out

```go
ms.FmtLong(60000)                 // "1 minute"
ms.FmtLong(2 * 60000)             // "2 minutes"
ms.FmtLong(-3 * 60000)            // "-3 minutes"
ms.FmtLong(ms.Parse('10 hours'))  // "10 hours"
```

## Tests
Tests can be run with `go test`
