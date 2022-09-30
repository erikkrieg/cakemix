run *ARGS:
  @go run ./main.go {{ ARGS }}

build:
  @go build

clean:
  @rm -rf tmp
  @mkdir tmp
