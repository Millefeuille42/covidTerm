# CovidTerm

A little Terminal Program that gets global or per country All-Time Covid19 stats.<br>
This program uses [covid19api.com](https://covid19api.com/).

## Examples
```bash
$ covidTerm

Global reports:
        Cases:          4977471
        Deaths:         329513
        Recoveries:     1838344
```
```bash
$ covidTerm france spain germany

Reports for France:
        Cases:        180933
        Deaths:       28025
        Recoveries:   62678

Reports for Spain:
        Cases:        232037
        Deaths:       27778
        Recoveries:   0

Reports for Germany:
        Cases:        177778
        Deaths:       8081
        Recoveries:   155681
```
## Syntax:
```bash
$ covidTerm
```
Returns global stats
```bash
$ covidTerm france spain germany
```
Returns per country stats for given countries

## Installation

This program requires that you have [Go installed](https://golang.org/dl/) 
```bash
$ go build -o ~/covidTerm github.com/Millefeuille42/covidTerm
```
```bash
$ git clone https://github.com/Millefeuille42/covidTerm
$ go build -o ~/covidTerm covidTerm/
```
Then move the binary to any directory in your `$PATH`