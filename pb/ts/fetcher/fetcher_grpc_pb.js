// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var fetcher_pb = require('./fetcher_pb.js');

function serialize_fetcher_FetchReply(arg) {
  if (!(arg instanceof fetcher_pb.FetchReply)) {
    throw new Error('Expected argument of type fetcher.FetchReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_fetcher_FetchReply(buffer_arg) {
  return fetcher_pb.FetchReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_fetcher_FetchRequest(arg) {
  if (!(arg instanceof fetcher_pb.FetchRequest)) {
    throw new Error('Expected argument of type fetcher.FetchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_fetcher_FetchRequest(buffer_arg) {
  return fetcher_pb.FetchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var FetcherService = exports.FetcherService = {
  fetch: {
    path: '/fetcher.Fetcher/Fetch',
    requestStream: false,
    responseStream: false,
    requestType: fetcher_pb.FetchRequest,
    responseType: fetcher_pb.FetchReply,
    requestSerialize: serialize_fetcher_FetchRequest,
    requestDeserialize: deserialize_fetcher_FetchRequest,
    responseSerialize: serialize_fetcher_FetchReply,
    responseDeserialize: deserialize_fetcher_FetchReply,
  },
};

exports.FetcherClient = grpc.makeGenericClientConstructor(FetcherService);
