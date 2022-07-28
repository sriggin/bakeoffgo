# BakeoffGo

## Setup

[Install Go](https://go.dev/doc/install)

```shell
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

See setup instructions for [catboost-go](https://github.com/bourbaki/catboost-go), and customize:
```shell
# you probably need to customize this
export CATBOOST_REPO=~/dev/thirdparty/catboost
export CATBOOST_DIR="$CATBOOST_REPO/catboost/libs/model_interface"
export C_INCLUDE_PATH=$CATBOOST_DIR:$C_INCLUDE_PATH
export LIBRARY_PATH=$CATBOOST_DIR:$LIBRARY_PATH
export LD_LIBRARY_PATH=$CATBOOST_DIR:$LD_LIBRARY_PATH
```