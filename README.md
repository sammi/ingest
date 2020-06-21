# ingest

Accept client gRPC request and save data into rocksdb.

## System requirements

* Ubuntu18.04
* gcc 7.5.0+
* cmake 3.17.0+

## Build and install 

```
mkdir build
cd build
cmake ..
make
sudo make install
```

## Start up ingest server
```
export LD_LIBRARY_PATH=/opt/grpc/lib
ingest
```
