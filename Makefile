lib/add.pb.go : lib/add.proto
	protoc --go_out=plugins=grpc:. lib/add.proto

lib/add_better.pb.go : lib/add_better.proto
	protoc --go_out=plugins=grpc:. lib/add_better.proto

lib/add.rb : lib/add.proto
	protoc --ruby_out=lib --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` lib/add.proto

lib/add_services.rb : lib/add.proto
	protoc --ruby_out=lib --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` lib/add.proto

all: lib/add.pb.go lib/add_better.pb.go lib/add.rb lib/add_services.rb

clean:
	rm lib/add.pb.go lib/add_better.pb.go lib/add.rb lib/add_services.rb
