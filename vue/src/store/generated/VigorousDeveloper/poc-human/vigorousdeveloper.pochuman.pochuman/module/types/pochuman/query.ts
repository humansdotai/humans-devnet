/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../pochuman/params";
import { FeeBalance } from "../pochuman/fee_balance";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { KeysignVoteData } from "../pochuman/keysign_vote_data";
import { ObserveVote } from "../pochuman/observe_vote";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetFeeBalanceRequest {
  index: string;
}

export interface QueryGetFeeBalanceResponse {
  feeBalance: FeeBalance | undefined;
}

export interface QueryAllFeeBalanceRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFeeBalanceResponse {
  feeBalance: FeeBalance[];
  pagination: PageResponse | undefined;
}

export interface QueryGetKeysignVoteDataRequest {
  index: string;
}

export interface QueryGetKeysignVoteDataResponse {
  keysignVoteData: KeysignVoteData | undefined;
}

export interface QueryAllKeysignVoteDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllKeysignVoteDataResponse {
  keysignVoteData: KeysignVoteData[];
  pagination: PageResponse | undefined;
}

export interface QueryGetObserveVoteRequest {
  index: string;
}

export interface QueryGetObserveVoteResponse {
  observeVote: ObserveVote | undefined;
}

export interface QueryAllObserveVoteRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllObserveVoteResponse {
  observeVote: ObserveVote[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetFeeBalanceRequest: object = { index: "" };

export const QueryGetFeeBalanceRequest = {
  encode(
    message: QueryGetFeeBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFeeBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFeeBalanceRequest,
    } as QueryGetFeeBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeBalanceRequest {
    const message = {
      ...baseQueryGetFeeBalanceRequest,
    } as QueryGetFeeBalanceRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetFeeBalanceRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeBalanceRequest>
  ): QueryGetFeeBalanceRequest {
    const message = {
      ...baseQueryGetFeeBalanceRequest,
    } as QueryGetFeeBalanceRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetFeeBalanceResponse: object = {};

export const QueryGetFeeBalanceResponse = {
  encode(
    message: QueryGetFeeBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.feeBalance !== undefined) {
      FeeBalance.encode(message.feeBalance, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFeeBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFeeBalanceResponse,
    } as QueryGetFeeBalanceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.feeBalance = FeeBalance.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeBalanceResponse {
    const message = {
      ...baseQueryGetFeeBalanceResponse,
    } as QueryGetFeeBalanceResponse;
    if (object.feeBalance !== undefined && object.feeBalance !== null) {
      message.feeBalance = FeeBalance.fromJSON(object.feeBalance);
    } else {
      message.feeBalance = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFeeBalanceResponse): unknown {
    const obj: any = {};
    message.feeBalance !== undefined &&
      (obj.feeBalance = message.feeBalance
        ? FeeBalance.toJSON(message.feeBalance)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeBalanceResponse>
  ): QueryGetFeeBalanceResponse {
    const message = {
      ...baseQueryGetFeeBalanceResponse,
    } as QueryGetFeeBalanceResponse;
    if (object.feeBalance !== undefined && object.feeBalance !== null) {
      message.feeBalance = FeeBalance.fromPartial(object.feeBalance);
    } else {
      message.feeBalance = undefined;
    }
    return message;
  },
};

const baseQueryAllFeeBalanceRequest: object = {};

export const QueryAllFeeBalanceRequest = {
  encode(
    message: QueryAllFeeBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllFeeBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFeeBalanceRequest,
    } as QueryAllFeeBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeBalanceRequest {
    const message = {
      ...baseQueryAllFeeBalanceRequest,
    } as QueryAllFeeBalanceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFeeBalanceRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFeeBalanceRequest>
  ): QueryAllFeeBalanceRequest {
    const message = {
      ...baseQueryAllFeeBalanceRequest,
    } as QueryAllFeeBalanceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFeeBalanceResponse: object = {};

export const QueryAllFeeBalanceResponse = {
  encode(
    message: QueryAllFeeBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.feeBalance) {
      FeeBalance.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllFeeBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFeeBalanceResponse,
    } as QueryAllFeeBalanceResponse;
    message.feeBalance = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.feeBalance.push(FeeBalance.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeBalanceResponse {
    const message = {
      ...baseQueryAllFeeBalanceResponse,
    } as QueryAllFeeBalanceResponse;
    message.feeBalance = [];
    if (object.feeBalance !== undefined && object.feeBalance !== null) {
      for (const e of object.feeBalance) {
        message.feeBalance.push(FeeBalance.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFeeBalanceResponse): unknown {
    const obj: any = {};
    if (message.feeBalance) {
      obj.feeBalance = message.feeBalance.map((e) =>
        e ? FeeBalance.toJSON(e) : undefined
      );
    } else {
      obj.feeBalance = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFeeBalanceResponse>
  ): QueryAllFeeBalanceResponse {
    const message = {
      ...baseQueryAllFeeBalanceResponse,
    } as QueryAllFeeBalanceResponse;
    message.feeBalance = [];
    if (object.feeBalance !== undefined && object.feeBalance !== null) {
      for (const e of object.feeBalance) {
        message.feeBalance.push(FeeBalance.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetKeysignVoteDataRequest: object = { index: "" };

export const QueryGetKeysignVoteDataRequest = {
  encode(
    message: QueryGetKeysignVoteDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetKeysignVoteDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetKeysignVoteDataRequest,
    } as QueryGetKeysignVoteDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetKeysignVoteDataRequest {
    const message = {
      ...baseQueryGetKeysignVoteDataRequest,
    } as QueryGetKeysignVoteDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetKeysignVoteDataRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetKeysignVoteDataRequest>
  ): QueryGetKeysignVoteDataRequest {
    const message = {
      ...baseQueryGetKeysignVoteDataRequest,
    } as QueryGetKeysignVoteDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetKeysignVoteDataResponse: object = {};

export const QueryGetKeysignVoteDataResponse = {
  encode(
    message: QueryGetKeysignVoteDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.keysignVoteData !== undefined) {
      KeysignVoteData.encode(
        message.keysignVoteData,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetKeysignVoteDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetKeysignVoteDataResponse,
    } as QueryGetKeysignVoteDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keysignVoteData = KeysignVoteData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetKeysignVoteDataResponse {
    const message = {
      ...baseQueryGetKeysignVoteDataResponse,
    } as QueryGetKeysignVoteDataResponse;
    if (
      object.keysignVoteData !== undefined &&
      object.keysignVoteData !== null
    ) {
      message.keysignVoteData = KeysignVoteData.fromJSON(
        object.keysignVoteData
      );
    } else {
      message.keysignVoteData = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetKeysignVoteDataResponse): unknown {
    const obj: any = {};
    message.keysignVoteData !== undefined &&
      (obj.keysignVoteData = message.keysignVoteData
        ? KeysignVoteData.toJSON(message.keysignVoteData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetKeysignVoteDataResponse>
  ): QueryGetKeysignVoteDataResponse {
    const message = {
      ...baseQueryGetKeysignVoteDataResponse,
    } as QueryGetKeysignVoteDataResponse;
    if (
      object.keysignVoteData !== undefined &&
      object.keysignVoteData !== null
    ) {
      message.keysignVoteData = KeysignVoteData.fromPartial(
        object.keysignVoteData
      );
    } else {
      message.keysignVoteData = undefined;
    }
    return message;
  },
};

const baseQueryAllKeysignVoteDataRequest: object = {};

export const QueryAllKeysignVoteDataRequest = {
  encode(
    message: QueryAllKeysignVoteDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllKeysignVoteDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllKeysignVoteDataRequest,
    } as QueryAllKeysignVoteDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllKeysignVoteDataRequest {
    const message = {
      ...baseQueryAllKeysignVoteDataRequest,
    } as QueryAllKeysignVoteDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllKeysignVoteDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllKeysignVoteDataRequest>
  ): QueryAllKeysignVoteDataRequest {
    const message = {
      ...baseQueryAllKeysignVoteDataRequest,
    } as QueryAllKeysignVoteDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllKeysignVoteDataResponse: object = {};

export const QueryAllKeysignVoteDataResponse = {
  encode(
    message: QueryAllKeysignVoteDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.keysignVoteData) {
      KeysignVoteData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllKeysignVoteDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllKeysignVoteDataResponse,
    } as QueryAllKeysignVoteDataResponse;
    message.keysignVoteData = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keysignVoteData.push(
            KeysignVoteData.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllKeysignVoteDataResponse {
    const message = {
      ...baseQueryAllKeysignVoteDataResponse,
    } as QueryAllKeysignVoteDataResponse;
    message.keysignVoteData = [];
    if (
      object.keysignVoteData !== undefined &&
      object.keysignVoteData !== null
    ) {
      for (const e of object.keysignVoteData) {
        message.keysignVoteData.push(KeysignVoteData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllKeysignVoteDataResponse): unknown {
    const obj: any = {};
    if (message.keysignVoteData) {
      obj.keysignVoteData = message.keysignVoteData.map((e) =>
        e ? KeysignVoteData.toJSON(e) : undefined
      );
    } else {
      obj.keysignVoteData = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllKeysignVoteDataResponse>
  ): QueryAllKeysignVoteDataResponse {
    const message = {
      ...baseQueryAllKeysignVoteDataResponse,
    } as QueryAllKeysignVoteDataResponse;
    message.keysignVoteData = [];
    if (
      object.keysignVoteData !== undefined &&
      object.keysignVoteData !== null
    ) {
      for (const e of object.keysignVoteData) {
        message.keysignVoteData.push(KeysignVoteData.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetObserveVoteRequest: object = { index: "" };

export const QueryGetObserveVoteRequest = {
  encode(
    message: QueryGetObserveVoteRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetObserveVoteRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetObserveVoteRequest,
    } as QueryGetObserveVoteRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetObserveVoteRequest {
    const message = {
      ...baseQueryGetObserveVoteRequest,
    } as QueryGetObserveVoteRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetObserveVoteRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetObserveVoteRequest>
  ): QueryGetObserveVoteRequest {
    const message = {
      ...baseQueryGetObserveVoteRequest,
    } as QueryGetObserveVoteRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetObserveVoteResponse: object = {};

export const QueryGetObserveVoteResponse = {
  encode(
    message: QueryGetObserveVoteResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.observeVote !== undefined) {
      ObserveVote.encode(
        message.observeVote,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetObserveVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetObserveVoteResponse,
    } as QueryGetObserveVoteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.observeVote = ObserveVote.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetObserveVoteResponse {
    const message = {
      ...baseQueryGetObserveVoteResponse,
    } as QueryGetObserveVoteResponse;
    if (object.observeVote !== undefined && object.observeVote !== null) {
      message.observeVote = ObserveVote.fromJSON(object.observeVote);
    } else {
      message.observeVote = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetObserveVoteResponse): unknown {
    const obj: any = {};
    message.observeVote !== undefined &&
      (obj.observeVote = message.observeVote
        ? ObserveVote.toJSON(message.observeVote)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetObserveVoteResponse>
  ): QueryGetObserveVoteResponse {
    const message = {
      ...baseQueryGetObserveVoteResponse,
    } as QueryGetObserveVoteResponse;
    if (object.observeVote !== undefined && object.observeVote !== null) {
      message.observeVote = ObserveVote.fromPartial(object.observeVote);
    } else {
      message.observeVote = undefined;
    }
    return message;
  },
};

const baseQueryAllObserveVoteRequest: object = {};

export const QueryAllObserveVoteRequest = {
  encode(
    message: QueryAllObserveVoteRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllObserveVoteRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllObserveVoteRequest,
    } as QueryAllObserveVoteRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllObserveVoteRequest {
    const message = {
      ...baseQueryAllObserveVoteRequest,
    } as QueryAllObserveVoteRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllObserveVoteRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllObserveVoteRequest>
  ): QueryAllObserveVoteRequest {
    const message = {
      ...baseQueryAllObserveVoteRequest,
    } as QueryAllObserveVoteRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllObserveVoteResponse: object = {};

export const QueryAllObserveVoteResponse = {
  encode(
    message: QueryAllObserveVoteResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.observeVote) {
      ObserveVote.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllObserveVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllObserveVoteResponse,
    } as QueryAllObserveVoteResponse;
    message.observeVote = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.observeVote.push(ObserveVote.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllObserveVoteResponse {
    const message = {
      ...baseQueryAllObserveVoteResponse,
    } as QueryAllObserveVoteResponse;
    message.observeVote = [];
    if (object.observeVote !== undefined && object.observeVote !== null) {
      for (const e of object.observeVote) {
        message.observeVote.push(ObserveVote.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllObserveVoteResponse): unknown {
    const obj: any = {};
    if (message.observeVote) {
      obj.observeVote = message.observeVote.map((e) =>
        e ? ObserveVote.toJSON(e) : undefined
      );
    } else {
      obj.observeVote = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllObserveVoteResponse>
  ): QueryAllObserveVoteResponse {
    const message = {
      ...baseQueryAllObserveVoteResponse,
    } as QueryAllObserveVoteResponse;
    message.observeVote = [];
    if (object.observeVote !== undefined && object.observeVote !== null) {
      for (const e of object.observeVote) {
        message.observeVote.push(ObserveVote.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a FeeBalance by index. */
  FeeBalance(
    request: QueryGetFeeBalanceRequest
  ): Promise<QueryGetFeeBalanceResponse>;
  /** Queries a list of FeeBalance items. */
  FeeBalanceAll(
    request: QueryAllFeeBalanceRequest
  ): Promise<QueryAllFeeBalanceResponse>;
  /** Queries a KeysignVoteData by index. */
  KeysignVoteData(
    request: QueryGetKeysignVoteDataRequest
  ): Promise<QueryGetKeysignVoteDataResponse>;
  /** Queries a list of KeysignVoteData items. */
  KeysignVoteDataAll(
    request: QueryAllKeysignVoteDataRequest
  ): Promise<QueryAllKeysignVoteDataResponse>;
  /** Queries a ObserveVote by index. */
  ObserveVote(
    request: QueryGetObserveVoteRequest
  ): Promise<QueryGetObserveVoteResponse>;
  /** Queries a list of ObserveVote items. */
  ObserveVoteAll(
    request: QueryAllObserveVoteRequest
  ): Promise<QueryAllObserveVoteResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  FeeBalance(
    request: QueryGetFeeBalanceRequest
  ): Promise<QueryGetFeeBalanceResponse> {
    const data = QueryGetFeeBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "FeeBalance",
      data
    );
    return promise.then((data) =>
      QueryGetFeeBalanceResponse.decode(new Reader(data))
    );
  }

  FeeBalanceAll(
    request: QueryAllFeeBalanceRequest
  ): Promise<QueryAllFeeBalanceResponse> {
    const data = QueryAllFeeBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "FeeBalanceAll",
      data
    );
    return promise.then((data) =>
      QueryAllFeeBalanceResponse.decode(new Reader(data))
    );
  }

  KeysignVoteData(
    request: QueryGetKeysignVoteDataRequest
  ): Promise<QueryGetKeysignVoteDataResponse> {
    const data = QueryGetKeysignVoteDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "KeysignVoteData",
      data
    );
    return promise.then((data) =>
      QueryGetKeysignVoteDataResponse.decode(new Reader(data))
    );
  }

  KeysignVoteDataAll(
    request: QueryAllKeysignVoteDataRequest
  ): Promise<QueryAllKeysignVoteDataResponse> {
    const data = QueryAllKeysignVoteDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "KeysignVoteDataAll",
      data
    );
    return promise.then((data) =>
      QueryAllKeysignVoteDataResponse.decode(new Reader(data))
    );
  }

  ObserveVote(
    request: QueryGetObserveVoteRequest
  ): Promise<QueryGetObserveVoteResponse> {
    const data = QueryGetObserveVoteRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "ObserveVote",
      data
    );
    return promise.then((data) =>
      QueryGetObserveVoteResponse.decode(new Reader(data))
    );
  }

  ObserveVoteAll(
    request: QueryAllObserveVoteRequest
  ): Promise<QueryAllObserveVoteResponse> {
    const data = QueryAllObserveVoteRequest.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Query",
      "ObserveVoteAll",
      data
    );
    return promise.then((data) =>
      QueryAllObserveVoteResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
