use cosmwasm_std::{Binary,HumanAddr};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub struct InitMsg {
    pub admin: Option<HumanAddr>,
    pub source_chain_id: String,
}

#[derive(Serialize, Deserialize, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum HandleMsg {
    SendRequest { endpoint_info: String, method: String, call_data: Binary, callback_address: HumanAddr, callback_function: String},
    SetResponse { request_id: String,err_msg: String, output: String},
    SetRelayer {relayer: Option<HumanAddr>},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    // GetCount returns the current count as a json-encoded number
    GetReuqest {requst_id: String},
    GetSequence {},
    GetRelayer {},
}