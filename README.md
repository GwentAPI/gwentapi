# GwentAPI
The code for GwentAPI

GwentAPI is a RESTful API dedicated to provide information regarding GWENT®: The Witcher Card Game. You may use its interface to consume information on cards, factions, etc. You can find the API documentation at [https://gwentapi.com/swagger/](https://gwentapi.com/swagger/).


# How to use

GwentAPI is release for both Linux and Windows. You can find the download [here](https://github.com/GwentAPI/gwentapi/releases).
The data allowing you to seed the database is included in the releases. Read the file ``DATABASE_INSTRUCTIONS.md`` to know how to feed the database.

Change the values found in ``config.toml`` to suit your environment. The configuration file must be kept in the same directory as the executable. 


# Install from source

## Requirement & dependencies

* [Go 1.8](https://golang.org/dl/)
* MongoDB 3.4
* Git
* [dep](https://github.com/golang/dep)

## Installation

Assuming you installed and configure all the required softwares listed above and that you have a working Go setup:

### Install the go dependencies

``dep ensure``

### Build the software
1. ``go get github.com/GwentAPI/gwentapi``
3. GwentAPI uses the experimental [dep](https://github.com/golang/dep) tool to manage the Go dependencies, however they are included in the ``vendor`` directory to help build the software on CI platforms.
4. Change your monboDB credentials in the config file. By default it connects to localhost on no particular users (no SSL, access control not enabled) and uses the gwentapi database.
5. Compile:
    * ``go build``
6. Optionally set the version of the binary (displayed with the -v flag) by changing the value of ``version`` in ``main.go`` before compiling or set it at compile time with the following command:
    * ``go build -ldflags="-X main.version=<versionString>" ``
7. Optionally modify the ``baseURL`` in the ``config.toml`` file to change where resources are pointing to.

**The program will run on localhost on port 8080 by default.**

## How to contribute?

Read the ``CONTRIBUTING.md`` file.

## Where are the artworks?

At the time of writing, the artworks is taken from http://gwentify.com (a copy is saved, they aren't hotlinked).
The files aren't included in the repo because git doesn't like binary files and neither do I like including binary files in git.
