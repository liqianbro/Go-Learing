#!/bin/zsh
#protoc --go_out=. *.proto  --descriptor_set_out=pb.desc

protoc --go_out=. $(find protoc -iname "*.proto") --descriptor_set_out=pb.desc
