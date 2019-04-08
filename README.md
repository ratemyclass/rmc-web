# rmc-web
[![Build Status](http://img.shields.io/travis/badges/rmc-web.svg?style=flat-square)](https://travis-ci.org/badges/rmc-web)
[![Coverage Status](http://img.shields.io/coveralls/badges/rmc-web.svg?style=flat-square)](https://coveralls.io/r/badges/rmc-web)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://rmc-web.mit-license.org)

The web backend implementation for RateMyClass.

## Installation

In order to use and work on this project, you must have [Docker](https://docs.docker.com/install/) installed. Now all you need to do is clone the repository to your computer! There are two modes you may run the software in: Production and Development.

### Production

To build the production image, navigate to the root directory and run `make install`. To run the container, run `make prod` (or simply `docker-compose up prod` if you don't like Make).

**Note:** Every time you make changes to the source, you will need to rebuild the production image in order for the changes to take effect. If you intend on making many changes, it is advised that you run in development mode instead to save time.

### Development

To build the development image, navigate to the root directory and run `make install`. To run the container, run `make dev` (or simply `docker-compose up dev` if you don't like Make).

If you plan on making changes to files on the backend, be sure to restart the container to see the modifications.
