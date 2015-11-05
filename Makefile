%.pb.go : %.proto
	protoc --go_out=plugins=grpc:. $<

lib/add.rb lib/add_services.rb : proto/add.proto
	protoc -I proto --ruby_out=lib --grpc_out=lib --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` proto/add.proto

proto_targets = proto/add.pb.go proto/add_better.pb.go lib/add.rb lib/add_services.rb
all: $(proto_targets)

clean:
	rm $(proto_targets)
