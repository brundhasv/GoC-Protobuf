# GoC Protobuf Encoding/Decoding
Go functions for Protobuf Encoding/decoding integrated with C

# Usage:
protoc --go_out=. Student.proto

go build  -o student.so -buildmode=c-shared Student.pb.go student_en_dc.go

gcc -o student student_cgo.c ./student.so

# Go program can also be run in standalone mode:
go run student_en_dc.go Student.pb.go

# Explanation of Code here:
https://dev.to/brundhasv/go-protobuf-encoding-decoding-integrated-with-c-3ig0


