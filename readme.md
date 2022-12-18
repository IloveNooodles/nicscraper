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

## NEW MODE!!!: Teams Scrap

Use `-u` flag to use it. And if you're using it please provide `Jwt` and `Cvid` token

JWT And Cvid Token can be obtain from your network tab steps are below

1. Open `Micorosft Teams` in browser
2. Open developer console
3. Type something in the search bar
4. find the `suggestions?scenario=powerbar`
5. Right click it and `copy as fetch`
6. Then provide jwt-auth and the Cvid from there

In \*nix systems:

```bash
export JWT_TOKEN=<your-token-here>
export CVID_TOKEN=<your-token-here>
./nicscraper -p 135 -y 18 -u
```

In Windows systems:

```cmd
SET JWT_TOKEN=<your-token-here>
SET CVID_TOKEN=<your-token-here>
nicscraper -p 135 -y 18 -u
```

You can also provide your token via the `-j/--jwt` and `-c/--cvid`.

```bash
./nicscraper -p 135 -y 18 --jwt <your-token-here> --cvid <your-token-here>
```

### Using Docker

To run using Docker:

```bash
docker build -t nicscraper:latest .
docker run nicscraper -p 135 -y 18 --token <your-token-here>
```

## Contributors

<a href="https://github.com/mkamadeus/nicscraper/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=mkamadeus/nicscraper" />
</a>
