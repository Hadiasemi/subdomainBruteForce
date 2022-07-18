# Subdomain Brute Forcing:
This code using parrallel computing power of golang to make subdomain brute forcing faster.

## Build the command:

```bash
▶ go build brute.go
```

## Help:

```bash
▶ ./brute -h

Usage of ./brute:
  -f string
        subdomain file (default "./deepmagic.com-prefixes-top500.txt")
 -t int
        thread (default max)
  -u string
        specify the url
```

## Run:

```bash
▶ ./brute -f subdomainFile -u url
```

By default using the SEC List top 500 subdomain.
