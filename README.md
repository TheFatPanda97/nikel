<h1 align="center">
  <br>
  <a href="https://docs.nikel.ml"><img src="https://docs.nikel.ml/img/nikel-api-circle.png" alt="Nikel API" width="200"></a>
  <br>
  Nikel API
  <br>
</h1>

<h4 align="center">A collection of data APIs for the University of Toronto.</h4>

<p align="center">
  <a href="https://travis-ci.com/nikel-api/nikel">
    <img alt="Build Status" src="https://img.shields.io/travis/nikel-api/nikel">
  </a>
  <a href="https://hub.docker.com/r/nikelapi/nikel">
    <img alt="Docker Hub Build Status" src="https://img.shields.io/docker/cloud/build/nikelapi/nikel">
  </a>
  <a href="https://status.nikel.ml/">
    <img alt="API Status" src="https://img.shields.io/uptimerobot/status/m785541667-14c2f35b7d11487c0874bdd7">
  </a>
  <a href="https://goreportcard.com/report/github.com/nikel-api/nikel">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/nikel-api/nikel">
  </a>
  <a href="https://github.com/nikel-api/nikel/blob/master/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/nikel-api/nikel">
  </a>
</p>

<p align="center">
  <a href="#documentation">Documentation</a> •
  <a href="#api-wrappers">API Wrappers</a> •
  <a href="#self-hosting">Self Hosting</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

## Documentation

[Nikel API Documentation](https://docs.nikel.ml)

## API Wrappers

#### Official

* [nikel-ts (Node.js)](https://www.npmjs.com/package/nikel)

#### Unofficial

* [nikel-rs (Rust)](https://crates.io/crates/nikel-rs)

Please feel free to submit a pull request to add your own API wrapper to this list!

## Self Hosting

Please consult the [configuration](#configuration) section on what environment variables to set.

### Using Docker

#### Deployment via Docker Hub images (recommended)

You can pull Nikel API's prebuilt Docker images from Docker Hub.

1. Pull the latest image from Docker Hub:
```
docker pull nikelapi/nikel
```

2. Run image (you can tweak variables accordingly)
```
docker run --publish 8080:8080 --detach --name nikel-core nikelapi/nikel:latest
```

#### Deployment via local Docker compose

Make sure that your docker version supports the docker-compose version displayed in the `docker-compose.yaml` file.

1. Run `docker-compose`
```
docker-compose up -d
```

### Traditional Deployment

Please make sure you have the same go version displayed in the `go.mod` file. It should usually be the latest stable release. If you are unsure which go version you have, use `go version` to find out.

Nikel should work on any 32/64 bit system with go installed.

1. git clone
```
git clone https://github.com/nikel-api/nikel.git
```
2. cd into nikel
```
cd nikel
```
3. Update submodules
```
git submodule update --init
```
4. cd into nikel-core
```
cd nikel-core
```
5. Build nikel-core
```
go build
```
6. Run nikel-core
```
Windows
./nikel-core.exe

Linux and macOS
./nikel-core
```

## Configuration

* By default, nikel-core should be listening and serving on port 8080. To change the port, modify the `PORT` environment variable.
* To suppress debug logs, add the environment variable `GIN_MODE` with the value `release`.
* To add optional ratelimiting, add the environment variable `RATELIMIT` with a positive integer value representing the number of reqs/s.

## Contributing

For contributing, there are a few things to look out for:

* Always use `go fmt` to format code
* Consult the article [Godoc: documenting Go code](https://blog.golang.org/godoc) on how to write docstrings if you aren't 100% sure

If you find any inconsistencies or parts of code that can be reworked, any pull requests are greatly appreciated.

## License

[MIT](https://github.com/nikel-api/nikel/blob/master/LICENSE)