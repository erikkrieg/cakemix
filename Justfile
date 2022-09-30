run *ARGS:
  @go run ./main.go {{ ARGS }}

build:
  @go build

clean:
  @rm -rf tmp
  @mkdir tmp

alias e := examples
alias example := examples
examples: clean
  @just run -o tmp examples
  
