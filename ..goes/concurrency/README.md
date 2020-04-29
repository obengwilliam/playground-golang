### Implment Concurrency in Go in Python
http://www.doxsey.net/blog/go-concurrency-from-the-ground-up


Goroutine is lightweight thread managed by the go runtime.

A goroutine run "concurrently" with the main function.
When the main function ends go routines are teminated.


Channels are synchronous , meaning they block till both ends send and receive


buffered channels are asynchronous , they simply channels with capacity.
