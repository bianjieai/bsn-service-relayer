
use std::u64;
use json::object;
use sha2::{Sha256, Digest};

use cosmwasm_std::{
    attr,to_binary,HumanAddr, Binary, Deps, DepsMut, Env, HandleResponse, InitResponse, MessageInfo, StdResult,CosmosMsg, WasmMsg
};

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg, QueryMsg};
use crate::state::{CHAINID, RELAYER, REQUESTS, REQUESTSEQUENCE, CALLBACKS,Callback};

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
pub fn init(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InitMsg,
) -> Result<InitResponse, ContractError> {
    let mut relayer = msg.admin;

    if relayer.is_none(){
        relayer = Some(info.sender);
    }


    let sequence: u64 = 0;
    REQUESTSEQUENCE.save(deps.storage, &sequence)?;
    CHAINID.save(deps.storage,&msg.source_chain_id)?;
    RELAYER.set(deps,relayer)?;

    Ok(InitResponse::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
pub fn handle(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> Result<HandleResponse, ContractError> {
    match msg {
        HandleMsg::SendRequest {endpoint_info, method, call_data, callback_address, callback_function} 
        => send_request(deps,endpoint_info, method, call_data, callback_address, callback_function),
        HandleMsg::SetResponse { request_id,err_msg, output } => set_response(deps, request_id,err_msg, output),
        HandleMsg::SetRelayer{ relayer } => set_relayer(deps, relayer,info.sender),
    }
}

pub fn send_request(deps: DepsMut,endpoint_info: String, _method: String, call_data: Binary, callback_address: HumanAddr, call_back_function: String) -> Result<HandleResponse, ContractError> {
    let chain_id = CHAINID.load(deps.storage)?;
    let mut sequence = REQUESTSEQUENCE.load(deps.storage)?;

    let request_id_str = chain_id+&sequence.to_string();

    let mut hasher = Sha256::new();
    hasher.input(&request_id_str.into_bytes());
    let output = hasher.result();
    let request_id = format!("{:X}",output);

    sequence += 1;
    REQUESTSEQUENCE.save(deps.storage, &sequence)?;
    
    let call_back = Callback{
        address: callback_address,
        method: call_back_function,
    };
    CALLBACKS.save(deps.storage, &request_id, &call_back)?;

    Ok(HandleResponse {
        messages: vec![],
        data: None,
        attributes: vec![attr("request_id", request_id),attr("endpoint_info", endpoint_info),attr("method", _method),attr("callData", call_data)],
    })
}

pub fn set_response(deps: DepsMut,request_id: String, err_msg: String, output: String) -> Result<HandleResponse, ContractError> {
    let executed = REQUESTS.load(deps.storage, &request_id);
    if executed.is_ok() {
        Err(ContractError::Unauthorized{})
    }else{
        let mut result = err_msg;
        if output.len()>0 {
            result = output;
        }

        REQUESTS.save(deps.storage, &request_id, &true)?;

        let callback = CALLBACKS.load(deps.storage, &request_id)?;
        let result_obg = object!{ callback.method.as_str() => object!{"request_id"=>request_id,"words"=>result}};

        let msgs = vec![CosmosMsg::Wasm(WasmMsg::Execute {
            contract_addr: callback.address,
            msg: Binary::from(result_obg.to_string().as_bytes()),
            send: vec![],
        })];
        
        Ok(HandleResponse {
            messages: msgs,
            data: None,
            attributes: vec![],
        })
    }
}

pub fn set_relayer(deps: DepsMut, relayer: Option<HumanAddr>, caller:HumanAddr) -> Result<HandleResponse, ContractError> {
    RELAYER.assert_admin(deps.as_ref(), &caller);
    RELAYER.set(deps, relayer); 
    Ok(HandleResponse::default())
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::GetReuqest {requst_id} => to_binary(&query_request(deps,requst_id)?),
        QueryMsg::GetSequence {} => to_binary(&query_sequence(deps)?),
        QueryMsg::GetRelayer {} => to_binary( &RELAYER.query_admin(deps)?),
    }
}

fn query_request(deps: Deps,requst_id: String) -> StdResult<bool> {
    let state = REQUESTS.load(deps.storage,&requst_id)?;
    Ok(state)
}

fn query_sequence(deps: Deps) -> StdResult<u64> {
    let state = REQUESTSEQUENCE.load(deps.storage)?;
    Ok(state)
}