# db

The `db` package contains the implementation of the `RavelDatabase` struct and its functions. 
This is just a simple overlay on top of [BadgerDB](https://github.com/dgraph-io/badger) exposing common
functions like `Init`, `Close`, `Read`, `Write` and `Delete`



- `Init` - initialises the `RavelDatabase` struct and opens a connection to BadgerDB.
- `Close` - closes the connection to BadgerDB.
- `Read` - starts a read only transaction and returns the value for the given key.
- `Write` - starts a read-write transaction, writes the key and value to badgerDB and then commits the transaction.
- `Delete` - starts a read-write transaction, deletes the key value pair and commits the transaction.