import "./proto/eliza_grpc_web_pb.js"
import "./proto/eliza_pb.js"
//
// let introFinished = false;
//
// // Make the Eliza Service client
// const client = proto.buf.connect.demo.eliza.v1.ElizaServiceClient(
//     "https://localhost:8081"
// )
// var req = new proto.buf.connect.demo.eliza.v1.IntroduceRequest();
// var stream = client.();
// stream.on('data', function(response) {
//     console.log(response.getMessage());
// });
// stream.on('status', function(status) {
//     console.log(status.code);
//     console.log(status.details);
//     console.log(status.metadata);
// });
// stream.on('end', function(end) {
//     // stream end signal
// });
//
// // to close the stream
// stream.cancel()