# timeserver
A really simple Golang HTTP server for testing purposes. Requests to the root (`/`) reply with a message giving the current time and the host/port it's listening on.

    $ curl http://127.0.0.1:4321
    The time is 2015-10-25 15:50:19.332605809 -0400 EDT from 127.0.0.1:4321

## Usage
Make sure you run a `go get` to retrieve all the dependencies.

    $ go get
    $ go build

When you start the server, you must specify the host and port which the server will listen on with the `HOSTPORT` environment variable.

        $ HOSTPORT=127.0.0.1:4321 ./timeserver

You can also start the server using the `go run` command.

    $ HOSTPORT=127.0.0.1:4321 go run timeserver.go
    2015/10/25 15:55:01 Starting server on 127.0.0.1:4321...
    2015/10/25 15:55:41 Request received from 127.0.0.1:37494
    2015/10/25 15:55:42 Request received from 127.0.0.1:37495

The server logs out all requests it receives.
