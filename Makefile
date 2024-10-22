build:
    go build -o myapp main.go

test:
    go test ./...

load_test:
    k6 run tests/load/load.test.js

run: build
    ./myapp
