# Nikel API

[![Build Status](https://travis-ci.com/nikel-api/nikel.svg?branch=master)](https://travis-ci.com/nikel-api/nikel)

Nikel (pronunciation: `/'ni:kɛl/`) is a collection of data APIs on information about the University of Toronto.

### Documentation

[Nikel API Documentation](https://docs.nikel.ml/docs)

### API Wrappers

##### Official

* [nikel-ts (Node.js)](https://www.npmjs.com/package/nikel)

### Endpoints currently supported

* /courses
* /textbooks
* /exams
* /evals
* /food
* /services
* /buildings
* /parking

### Self Hosting

Please make sure you have the same go version displayed in the `go.mod` file. It should usually be the latest stable release. If you are unsure which go version you have, use `go version` to find out.

Nikel should work on any 32/64 bit system with go installed.

1. git clone
```
git clone https://github.com/nikel-api/nikel.git
```
2. cd into nikel-parser submodule
```
cd nikel/nikel-parser
```
3. Update nikel-parser submodule to latest
```
git pull origin master
```
4. cd into nikel-core
```
cd ../nikel-core
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

7. Optional configuration

* By default, nikel-core should be listening and serving on port 8080. To change the port, modify the `PORT` environment variable.
* To suppress debug logs, add the environment variable `GIN_MODE` with the value `release`.
* To add optional application metrics via New Relic APM, add the environment variable `NEW_RELIC_LICENSE_KEY` along with a license key.

### Contributing

For contributing, there are a few things to look out for:

* Always use `go fmt` to format code
* Consult the article [Godoc: documenting Go code](https://blog.golang.org/godoc) on how to write docstrings if you aren't 100% sure

If you find any inconsistencies or parts of code that can be reworked, any pull requests are greatly appreciated.
