# ingest

Accept client gRPC request and save data into rocksdb.

## System requirements

* Ubuntu18.04
* gcc 7.5.0+
* cmake 3.17.0+
* docker with buildkit support
```
{
  "features": {
    "buildkit": true
  }
}
```

## Build and install ingest server

```
cd src/main/cpp
mkdir build
cd build
cmake ..
make
sudo make install
```

## Start ingest server
```
export LD_LIBRARY_PATH=/opt/grpc/lib
ingest
```

## Run test application
```
cd src/test/go
docker build . -t test
docker run --rm test
```
