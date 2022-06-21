# http-sleep

HTTP endpoint that just sleeps for specified duration and returns 200.


## Usage

```
$ ./http-sleep -port 8000 &
$ time curl localhost:80/3s
OK
real	0m3.011s
```

Duration can be any [time.ParseDuration()](https://pkg.go.dev/time#ParseDuration) recognizes.
