# nicscraper

![](https://shields.io/badge/go-v1.15-blue?logo=go)
![](https://img.shields.io/github/issues/mkamadeus/nicscraper)
![](https://img.shields.io/github/forks/mkamadeus/nicscraper)
![](https://img.shields.io/github/stars/mkamadeus/nicscraper)

Tiny Go-based binary to scrape from nic.itb.ac.id.

## Prerequisite

For compiling:

- Go v1.15
- Docker (optional)

## How to Use

### Compiling the source code

You can compile the source code by yourself.
If preferable, a precompiled binary will be supplied in the releases page.
To compile by yourself:

```bash
make build
# ...or
go build -o nicscraper main.go
```

### Using precompiled binaries

In \*nix systems:

```bash
export NIC_CI_TOKEN=<your-token-here>
./nicscraper -p 135 -y 18
```

In Windows systems:

```cmd
SET NIC_CI_TOKEN=<your-token-here>
nicscraper -p 135 -y 18
```

You can also provide your token via the `-t`/`--token`.

```bash
./nicscraper -p 135 -y 18 --token <your-token-here>
```

### Using Docker

To run using Docker:

```bash
docker build -t nicscraper:latest .
docker run nicscraper -p 135 -y 18 --token <your-token-here>
```
