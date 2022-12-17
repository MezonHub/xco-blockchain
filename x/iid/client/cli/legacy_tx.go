package cli

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	libxco "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	"github.com/furylockerroom/xco-blockchain/lib/legacydid"
	"github.com/furylockerroom/xco-blockchain/x/iid/types"
	"github.com/spf13/cobra"
)

func NewCreateIidDocumentFormLegacyDidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-iid-from-legacy-did [did]",
		Short:   "create decentralized did (did) document from legacy did",
		Example: "creates a did document for users",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			xcodid, err := legacydid.Unmarshalxcodid(args[0])
			if err != nil {
				return err
			}

			// chaincmd.Flags().GetString(flags.FlagChainID)
			// if err != nil {
			// 	return err
			// }
			// did
			did := types.DID(xcodid.Did)
			// verification
			// signer := clientCtx.GetFromAddress()
			// pubkey

			pubKey := xcodid.VerifyKey

			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			// understand the vmType

			auth := types.NewVerification(
				types.NewVerificationMethod(
					xcodid.Did,
					did,
					types.NewPublicKeyMultibase(base58.Decode(pubKey), types.DIDVMethodTypeEd25519VerificationKey2018),
				),
				[]string{types.Authentication},
				nil,
			)
			// create the message
			msg := types.NewMsgCreateIidDocument(
				did.String(),
				types.Verifications{auth},
				types.Services{},
				types.AccordedRights{},
				types.LinkedResources{},
				types.LinkedEntities{},
				xcodid.Address().String(),
				types.Contexts{},
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// execute
			return libxco.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
