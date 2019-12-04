## tinybit

Tiny Bitcoin Node. In early development, probably will never be finished. Hopefully, will have some useful features.

Implementation is explained in a series of blog posts:

1. [Programming Bitcoin Network](https://jeiwan.net/posts/programming-bitcoin-network/)
1. [Programming Bitcoin Network, part 2](https://jeiwan.net/posts/programming-bitcoin-network-2/)
1. [Programming Bitcoin Network, part 3](https://jeiwan.net/posts/programming-bitcoin-network-3/)

### Running
1. Install [btcd](https://github.com/btcsuite/btcd).
1. `btcd --configfile ./btcd.conf`
1. `go build`
1. `./tinybit`


### Implemented so far

1. messages serialization and deserialization (see `binary` package)
1. 'version', 'verack' messages
1. 'inv'
1. 'tx'
1. 'getdata'
1. 'ping', 'pong'