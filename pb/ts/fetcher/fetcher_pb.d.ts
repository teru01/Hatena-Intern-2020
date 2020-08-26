// package: fetcher
// file: fetcher.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class FetchRequest extends jspb.Message { 
    getUri(): string;
    setUri(value: string): FetchRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FetchRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FetchRequest): FetchRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FetchRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FetchRequest;
    static deserializeBinaryFromReader(message: FetchRequest, reader: jspb.BinaryReader): FetchRequest;
}

export namespace FetchRequest {
    export type AsObject = {
        uri: string,
    }
}

export class FetchReply extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): FetchReply;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FetchReply.AsObject;
    static toObject(includeInstance: boolean, msg: FetchReply): FetchReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FetchReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FetchReply;
    static deserializeBinaryFromReader(message: FetchReply, reader: jspb.BinaryReader): FetchReply;
}

export namespace FetchReply {
    export type AsObject = {
        title: string,
    }
}
