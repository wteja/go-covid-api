# Covid-19 Summary API

This is the Covid-19 Summary API which will summarize raw Covid-19 report into 2 groups
- AgeGroup this is the number of cases by age range
- Province this the number of the cases in each province

## How to run

Run the command below to start API server

```bash
./go-covid-api
```

Or run it manually using `go` command

```bash
go run main.go
```

Then browse to this url: http://localhost:8080/covid/summary

## How to run test

Run this command below to run all test cases:

```bash
go test main_test.go
```

## Author

[Weerayut Teja](https://github.com/wteja)