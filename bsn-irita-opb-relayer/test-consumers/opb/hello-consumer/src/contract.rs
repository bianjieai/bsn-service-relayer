use cosmwasm_std::{attr, Binary, CosmosMsg, Deps, DepsMut, Env, HandleResponse, HumanAddr, InitResponse, MessageInfo, StdResult, WasmMsg, to_binary};
use json::object;
use base64::encode;

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg,QueryMsg};
use crate::state::{State,config};

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
pub fn init(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InitMsg,
) -> Result<InitResponse, ContractError> {
    let state = State {
        owner: deps.api.canonical_address(&info.sender)?,
        core_contract: msg.core_contract,
        target_contract: msg.target_contract,
    };
    config(deps.storage).save(&state)?;

    Ok(InitResponse::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
pub fn handle(
    deps: DepsMut,
    env: Env,
    _info: MessageInfo,
    msg: HandleMsg,
) -> Result<HandleResponse, ContractError> {
    match msg {
        HandleMsg::Hello{words} => try_hello(deps,words,env.contract.address),
        HandleMsg::CallBack{request_id, words} => call_back(deps, request_id, words),
    }
}

pub fn try_hello(deps: DepsMut, words: String, self_address: HumanAddr) -> Result<HandleResponse, ContractError> {
    let state = config(deps.storage).load()?;
    let endpoint_address = state.target_contract.as_str();
    let self_address_str = self_address.as_str();
    let endpoint_info = object!{"dest_chain_id"=>"1","dest_chain_type"=>"opb","endpoint_address"=>endpoint_address,"endpoint_type"=>"contract"};
    let call_data = object!{"hello"=>object! {"words"=>words}};
    let call_data_base64 = encode(call_data.to_string().as_bytes());
    let msg = object!{"send_request"=> object!{"endpoint_info"=>endpoint_info.to_string(),"method"=>"hello","call_data"=>call_data_base64,"callback_address"=>self_address_str,"callback_function"=>"call_back"}};

    let msgs = vec![CosmosMsg::Wasm(WasmMsg::Execute {
        contract_addr: state.core_contract,
        msg:  Binary::from(msg.to_string().as_bytes()),
        send: vec![],
    })];

    Ok(HandleResponse {
        messages: msgs,
        data: None,
        attributes: vec![],
    })
}

pub fn call_back(_deps: DepsMut, request_id: String, words: String) -> Result<HandleResponse, ContractError> {
    Ok(HandleResponse {
        messages: vec![],
        data: None,
        attributes: vec![attr("request_id",request_id),attr("words",words)],
    })
}

pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    to_binary("no query function")
}