protoc -I . --openapiv2_out ./pkg \
    --openapiv2_opt logtostderr=true \
    ./api/api.proto