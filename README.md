# Curso Go Expert

## Coverage

exec: `go test -coverprofile=coverage.out`

### go tool

exec: `go tool cover -html=coverage.out`

### exec Benchmark

exec: `go test -bench=.` com os Tests
exec: `go test -bench=. -run=^#` apenas o Benchmark
exec: `go test -bench=. -run=^# -count=10` com 10 chamadas
exec: `go test -bench=. -run=^# -count=10 -benchtime=3s` executa por três 3s

### Fuzzing

Teste de mutação
exec: `go test -fuzz=. -run=ˆ#`

### Template Go Projects

`https://github.com/golang-standards/project-layout`
