# urgo

bulk url metadata extractor

**INSTALLATION**
> go install github.com/aristosMiliaressis/urgo@latest

**USAGE**
```
USAGE: cat urls.txt | urgo [OPTIONS]

REQUEST OPTIONS:
        -m, --method
                specifies http method to use
        -h, --req-header
                adds request header (can be used multiple times)

EXTRACTION OPTIONS:
        -sC, --status-code
                textracts the response status code
        -rT, --response-time
                extracts the response time
        -rH, --resp-header
                extracts the specified response header (can be used multiple times)
        -f, --favicon
                extracts favicon hash

OUTPUT OPTIONS:
        -oJ, --json
                outputs in JSONL(ine) format
        -h, --help
                prints this help page

PERFORMANCE OPTIONS:
        -t, --threads
                specifies number of threads to use [default: 20]
```