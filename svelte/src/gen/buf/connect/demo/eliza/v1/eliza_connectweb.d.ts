// Copyright 2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-connect-web v0.0.9
// @generated from file buf/connect/demo/eliza/v1/eliza.proto (package buf.connect.demo.eliza.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {
    ConverseRequest,
    ConverseResponse,
    IntroduceRequest,
    IntroduceResponse,
    SayRequest,
    SayResponse,
} from './eliza_pb.js'
import { MethodKind } from '@bufbuild/protobuf'

/**
 * ElizaService provides a way to talk to the ELIZA, which is a port of
 * the DOCTOR script for Joseph Weizenbaum's original ELIZA program.
 * Created in the mid-1960s at the MIT Artificial Intelligence Laboratory,
 * ELIZA demonstrates the superficiality of human-computer communication.
 * DOCTOR simulates a psychotherapist, and is commonly found as an Easter
 * egg in emacs distributions.
 *
 * @generated from service buf.connect.demo.eliza.v1.ElizaService
 */
export declare const ElizaService: {
    readonly typeName: 'buf.connect.demo.eliza.v1.ElizaService'
    readonly methods: {
        /**
         * Say is a unary request demo. This method should allow for a one sentence
         * response given a one sentence request.
         *
         * @generated from rpc buf.connect.demo.eliza.v1.ElizaService.Say
         */
        readonly say: {
            readonly name: 'Say'
            readonly I: typeof SayRequest
            readonly O: typeof SayResponse
            readonly kind: MethodKind.Unary
        }
        /**
         * Converse is a bi-directional streaming request demo. This method should allow for
         * many requests and many responses.
         *
         * @generated from rpc buf.connect.demo.eliza.v1.ElizaService.Converse
         */
        readonly converse: {
            readonly name: 'Converse'
            readonly I: typeof ConverseRequest
            readonly O: typeof ConverseResponse
            readonly kind: MethodKind.BiDiStreaming
        }
        /**
         * Introduce is a server-streaming request demo.  This method allows for a single request that will return a series
         * of responses
         *
         * @generated from rpc buf.connect.demo.eliza.v1.ElizaService.Introduce
         */
        readonly introduce: {
            readonly name: 'Introduce'
            readonly I: typeof IntroduceRequest
            readonly O: typeof IntroduceResponse
            readonly kind: MethodKind.ServerStreaming
        }
    }
}
