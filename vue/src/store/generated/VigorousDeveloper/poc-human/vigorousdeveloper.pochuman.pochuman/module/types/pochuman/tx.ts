/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

export interface MsgRequestTransaction {
  creator: string;
  originChain: string;
  originAddress: string;
  targetChain: string;
  targetAddress: string;
  amount: string;
  fee: string;
}

export interface MsgRequestTransactionResponse {}

export interface MsgObservationVote {
  creator: string;
  txHash: string;
  chainName: string;
  from: string;
  to: string;
  amount: string;
}

export interface MsgObservationVoteResponse {}

const baseMsgRequestTransaction: object = {
  creator: "",
  originChain: "",
  originAddress: "",
  targetChain: "",
  targetAddress: "",
  amount: "",
  fee: "",
};

export const MsgRequestTransaction = {
  encode(
    message: MsgRequestTransaction,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.originChain !== "") {
      writer.uint32(18).string(message.originChain);
    }
    if (message.originAddress !== "") {
      writer.uint32(26).string(message.originAddress);
    }
    if (message.targetChain !== "") {
      writer.uint32(34).string(message.targetChain);
    }
    if (message.targetAddress !== "") {
      writer.uint32(42).string(message.targetAddress);
    }
    if (message.amount !== "") {
      writer.uint32(50).string(message.amount);
    }
    if (message.fee !== "") {
      writer.uint32(58).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestTransaction } as MsgRequestTransaction;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.originChain = reader.string();
          break;
        case 3:
          message.originAddress = reader.string();
          break;
        case 4:
          message.targetChain = reader.string();
          break;
        case 5:
          message.targetAddress = reader.string();
          break;
        case 6:
          message.amount = reader.string();
          break;
        case 7:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestTransaction {
    const message = { ...baseMsgRequestTransaction } as MsgRequestTransaction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = String(object.originChain);
    } else {
      message.originChain = "";
    }
    if (object.originAddress !== undefined && object.originAddress !== null) {
      message.originAddress = String(object.originAddress);
    } else {
      message.originAddress = "";
    }
    if (object.targetChain !== undefined && object.targetChain !== null) {
      message.targetChain = String(object.targetChain);
    } else {
      message.targetChain = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = String(object.targetAddress);
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgRequestTransaction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.originChain !== undefined &&
      (obj.originChain = message.originChain);
    message.originAddress !== undefined &&
      (obj.originAddress = message.originAddress);
    message.targetChain !== undefined &&
      (obj.targetChain = message.targetChain);
    message.targetAddress !== undefined &&
      (obj.targetAddress = message.targetAddress);
    message.amount !== undefined && (obj.amount = message.amount);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRequestTransaction>
  ): MsgRequestTransaction {
    const message = { ...baseMsgRequestTransaction } as MsgRequestTransaction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = object.originChain;
    } else {
      message.originChain = "";
    }
    if (object.originAddress !== undefined && object.originAddress !== null) {
      message.originAddress = object.originAddress;
    } else {
      message.originAddress = "";
    }
    if (object.targetChain !== undefined && object.targetChain !== null) {
      message.targetChain = object.targetChain;
    } else {
      message.targetChain = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = object.targetAddress;
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgRequestTransactionResponse: object = {};

export const MsgRequestTransactionResponse = {
  encode(
    _: MsgRequestTransactionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRequestTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestTransactionResponse,
    } as MsgRequestTransactionResponse;
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

  fromJSON(_: any): MsgRequestTransactionResponse {
    const message = {
      ...baseMsgRequestTransactionResponse,
    } as MsgRequestTransactionResponse;
    return message;
  },

  toJSON(_: MsgRequestTransactionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRequestTransactionResponse>
  ): MsgRequestTransactionResponse {
    const message = {
      ...baseMsgRequestTransactionResponse,
    } as MsgRequestTransactionResponse;
    return message;
  },
};

const baseMsgObservationVote: object = {
  creator: "",
  txHash: "",
  chainName: "",
  from: "",
  to: "",
  amount: "",
};

export const MsgObservationVote = {
  encode(
    message: MsgObservationVote,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.txHash !== "") {
      writer.uint32(18).string(message.txHash);
    }
    if (message.chainName !== "") {
      writer.uint32(26).string(message.chainName);
    }
    if (message.from !== "") {
      writer.uint32(34).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(42).string(message.to);
    }
    if (message.amount !== "") {
      writer.uint32(50).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgObservationVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgObservationVote } as MsgObservationVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.txHash = reader.string();
          break;
        case 3:
          message.chainName = reader.string();
          break;
        case 4:
          message.from = reader.string();
          break;
        case 5:
          message.to = reader.string();
          break;
        case 6:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgObservationVote {
    const message = { ...baseMsgObservationVote } as MsgObservationVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = String(object.txHash);
    } else {
      message.txHash = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName);
    } else {
      message.chainName = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = String(object.from);
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = String(object.to);
    } else {
      message.to = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: MsgObservationVote): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.txHash !== undefined && (obj.txHash = message.txHash);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgObservationVote>): MsgObservationVote {
    const message = { ...baseMsgObservationVote } as MsgObservationVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = object.txHash;
    } else {
      message.txHash = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName;
    } else {
      message.chainName = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = object.from;
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = object.to;
    } else {
      message.to = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMsgObservationVoteResponse: object = {};

export const MsgObservationVoteResponse = {
  encode(
    _: MsgObservationVoteResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgObservationVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgObservationVoteResponse,
    } as MsgObservationVoteResponse;
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

  fromJSON(_: any): MsgObservationVoteResponse {
    const message = {
      ...baseMsgObservationVoteResponse,
    } as MsgObservationVoteResponse;
    return message;
  },

  toJSON(_: MsgObservationVoteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgObservationVoteResponse>
  ): MsgObservationVoteResponse {
    const message = {
      ...baseMsgObservationVoteResponse,
    } as MsgObservationVoteResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestTransaction(
    request: MsgRequestTransaction
  ): Promise<MsgRequestTransactionResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ObservationVote(
    request: MsgObservationVote
  ): Promise<MsgObservationVoteResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RequestTransaction(
    request: MsgRequestTransaction
  ): Promise<MsgRequestTransactionResponse> {
    const data = MsgRequestTransaction.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Msg",
      "RequestTransaction",
      data
    );
    return promise.then((data) =>
      MsgRequestTransactionResponse.decode(new Reader(data))
    );
  }

  ObservationVote(
    request: MsgObservationVote
  ): Promise<MsgObservationVoteResponse> {
    const data = MsgObservationVote.encode(request).finish();
    const promise = this.rpc.request(
      "vigorousdeveloper.pochuman.pochuman.Msg",
      "ObservationVote",
      data
    );
    return promise.then((data) =>
      MsgObservationVoteResponse.decode(new Reader(data))
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
