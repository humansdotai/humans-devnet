import { Client, registry, MissingWalletError } from 'humansdotai-humans-client-ts'

import { FeeBalance } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { KeysignVoteData } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { ObserveVote } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { Params } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { PoolBalance } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { Pubkeys } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { QueryGetSuperadminRequest } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { QueryGetSuperadminResponse } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { QueryAllSuperadminRequest } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { QueryAllSuperadminResponse } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { Superadmin } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { TransactionData } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"
import { WhitelistedNode } from "humansdotai-humans-client-ts/humansdotai.humans.humans/types"


export { FeeBalance, KeysignVoteData, ObserveVote, Params, PoolBalance, Pubkeys, QueryGetSuperadminRequest, QueryGetSuperadminResponse, QueryAllSuperadminRequest, QueryAllSuperadminResponse, Superadmin, TransactionData, WhitelistedNode };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				FeeBalance: {},
				FeeBalanceAll: {},
				KeysignVoteData: {},
				KeysignVoteDataAll: {},
				ObserveVote: {},
				ObserveVoteAll: {},
				PoolBalance: {},
				PoolBalanceAll: {},
				TransactionData: {},
				TransactionDataAll: {},
				Pubkeys: {},
				PubkeysAll: {},
				WhitelistedNode: {},
				WhitelistedNodeAll: {},
				
				_Structure: {
						FeeBalance: getStructure(FeeBalance.fromPartial({})),
						KeysignVoteData: getStructure(KeysignVoteData.fromPartial({})),
						ObserveVote: getStructure(ObserveVote.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PoolBalance: getStructure(PoolBalance.fromPartial({})),
						Pubkeys: getStructure(Pubkeys.fromPartial({})),
						QueryGetSuperadminRequest: getStructure(QueryGetSuperadminRequest.fromPartial({})),
						QueryGetSuperadminResponse: getStructure(QueryGetSuperadminResponse.fromPartial({})),
						QueryAllSuperadminRequest: getStructure(QueryAllSuperadminRequest.fromPartial({})),
						QueryAllSuperadminResponse: getStructure(QueryAllSuperadminResponse.fromPartial({})),
						Superadmin: getStructure(Superadmin.fromPartial({})),
						TransactionData: getStructure(TransactionData.fromPartial({})),
						WhitelistedNode: getStructure(WhitelistedNode.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getFeeBalance: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FeeBalance[JSON.stringify(params)] ?? {}
		},
				getFeeBalanceAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FeeBalanceAll[JSON.stringify(params)] ?? {}
		},
				getKeysignVoteData: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.KeysignVoteData[JSON.stringify(params)] ?? {}
		},
				getKeysignVoteDataAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.KeysignVoteDataAll[JSON.stringify(params)] ?? {}
		},
				getObserveVote: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ObserveVote[JSON.stringify(params)] ?? {}
		},
				getObserveVoteAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ObserveVoteAll[JSON.stringify(params)] ?? {}
		},
				getPoolBalance: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PoolBalance[JSON.stringify(params)] ?? {}
		},
				getPoolBalanceAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PoolBalanceAll[JSON.stringify(params)] ?? {}
		},
				getTransactionData: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TransactionData[JSON.stringify(params)] ?? {}
		},
				getTransactionDataAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TransactionDataAll[JSON.stringify(params)] ?? {}
		},
				getPubkeys: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Pubkeys[JSON.stringify(params)] ?? {}
		},
				getPubkeysAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PubkeysAll[JSON.stringify(params)] ?? {}
		},
				getWhitelistedNode: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.WhitelistedNode[JSON.stringify(params)] ?? {}
		},
				getWhitelistedNodeAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.WhitelistedNodeAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: humansdotai.humans.humans initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFeeBalance({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryFeeBalance( key.index)).data
				
					
				commit('QUERY', { query: 'FeeBalance', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFeeBalance', payload: { options: { all }, params: {...key},query }})
				return getters['getFeeBalance']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFeeBalance API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFeeBalanceAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryFeeBalanceAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryFeeBalanceAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'FeeBalanceAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFeeBalanceAll', payload: { options: { all }, params: {...key},query }})
				return getters['getFeeBalanceAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFeeBalanceAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryKeysignVoteData({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryKeysignVoteData( key.index)).data
				
					
				commit('QUERY', { query: 'KeysignVoteData', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryKeysignVoteData', payload: { options: { all }, params: {...key},query }})
				return getters['getKeysignVoteData']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryKeysignVoteData API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryKeysignVoteDataAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryKeysignVoteDataAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryKeysignVoteDataAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'KeysignVoteDataAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryKeysignVoteDataAll', payload: { options: { all }, params: {...key},query }})
				return getters['getKeysignVoteDataAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryKeysignVoteDataAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryObserveVote({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryObserveVote( key.index)).data
				
					
				commit('QUERY', { query: 'ObserveVote', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryObserveVote', payload: { options: { all }, params: {...key},query }})
				return getters['getObserveVote']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryObserveVote API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryObserveVoteAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryObserveVoteAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryObserveVoteAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ObserveVoteAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryObserveVoteAll', payload: { options: { all }, params: {...key},query }})
				return getters['getObserveVoteAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryObserveVoteAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPoolBalance({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryPoolBalance( key.index)).data
				
					
				commit('QUERY', { query: 'PoolBalance', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPoolBalance', payload: { options: { all }, params: {...key},query }})
				return getters['getPoolBalance']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPoolBalance API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPoolBalanceAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryPoolBalanceAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryPoolBalanceAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PoolBalanceAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPoolBalanceAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPoolBalanceAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPoolBalanceAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTransactionData({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryTransactionData( key.index)).data
				
					
				commit('QUERY', { query: 'TransactionData', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTransactionData', payload: { options: { all }, params: {...key},query }})
				return getters['getTransactionData']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTransactionData API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTransactionDataAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryTransactionDataAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryTransactionDataAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TransactionDataAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTransactionDataAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTransactionDataAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTransactionDataAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPubkeys({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryPubkeys( key.index)).data
				
					
				commit('QUERY', { query: 'Pubkeys', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPubkeys', payload: { options: { all }, params: {...key},query }})
				return getters['getPubkeys']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPubkeys API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPubkeysAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryPubkeysAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryPubkeysAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PubkeysAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPubkeysAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPubkeysAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPubkeysAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryWhitelistedNode({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryWhitelistedNode( key.index)).data
				
					
				commit('QUERY', { query: 'WhitelistedNode', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryWhitelistedNode', payload: { options: { all }, params: {...key},query }})
				return getters['getWhitelistedNode']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryWhitelistedNode API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryWhitelistedNodeAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.HumansdotaiHumansHumans.query.queryWhitelistedNodeAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.HumansdotaiHumansHumans.query.queryWhitelistedNodeAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'WhitelistedNodeAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryWhitelistedNodeAll', payload: { options: { all }, params: {...key},query }})
				return getters['getWhitelistedNodeAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryWhitelistedNodeAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgObservationVote({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.HumansdotaiHumansHumans.tx.sendMsgObservationVote({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgObservationVote:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgObservationVote:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgApproveTransaction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.HumansdotaiHumansHumans.tx.sendMsgApproveTransaction({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveTransaction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgApproveTransaction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateBalance({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.HumansdotaiHumansHumans.tx.sendMsgUpdateBalance({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateBalance:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateBalance:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgKeysignVote({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.HumansdotaiHumansHumans.tx.sendMsgKeysignVote({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgKeysignVote:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgKeysignVote:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRequestTransaction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.HumansdotaiHumansHumans.tx.sendMsgRequestTransaction({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRequestTransaction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRequestTransaction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgTranfserPoolcoin({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.HumansdotaiHumansHumans.tx.sendMsgTranfserPoolcoin({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTranfserPoolcoin:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgTranfserPoolcoin:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgObservationVote({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.HumansdotaiHumansHumans.tx.msgObservationVote({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgObservationVote:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgObservationVote:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgApproveTransaction({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.HumansdotaiHumansHumans.tx.msgApproveTransaction({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveTransaction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgApproveTransaction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateBalance({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.HumansdotaiHumansHumans.tx.msgUpdateBalance({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateBalance:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateBalance:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgKeysignVote({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.HumansdotaiHumansHumans.tx.msgKeysignVote({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgKeysignVote:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgKeysignVote:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRequestTransaction({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.HumansdotaiHumansHumans.tx.msgRequestTransaction({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRequestTransaction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRequestTransaction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgTranfserPoolcoin({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.HumansdotaiHumansHumans.tx.msgTranfserPoolcoin({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTranfserPoolcoin:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgTranfserPoolcoin:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
