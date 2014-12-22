# Proposal Converter

A simple tool to convert the CfPs from .csv format to respective formats for the webpage.

## Build

[Go](http://golang.org) needs to be installed. Then simply:

```
go build
```

## Usage for Schedule

```
propConv -i <CSV-FILE> -d 1 -t 1
```

This will convert the Proposals for day 1 and track 1 and print on Standard Out. (redirect as you like)


## Usage for generating subpages

```
propConv -i <CSV-FILE> -detail -outdir <OUTPUT-DIR>
```

This will generate all single pages for all talks and save them in OUTPUT-DIR.

