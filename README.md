# Monit

master:  [![CircleCI](https://circleci.com/gh/ernestio/monit/tree/master.svg?style=shield)](https://circleci.com/gh/ernestio/monit/tree/master)  
develop: [![CircleCI](https://circleci.com/gh/ernestio/monit/tree/develop.svg?style=shield)](https://circleci.com/gh/ernestio/monit/tree/develop)

## Synopsis

This microservice listens to all `monitor.user` events that are fired off by the FSM. It allows users to listen to the event stream over SSE (Server Side Events)

These events are collected into an inbox based on the monitor id passed by the user on an action. The inbox is opened on `service.create`, `service.delete` and closed on `service.create.done` and `service.delete.done`. If no inbox with the given ID exists when a user connects, one will be created.

## Installation

```
make deps
make install
```

## Tests

Running the tests:
```
make test
```

### Authentication

Authentication is handled the same way as every other REST API under flow. The auth token retuned from authenticating against `/sessions` should be sent as the HTTP header `X-Auth-Token`.

### Notifications

An example of the `monitor.user` message is as followed. Each message will be sent as newline delimited `data:` over the SSE protocol. The last part of the service field should include the `monitor_id` included in the payload sent when creating or updating a service.

### Usage

To connect to a stream, make the following request with the correct stream (monitor_id) and auth token.

```
$ curl -i -k -H 'X-Auth-Token: ***' https://ernest.local/events?stream=***
```


## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 ernest.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
