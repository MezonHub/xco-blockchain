package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/xcohub/xco-blockchain/x/token/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	tokenQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "token query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tokenQueryCmd.AddCommand(
		GetCmdTokenDocs(),
	)

	return tokenQueryCmd
}

func GetCmdTokenDocs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-token-doc",
		Short: "Query TokenDocs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			minterDid := args[0]

			res, err := queryClient.TokenList(context.Background(), &types.QueryTokenListRequest{MinterDid: minterDid})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// func GetCmdProjectAccounts() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "get-project-accounts [did]",
// 	Short: "Get a Project accounts of a Project by Did",
// 	Args:  cobra.ExactArgs(1),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		clientCtx, err := client.GetClientQueryContext(cmd)
// 		if err != nil {
// 			return err
// 		}

// 		projectDid := args[0]

// 		queryClient := types.NewQueryClient(clientCtx)

// 		res, err := queryClient.ProjectAccounts(context.Background(), &types.QueryProjectAccountsRequest{ProjectDid: projectDid})
// 		if err != nil {
// 			return err
// 		}

// 		if len(res.GetAccountMap().Map) == 0 {
// 			return errors.New("project does not exist")
// 		}

// 		return clientCtx.PrintProto(res)
// 	},
// }

// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// func GetCmdProjectTxs() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "get-project-txs [project-did]",
// 	Short: "Get a Project txs for a projectDid",
// 	Args:  cobra.ExactArgs(1),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		clientCtx, err := client.GetClientQueryContext(cmd)
// 		if err != nil {
// 			return err
// 		}

// 		projectDid := args[0]

// 		queryClient := types.NewQueryClient(clientCtx)

// 		res, err := queryClient.ProjectTx(context.Background(), &types.QueryProjectTxRequest{ProjectDid: projectDid})
// 		if err != nil {
// 			return err
// 		}

// 		if len(res.GetTxs().DocsList) == 0 {
// 			return errors.New("project does not have any transactions")
// 		}

// 		return clientCtx.PrintProto(res)
// 	},
// }

// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// func GetParamsRequestHandler() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "params",
// 	Short: "Query params",
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		clientCtx, err := client.GetClientQueryContext(cmd)
// 		if err != nil {
// 			return err
// 		}

// 		queryClient := types.NewQueryClient(clientCtx)

// 		res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
// 		if err != nil {
// 			return err
// 		}

// 		return clientCtx.PrintProto(res)
// 	},
// }

// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }
