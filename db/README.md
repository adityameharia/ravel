# db

The `db` package contains the implementation of the `RavelDatabase` struct and its functions. 
This is just a simple overlay on top of [BadgerDB](https://github.com/dgraph-io/badger) exposing common
functions like `Init`, `Close`, `Read`, `Write` and `Delete`

- The Init functions initialises a new BadgerDb instance on the path provided
- The Close function closes the BadgerDB instance
- The Read function starts a read only transaction and returns the value for the given key
- The Write function starts a read-write transaction,writes the key and value to badgerDB and then commits the transaction on a successful write.
- The Delete function also starts a read-write transaction, deletes the key value pair with the corresponding key from BadgerDB and commits the transaction on successful execution