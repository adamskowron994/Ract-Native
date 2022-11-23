"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const ElizaServiceClientPb_1 = require("./proto/ElizaServiceClientPb");
const eliza_pb_1 = require("./proto/eliza_pb");
const bun_promptx_1 = require("bun-promptx");
const elizaService = new ElizaServiceClientPb_1.ElizaServiceClient("http://localhost:8080");
const req = new eliza_pb_1.SayRequest();
req.setSentence("Hi!");
function main() {
    return __awaiter(this, void 0, void 0, function* () {
        const res = yield elizaService.say(req, {});
        console.log(res);
    });
}
main();
const username = (0, bun_promptx_1.createPrompt)("Enter username: ");
