var grpc = require('grpc');
var pb = grpc.load(__dirname + '/add.proto');
pb = pb.main;

var testClient = new pb.Test('localhost:1234', grpc.Credentials.createInsecure());

testClient.add({v1: parseInt(process.argv[2]), v2: parseInt(process.argv[3])}, function(error, response) {
  if (error) {
    console.log('error', error);
  } else {
    console.log('response', response);
  }
});
