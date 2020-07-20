#!/usr/bin/env bash

# TODO: No longer a valid demo

echo "Broadcasting..."
TX='{"mode":"block","tx":"0x7B227061796C6F6164223A5B7B2274797065223A226469642F416464446964222C2276616C7565223A7B22646964446F63223A7B22646964223A226469643A69786F3A55347453707A7A763931484871575731596D466B484A222C227075624B6579223A22466B65447565356974383274616568654D7072646150726374664B3344655656394E6E455059446777775247222C2263726564656E7469616C73223A5B5D7D7D7D5D2C22666565223A7B22616D6F756E74223A5B7B22616D6F756E74223A2235303030222C2264656E6F6D223A227569786F227D5D2C22676173223A22323030303030227D2C227369676E617475726573223A5B7B227369676E617475726556616C7565223A22464D392B724653452B565633687033397753384A65666B3574657A4B65466E393732776247666D5651687846336B2B6D666C5A7456443648692F7247416F4965784A35417564595A66797068594B324B3753567942673D3D222C2263726561746564223A22323032302D30372D30365431323A30343A32352E3736385A227D5D7D"}'
curl -X POST localhost:1317/txs --data-binary "$TX"

echo ""

echo "Now we can query the DID..."
curl -X GET localhost:1317/did/did:ixo:U4tSpzzv91HHqWW1YmFkHJ

# Decoded transaction (what we're actually broadcasting):
# {
#   "payload": [
#     {
#       "type": "did/AddDid",
#       "value": {
#         "didDoc": {
#           "did": "did:ixo:U4tSpzzv91HHqWW1YmFkHJ",
#           "pubKey": "FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
#           "credentials": []
#         }
#       }
#     }
#   ],
#   "fee": {
#     "amount": [
#       {
#         "amount": "5000",
#         "denom": "uixo"
#       }
#     ],
#     "gas": "200000"
#   },
#   "signatures": [
#     {
#       "signatureValue": "FM9+rFSE+VV3hp39wS8Jefk5tezKeFn972wbGfmVQhxF3k+mflZtVD6Hi/rGAoIexJ5AudYZfyphYK2K7SVyBg==",
#       "created": "2020-07-06T12:04:25.768Z"
#     }
#   ]
# }