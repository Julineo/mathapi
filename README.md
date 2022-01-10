# Math API

Webservice for basic math functions.

Well tested Golang statistics library used to calculate percentile github.com/montanaflynn/stats.
Use `go get github.com/montanaflynn/stats` to install package.

Min/max functions use quantifier to return top/bottom values.


## Installation and usage

Use `go build` to build executable inside mathapi directory.

Use `./mathapi` or `go run main.go` to start the server.

Use `go test` to run tests.

Use `go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 math_test.go` to run benchmarks.

Use `go tool pprof cpu.out` and `go tool pprof cpu.out` with command `web` to visualise runtime and memory usage.



See API usage bellow.


## /min

curl -i -X POST http://localhost:3000/min -H 'Content-Type: application/json' -d '{"list":[10,2,4,8,6,4,10], "Quantifier": 3}'


## /max

curl -i -X POST http://localhost:3000/max -H 'Content-Type: application/json' -d '{"list":[10,2,4,8,6,4,10], "Quantifier": 3}'


## /avg

curl -i -X POST http://localhost:3000/avg -H 'Content-Type: application/json' -d '{"list":[1,2,4,8,6,4,10]}'


## /median

curl -i -X POST http://localhost:3000/median -H 'Content-Type: application/json' -d '{"list":[1,2,4,8,6,4,10]}'


## /percentile

curl -i -X POST http://localhost:3000/percentile -H 'Content-Type: application/json' -d '{"list":[15,20,35,40,50], "Quantifier": 30}'



