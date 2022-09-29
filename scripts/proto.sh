echo "generating protobuf codes"
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./api/proto/src/*.proto
  rm -rf ./api/proto/*.pb.go
  mv ./api/proto/src/*.pb.go ./api/proto/
  for f in ./api/proto/*.pb.go; do
    protoc-go-inject-tag -input="$f" -XXX_skip=json,xml,yaml
  done
