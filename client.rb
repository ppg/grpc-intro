this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'add_services'

def main
  stub = Main::Test::Stub.new('localhost:1234')
  r = stub.add(Main::NumericRequest.new(v1: ARGV[0].to_i, v2: ARGV[1].to_i)).r
  p "r: #{r}"
end

main
