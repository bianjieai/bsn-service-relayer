use std::str;

use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use cw_controllers::{Admin};

use cosmwasm_std::HumanAddr;

use cw_storage_plus::{Map,Item};

// REQUESTS stores the requests whether the response has been returned: <request_id, responded>
pub const REQUESTS: Map<&[u8], bool> = Map::new("people");

// CALLBACKS stores the requests whether the response has been returned: <request_id, callback>
pub const CALLBACKS: Map<&str, Callback> = Map::new("callbacks");

pub const RELAYER: Admin = Admin::new("relayer");

pub const REQUESTSEQUENCE: Item<u64> = Item::new("request_sequence");

pub const CHAINID: Item<String> = Item::new("chain_id");

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
#[serde(rename_all = "snake_case")]
pub struct Callback{
    pub address: HumanAddr,
    pub method: String,
}
