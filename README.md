# Nezar

## NATS message publisher for Morpheus processes

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/spoonboy-io/nezar?style=flat-square)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/spoonboy-io/nezar?style=flat-square)](https://goreportcard.com/report/github.com/spoonboy-io/nezar)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/spoonboy-io/nezar/build.yml?branch=main&style=flat-square)](https://github.com/spoonboy-io/nezar/actions/workflows/build.yml)
[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/spoonboy-io/nezar/unit_test.yml?branch=main&label=tests&style=flat-square)](https://github.com/spoonboy-io/nezar/actions/workflows/unit_test.yml)
[![GitHub Release Date](https://img.shields.io/github/release-date/spoonboy-io/nezar?style=flat-square)](https://github.com/spoonboy-io/nezar/releases)
[![GitHub commits since latest release (by date)](https://img.shields.io/github/commits-since/spoonboy-io/nezar/latest?style=flat-square)](https://github.com/spoonboy-io/nezar/commits)
[![GitHub](https://img.shields.io/github/license/spoonboy-io/nezar?label=license&style=flat-square)](LICENSE)

## About

Nezar watches [Morpheus CMP](https://morpheusdata.com) processes/events.

## Releases

You can find the [latest software here](https://github.com/spoonboy-io/nezar/releases/latest).

### Get Started

Nezar polls the Morpheus database so needs credentials. The `morpheus` user can be used, but it is preferable to 
create an additional user with SELECT privileges on the `process` and `process_type` tables.

A `mysql.env` file should be created in the same directory as the application from which the database user configuration
will be read. The following example shows the environment variables used by Dozer which should be included in `mysql.env`:

```bash
## MySQL Config
MYSQL_SERVER=127.0.0.1
MYSQL_USER=dozer
MYSQL_PASSWORD=xxxxa8aca0de5dab5fa1bxxxxx

## Optional to override defaults
MYSQL_DATABASE=morpheus
POLL_INTERVAL_SECONDS=3
```


### Installation
Grab the tar.gz or zip archive for your OS from the [releases page](https://github.com/spoonboy-io/nezar/releases/latest).

Unpack it to the target host, and then start the server.

```
./nezar
```

Or with nohup..

```
nohup ./nezar &
```

### License
Licensed under [Mozilla Public License 2.0](LICENSE)
