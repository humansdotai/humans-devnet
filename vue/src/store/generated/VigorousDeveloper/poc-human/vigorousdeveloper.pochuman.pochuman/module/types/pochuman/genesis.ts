/* eslint-disable */
import { Params } from "../pochuman/params";
import { FeeBalance } from "../pochuman/fee_balance";
import { KeysignVoteData } from "../pochuman/keysign_vote_data";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "vigorousdeveloper.pochuman.pochuman";

/** GenesisState defines the pochuman module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  feeBalanceList: FeeBalance[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  keysignVoteDataList: KeysignVoteData[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.feeBalanceList) {
      FeeBalance.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.keysignVoteDataList) {
      KeysignVoteData.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.feeBalanceList = [];
    message.keysignVoteDataList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.feeBalanceList.push(
            FeeBalance.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.keysignVoteDataList.push(
            KeysignVoteData.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.feeBalanceList = [];
    message.keysignVoteDataList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.feeBalanceList !== undefined && object.feeBalanceList !== null) {
      for (const e of object.feeBalanceList) {
        message.feeBalanceList.push(FeeBalance.fromJSON(e));
      }
    }
    if (
      object.keysignVoteDataList !== undefined &&
      object.keysignVoteDataList !== null
    ) {
      for (const e of object.keysignVoteDataList) {
        message.keysignVoteDataList.push(KeysignVoteData.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.feeBalanceList) {
      obj.feeBalanceList = message.feeBalanceList.map((e) =>
        e ? FeeBalance.toJSON(e) : undefined
      );
    } else {
      obj.feeBalanceList = [];
    }
    if (message.keysignVoteDataList) {
      obj.keysignVoteDataList = message.keysignVoteDataList.map((e) =>
        e ? KeysignVoteData.toJSON(e) : undefined
      );
    } else {
      obj.keysignVoteDataList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.feeBalanceList = [];
    message.keysignVoteDataList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.feeBalanceList !== undefined && object.feeBalanceList !== null) {
      for (const e of object.feeBalanceList) {
        message.feeBalanceList.push(FeeBalance.fromPartial(e));
      }
    }
    if (
      object.keysignVoteDataList !== undefined &&
      object.keysignVoteDataList !== null
    ) {
      for (const e of object.keysignVoteDataList) {
        message.keysignVoteDataList.push(KeysignVoteData.fromPartial(e));
      }
    }
    return message;
  },
};

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
