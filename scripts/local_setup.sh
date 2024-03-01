#!/bin/bash
set -e
set -x

BIN=./build/sedad
CONFIG_PATH=$HOME/.sedad/config

function add_key_and_account() {
    local name=$1
    local amount=$2
    $BIN keys add $name --keyring-backend test
    $BIN add-genesis-account $name $amount --keyring-backend test
}

#
# Local Single-node Setup
#
# NOTE: Run this script from project root.
#

# build the binary
make build

# reset the chain
$BIN tendermint unsafe-reset-all
rm -rf ~/.sedad || true

# configure sedad
$BIN config set client chain-id sedachain

# initialize the chain
$BIN init node0 --default-denom aseda

cat $HOME/.sedad/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="30s"' > $HOME/.sedad/config/tmp_genesis.json && mv $HOME/.sedad/config/tmp_genesis.json $HOME/.sedad/config/genesis.json
cat $HOME/.sedad/config/genesis.json | jq '.app_state["gov"]["params"]["voting_period"]="30s"' > $HOME/.sedad/config/tmp_genesis.json && mv $HOME/.sedad/config/tmp_genesis.json $HOME/.sedad/config/genesis.json
cat $HOME/.sedad/config/genesis.json | jq '.app_state["gov"]["params"]["expedited_voting_period"]="15s"' > $HOME/.sedad/config/tmp_genesis.json && mv $HOME/.sedad/config/tmp_genesis.json $HOME/.sedad/config/genesis.json
cat $HOME/.sedad/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="100000000"' > $HOME/.sedad/config/tmp_genesis.json && mv $HOME/.sedad/config/tmp_genesis.json $HOME/.sedad/config/genesis.json

# update genesis
add_key_and_account "satoshi" "100000000000000000seda"
add_key_and_account "acc1" "100000000000000000seda"

# create a default validator
$BIN gentx satoshi 10000000000000000seda --keyring-backend test --from satoshi

# collect genesis txns
$BIN collect-gentxs

# start the chain
$BIN start --log_level debug || echo "Failed to start the chain"