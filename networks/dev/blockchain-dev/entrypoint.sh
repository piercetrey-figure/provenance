#!/usr/bin/env bash

##
## Input parameters
##
export PIO_HOME="/provenance/nodedev"
BINARY=/usr/bin/${BINARY:-provenanced}
LOG=${LOG:-provenance.log}

##
## Generate genesis if not already present
##
if ! [ -f ${PIO_HOME}/config/genesis.json ]; then
  "${BINARY}" -t --home "${PIO_HOME}" init --chain-id=chain-dev testing ; \
  hash_supply=0
  num_accounts=0
  for f in /mnemonics/*.txt
    do
      key_name=$(basename $f .txt)
      echo "Adding account $key_name from mnemonic file $f with 100000000000000000000nhash"
      "${BINARY}" -t --home "${PIO_HOME}" keys add $key_name --recover --keyring-backend test < $f  
      "${BINARY}" -t --home "${PIO_HOME}" add-genesis-account $key_name 100000000000000000000nhash --keyring-backend test
      let num_accounts=num_accounts+1
  done
  "${BINARY}" -t --home "${PIO_HOME}" add-genesis-root-name validator pio --keyring-backend test
  "${BINARY}" -t --home "${PIO_HOME}" add-genesis-root-name validator pb --restrict=false --keyring-backend test  
  "${BINARY}" -t --home "${PIO_HOME}" add-genesis-root-name validator io --restrict --keyring-backend test 
  "${BINARY}" -t --home "${PIO_HOME}" add-genesis-root-name validator provenance --keyring-backend test 
  "${BINARY}" -t --home "${PIO_HOME}" add-genesis-marker ${num_accounts}00000000000000000000nhash --manager validator --access mint,burn,admin,withdraw,deposit --activate --keyring-backend test 
  "${BINARY}" -t --home "${PIO_HOME}" gentx validator 100000000000000nhash --keyring-backend test --chain-id=chain-dev
  "${BINARY}" -t --home "${PIO_HOME}" collect-gentxs 
  "${BINARY}" -t --home "${PIO_HOME}" config set rpc.laddr tcp://0.0.0.0:26657
  "${BINARY}" -t --home "${PIO_HOME}" config set api.enable true
  "${BINARY}" -t --home "${PIO_HOME}" config set api.swagger true
  "${BINARY}" -t config set minimum-gas-prices "0vspn"
  
fi

jq '.consensus_params.block.max_gas = "1000000000" | .conssensus_params.block.max_bytes = "2200020096" | .app_state.msgfees.params.floor_gas_price.amount = "0"' /provenance/nodedev/config/genesis.json > /provenance/nodedev/config/genesis.json.tmp && mv /provenance/nodedev/config/genesis.json.tmp /provenance/nodedev/config/genesis.json
# cat /provenance/nodedev/config/genesis.json

##
## Run binary with all parameters
##
if [ -d "$(dirname "${PIO_HOME}"/"${LOG}")" ]; then
 "${BINARY}" -t --home "${PIO_HOME}" "$@" | tee "${PIO_HOME}/${LOG}"
else
 "${BINARY}" -t --home "${PIO_HOME}" "$@" 
fi

