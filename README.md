Mesures for golang
==================

Simple _go_ proxy for tcp -> eventsource,
for the [mesures](https://github.com/bearstech/node-mesures) project.

Go is lighter than nodejs? Yes. RSS is 20% of nodejs version.

Memory is cheap, but you can be greedy.

Test
----

	go test

Try it
------

    ./build.sh
    ./mesures

In an other terminal:

    telnet localhost 5000

You can test command, and the server will answer.

    popo 42[ENTER]

In another terminal

    curl -v http://localhost:8000/events

You will see an eventsource (a never ending webpage),
with the current values as first event.
Try to type more events in telnet.

Licence
-------

GPLv3
