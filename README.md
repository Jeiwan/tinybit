## tinybit

Tiny Bitcoin Node. In early development, probably will never be finished. 
Implementation is explained in a series of blog posts:

1. [Programming the Bitcoin Network](https://jeiwan.cc/posts/programming-bitcoin-network/)

### Running
1. Install [btcd](https://github.com/btcsuite/btcd).
1. `btcd --configfile ./btcd.conf`
1. `go build`
1. `./tinybit`