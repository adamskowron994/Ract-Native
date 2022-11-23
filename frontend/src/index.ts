import * as grpcWeb from 'grpc-web';
import {ElizaServiceClient} from './proto/ElizaServiceClientPb';
import {SayRequest} from './proto/eliza_pb';


const elizaService = new ElizaServiceClient("http://localhost:8080")
const req = new SayRequest();
req.setSentence("Hi!");

async function main() {
    const res = await elizaService.say(req, {});
    console.log(res);
}
main()
