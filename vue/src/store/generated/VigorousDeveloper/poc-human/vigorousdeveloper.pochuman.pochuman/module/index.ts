// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestTransaction } from "./types/humans/tx";
import { MsgTranfserPoolcoin } from "./types/humans/tx";
import { MsgKeysignVote } from "./types/humans/tx";
import { MsgApproveTransaction } from "./types/humans/tx";
import { MsgObservationVote } from "./types/humans/tx";
import { MsgUpdateBalance } from "./types/humans/tx";


const types = [
  ["/vigorousdeveloper.humans.humans.MsgRequestTransaction", MsgRequestTransaction],
  ["/vigorousdeveloper.humans.humans.MsgTranfserPoolcoin", MsgTranfserPoolcoin],
  ["/vigorousdeveloper.humans.humans.MsgKeysignVote", MsgKeysignVote],
  ["/vigorousdeveloper.humans.humans.MsgApproveTransaction", MsgApproveTransaction],
  ["/vigorousdeveloper.humans.humans.MsgObservationVote", MsgObservationVote],
  ["/vigorousdeveloper.humans.humans.MsgUpdateBalance", MsgUpdateBalance],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRequestTransaction: (data: MsgRequestTransaction): EncodeObject => ({ typeUrl: "/vigorousdeveloper.humans.humans.MsgRequestTransaction", value: MsgRequestTransaction.fromPartial( data ) }),
    msgTranfserPoolcoin: (data: MsgTranfserPoolcoin): EncodeObject => ({ typeUrl: "/vigorousdeveloper.humans.humans.MsgTranfserPoolcoin", value: MsgTranfserPoolcoin.fromPartial( data ) }),
    msgKeysignVote: (data: MsgKeysignVote): EncodeObject => ({ typeUrl: "/vigorousdeveloper.humans.humans.MsgKeysignVote", value: MsgKeysignVote.fromPartial( data ) }),
    msgApproveTransaction: (data: MsgApproveTransaction): EncodeObject => ({ typeUrl: "/vigorousdeveloper.humans.humans.MsgApproveTransaction", value: MsgApproveTransaction.fromPartial( data ) }),
    msgObservationVote: (data: MsgObservationVote): EncodeObject => ({ typeUrl: "/vigorousdeveloper.humans.humans.MsgObservationVote", value: MsgObservationVote.fromPartial( data ) }),
    msgUpdateBalance: (data: MsgUpdateBalance): EncodeObject => ({ typeUrl: "/vigorousdeveloper.humans.humans.MsgUpdateBalance", value: MsgUpdateBalance.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
