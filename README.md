# Fire Dragon

[![Build Status](https://travis-ci.org/bubblebox/server.svg?branch=master)](https://travis-ci.org/bubblebox/server)

A modern content hosting service builtd on Go and Ember.js.

See [the wiki](https://github.com/bubblebox/server/wiki) for more details.

## Running

_Fire Dragon is a work in progress and will probably break a lot at this time._

If all you want is a working application, simple run:

 * make setup && make
 * cd dist
 * ./firedragon and hop over to [http://admin.127.0.0.1.xip.io:8042](http://admin.127.0.0.1.xip.io:8042)

This handy makefile builds both client and server and puts them in the `dist/`
directory.

## Running tests

There's a combined tests target: `make test` that will run tests for
both the server and client. Tests can be run separately for each component,
again by running `make test` in either the `client` or `server` directory.
