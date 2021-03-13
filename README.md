# samsocks

This is a tool which will automatically sets up a SOCKS5 proxy to the
I2P network, supporting both TCP and UDP communication, automatically
using the SAM API. It will eventually be used like `torsocks` via the
`LD_PRELOAD` hack, acting as a transparently isolating socksifier.