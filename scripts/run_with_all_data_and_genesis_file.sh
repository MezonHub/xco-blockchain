#!/usr/bin/env bash

PASSWORD="12345678"

xcod init local --chain-id pandora-4

echo "Backing up existing genesis file..."
cp "$HOME"/.xcod/config/genesis.json "$HOME"/.xcod/config/genesis.json.backup

echo "Copying new genesis file to $HOME/.xcod/config/genesis.json..."
cp genesis.json "$HOME"/.xcod/config/genesis.json

yes 'y' | xcod keys delete miguel --force
yes 'y' | xcod keys delete francesco --force
yes 'y' | xcod keys delete shaun --force
yes 'y' | xcod keys delete fee --force
yes 'y' | xcod keys delete fee2 --force
yes 'y' | xcod keys delete fee3 --force
yes 'y' | xcod keys delete fee4 --force
yes 'y' | xcod keys delete fee5 --force
yes 'y' | xcod keys delete reserveOut --force

yes $PASSWORD | xcod keys add miguel
yes $PASSWORD | xcod keys add francesco
yes $PASSWORD | xcod keys add shaun
yes $PASSWORD | xcod keys add fee
yes $PASSWORD | xcod keys add fee2
yes $PASSWORD | xcod keys add fee3
yes $PASSWORD | xcod keys add fee4
yes $PASSWORD | xcod keys add fee5
yes $PASSWORD | xcod keys add reserveOut

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | xcod add-genesis-account "$(xcod keys show miguel -a)" 1000000000000uxco,1000000000000res,1000000000000rez,1000000000000uxgbp
yes $PASSWORD | xcod add-genesis-account "$(xcod keys show francesco -a)" 1000000000000uxco,1000000000000res,1000000000000rez
yes $PASSWORD | xcod add-genesis-account "$(xcod keys show shaun -a)" 1000000000000uxco,1000000000000res,1000000000000rez

# Add pubkey-based genesis accounts
MIGUEL_ADDR="ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"    # address from did:ixo:4XJLBfGtWSGKSz4BeRxdun's pubkey
FRANCESCO_ADDR="ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y" # address from did:ixo:UKzkhVSHc3qEFva5EY2XHt's pubkey
SHAUN_ADDR="ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"     # address from did:ixo:U4tSpzzv91HHqWW1YmFkHJ's pubkey
yes $PASSWORD | xcod add-genesis-account "$MIGUEL_ADDR" 1000000000000uxco,1000000000000res,1000000000000rez
yes $PASSWORD | xcod add-genesis-account "$FRANCESCO_ADDR" 1000000000000uxco,1000000000000res,1000000000000rez
yes $PASSWORD | xcod add-genesis-account "$SHAUN_ADDR" 1000000000000uxco,1000000000000res,1000000000000rez

# Add ixo did
XCO_DID="did:ixo:U4tSpzzv91HHqWW1YmFkHJ"
FROM="\"ixo_did\": \"\""
TO="\"ixo_did\": \"$XCO_DID\""
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
STAKING_TOKEN="uxco"
FROM="\"bond_denom\": \"stake\""
TO="\"bond_denom\": \"$STAKING_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/genesis.json
FROM="\"mint_denom\": \"stake\""
TO="\"mint_denom\": \"$STAKING_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
FEE_TOKEN="uxco"
FROM="\"stake\""
TO="\"$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/genesis.json

# Set reserved bond tokens
RESERVED_BOND_TOKENS=""  # example: " \"abc\", \"def\", \"ghi\" "
FROM="\"reserved_bond_tokens\": \[\]"
TO="\"reserved_bond_tokens\": \[$RESERVED_BOND_TOKENS\]"
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/genesis.json

# Set min-gas-prices (using fee token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/app.toml

# TODO: config missing from new version (REF: https://github.com/cosmos/cosmos-sdk/issues/8529)
#xcod config chain-id pandora-4
#xcod config output json
#xcod config indent true
#xcod config trust-node true

xcod gentx miguel 1000000uxco --chain-id pandora-4

xcod collect-gentxs
xcod validate-genesis

# Enable REST API (assumed to be at line 104 of app.toml)
FROM="enable = false"
TO="enable = true"
sed -i "104s/$FROM/$TO/" "$HOME"/.xcod/config/app.toml

# Enable Swagger docs (assumed to be at line 107 of app.toml)
FROM="swagger = false"
TO="swagger = true"
sed -i "107s/$FROM/$TO/" "$HOME"/.xcod/config/app.toml

# Uncomment the below to broadcast node RPC endpoint
#FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
#TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
#sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/config.toml

# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$HOME"/.xcod/config/config.toml
#sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$HOME"/.xcod/config/config.toml

xcod start --pruning "nothing"
