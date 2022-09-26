import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateBalance } from "./types/humans/tx";
import { MsgTranfserPoolcoin } from "./types/humans/tx";
import { MsgObservationVote } from "./types/humans/tx";
import { MsgKeysignVote } from "./types/humans/tx";
import { MsgApproveTransaction } from "./types/humans/tx";
import { MsgRequestTransaction } from "./types/humans/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/humansdotai.humans.humans.MsgUpdateBalance", MsgUpdateBalance],
    ["/humansdotai.humans.humans.MsgTranfserPoolcoin", MsgTranfserPoolcoin],
    ["/humansdotai.humans.humans.MsgObservationVote", MsgObservationVote],
    ["/humansdotai.humans.humans.MsgKeysignVote", MsgKeysignVote],
    ["/humansdotai.humans.humans.MsgApproveTransaction", MsgApproveTransaction],
    ["/humansdotai.humans.humans.MsgRequestTransaction", MsgRequestTransaction],
    
];

export { msgTypes }