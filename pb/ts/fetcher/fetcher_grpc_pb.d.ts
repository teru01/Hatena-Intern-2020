// package: fetcher
// file: fetcher.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as fetcher_pb from "./fetcher_pb";

interface IFetcherService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    fetch: IFetcherService_IFetch;
}

interface IFetcherService_IFetch extends grpc.MethodDefinition<fetcher_pb.FetchRequest, fetcher_pb.FetchReply> {
    path: string; // "/fetcher.Fetcher/Fetch"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<fetcher_pb.FetchRequest>;
    requestDeserialize: grpc.deserialize<fetcher_pb.FetchRequest>;
    responseSerialize: grpc.serialize<fetcher_pb.FetchReply>;
    responseDeserialize: grpc.deserialize<fetcher_pb.FetchReply>;
}

export const FetcherService: IFetcherService;

export interface IFetcherServer {
    fetch: grpc.handleUnaryCall<fetcher_pb.FetchRequest, fetcher_pb.FetchReply>;
}

export interface IFetcherClient {
    fetch(request: fetcher_pb.FetchRequest, callback: (error: grpc.ServiceError | null, response: fetcher_pb.FetchReply) => void): grpc.ClientUnaryCall;
    fetch(request: fetcher_pb.FetchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: fetcher_pb.FetchReply) => void): grpc.ClientUnaryCall;
    fetch(request: fetcher_pb.FetchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: fetcher_pb.FetchReply) => void): grpc.ClientUnaryCall;
}

export class FetcherClient extends grpc.Client implements IFetcherClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public fetch(request: fetcher_pb.FetchRequest, callback: (error: grpc.ServiceError | null, response: fetcher_pb.FetchReply) => void): grpc.ClientUnaryCall;
    public fetch(request: fetcher_pb.FetchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: fetcher_pb.FetchReply) => void): grpc.ClientUnaryCall;
    public fetch(request: fetcher_pb.FetchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: fetcher_pb.FetchReply) => void): grpc.ClientUnaryCall;
}
