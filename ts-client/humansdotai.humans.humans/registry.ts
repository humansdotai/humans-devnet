import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgObservationVote } from "./types/humans/tx";
import { MsgApproveTransaction } from "./types/humans/tx";
import { MsgUpdateBalance } from "./types/humans/tx";
import { MsgKeysignVote } from "./types/humans/tx";
import { MsgRequestTransaction } from "./types/humans/tx";
import { MsgTranfserPoolcoin } from "./types/humans/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/humansdotai.humans.humans.MsgObservationVote", MsgObservationVote],
    ["/humansdotai.humans.humans.MsgApproveTransaction", MsgApproveTransaction],
    ["/humansdotai.humans.humans.MsgUpdateBalance", MsgUpdateBalance],
    ["/humansdotai.humans.humans.MsgKeysignVote", MsgKeysignVote],
    ["/humansdotai.humans.humans.MsgRequestTransaction", MsgRequestTransaction],
    ["/humansdotai.humans.humans.MsgTranfserPoolcoin", MsgTranfserPoolcoin],
    
];

export { msgTypes }