#!/usr/bin/env bash

echo "Exporting app state to genesis file..."
xcod export >genesis.json

echo "Fixing genesis file..."
sed -i 's/"genutil":null/"genutil":{"gentxs":null}/g' genesis.json
# https://github.com/cosmos/cosmos-sdk/issues/5086

echo "Backing up existing genesis file..."
cp "$HOME"/.xcod/config/genesis.json "$HOME"/.xcod/config/genesis.json.backup

echo "Moving new genesis file to $HOME/.xcod/config/genesis.json..."
mv genesis.json "$HOME"/.xcod/config/genesis.json

xcod unsafe-reset-all
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
