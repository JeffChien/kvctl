# kvctl

`kvctl` is a cli tool which provide unix like commands to manipulate key-value storage.

# features
- cat
- mkdir
- touch
- rm
- ls
- tar (not compatible with each backend)
- ... more

# support backend
any backend support by `docker/libkv` should work.
- consul
- etcd
- zookeeper
- BoltDB

currently only test with `consul`, `etcd`

# Notes
- need `go1.5` with `GO15VENDOREXPERIMENT=1`
- `kvctl` is under heavy development, command behavior may be changed.
- use `docker/libkv` to handle different kv backends, but currently use my own fork,
  will migrate to upstream when some issues are solved.

# Roadmap
- `cp`: copy from a to b
- `repl`: interactive mode
- `watch`: watch change under some path and return changed key
- support `ttl`
- document
