
# Fxproxy Case

Neda Divbandi fxproxy challenge
i used https://github.com/traefik/whoami to create downstream a tiny Go webserver  

## build and Run Project on local

With Docker:

```sh
$  docker build -t fxproxy .
$  docker run -it --rm -p 8080:8080 -e PORT=':8080' -e SCHEMA='http' -e DOWNSTREAM='localhost
:49153' fxproxy 

```
##### To Run the whole tests

```

go test ./...
```
