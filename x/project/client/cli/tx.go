package cli

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	xcotypes "github.com/furylockerroom/xco-blockchain/lib/xco"
	didtypes "github.com/furylockerroom/xco-blockchain/lib/legacydid"
	iidtypes "github.com/furylockerroom/xco-blockchain/x/iid/types"
	"github.com/furylockerroom/xco-blockchain/x/project/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	projectTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "project transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	projectTxCmd.AddCommand(
		NewCmdCreateProject(),
		NewCmdCreateAgent(),
		NewCmdUpdateProjectStatus(),
		NewCmdUpdateAgent(),
		NewCmdCreateClaim(),
		NewCmdCreateEvaluation(),
		NewCmdWithdrawFunds(),
		NewCmdUpdateProjectDoc(),
	)

	return projectTxCmd
}

func NewCmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [sender-did] [project-iid-json] [xco-did]",
		Short: "Create a new ProjectDoc signed by the xcodid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDataStr := args[1]
			xcodidStr := args[2]

			xcodid, err := didtypes.Unmarshalxcodid(xcodidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgCreateProject(
				iidtypes.DIDFragment(senderDid), json.RawMessage(projectDataStr), xcodid.Did, xcodid.VerifyKey, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			res, err := xcotypes.SignAndBroadcastTxFromStdSignMsg(clientCtx, msg, xcodid, cmd.Flags())
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateProjectStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-status [sender-did] [status] [xco-did]",
		Short: "Update the status of a project signed by the xcodid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			status := args[1]
			xcodid, err := didtypes.Unmarshalxcodid(args[2])
			if err != nil {
				return err
			}

			projectStatus := types.ProjectStatus(status)
			//if projectStatus != types.CreatedProject &&
			//	projectStatus != types.PendingStatus &&
			//	projectStatus != types.FundedStatus &&
			//	projectStatus != types.StartedStatus &&
			//	projectStatus != types.StoppedStatus &&
			//	projectStatus != types.PaidoutStatus {
			//	return errors.New("The status must be one of 'CREATED', " +
			//		"'PENDING', 'FUNDED', 'STARTED', 'STOPPED' or 'PAIDOUT'")
			//}

			updateProjectStatusDoc := types.NewUpdateProjectStatusDoc(
				projectStatus, "")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgUpdateProjectStatus(iidtypes.DIDFragment(senderDid), updateProjectStatusDoc, xcodid.Did, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-agent [tx-hash] [sender-did] [agent-did] " +
			"[role] [project-did]",
		Short: "Create a new agent on a project signed by the xcodid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			role := args[3]
			if role != "SA" && role != "EA" && role != "IA" {
				return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
			}

			createAgentDoc := types.NewCreateAgentDoc(agentDid, role)

			xcodid, err := didtypes.Unmarshalxcodid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgCreateAgent(txHash, iidtypes.DIDFragment(senderDid), createAgentDoc, xcodid.Did, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-agent [tx-hash] [sender-did] [agent-did] " +
			"[status] [xco-did]",
		Short: "Update the status of an agent on a project signed by the xcodid of the project",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			agentRole := args[4]
			agentStatus := types.AgentStatus(args[3])
			if agentStatus != types.PendingAgent && agentStatus != types.ApprovedAgent && agentStatus != types.RevokedAgent {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
			}

			updateAgentDoc := types.NewUpdateAgentDoc(
				agentDid, agentStatus, agentRole)

			xcodid, err := didtypes.Unmarshalxcodid(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgUpdateAgent(txHash, iidtypes.DIDFragment(senderDid), updateAgentDoc, xcodid.Did, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-claim [tx-hash] [sender-did] [claim-id] [claim-template-id] [xco-did]",
		Short: "Create a new claim on a project signed by the xcodid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimTemplateId := args[3]
			createClaimDoc := types.NewCreateClaimDoc(claimId, claimTemplateId)

			xcodid, err := didtypes.Unmarshalxcodid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgCreateClaim(txHash, iidtypes.DIDFragment(senderDid), createClaimDoc, xcodid.Did, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateEvaluation() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-evaluation [tx-hash] [sender-did] [claim-id] " +
			"[status] [xco-did]",
		Short: "Create a new claim evaluation on a project signed by the xcodid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimStatus := types.ClaimStatus(args[3])
			if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			}

			createEvaluationDoc := types.NewCreateEvaluationDoc(
				claimId, claimStatus)

			xcodid, err := didtypes.Unmarshalxcodid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgCreateEvaluation(txHash, iidtypes.DIDFragment(senderDid), createEvaluationDoc, xcodid.Did, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdWithdrawFunds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-funds [sender-did] [iid]",
		Short: "Withdraw funds.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			xcodid, err := didtypes.Unmarshalxcodid(args[0])
			if err != nil {
				return err
			}

			var data types.WithdrawFundsDoc
			err = json.Unmarshal([]byte(args[1]), &data)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgWithdrawFunds(iidtypes.DIDFragment(xcodid.Did), data, xcodid.Address().String())
			msg.SenderAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateProjectDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-doc [sender-did] [project-iid-json] [xco-did]",
		Short: "Update a project's iid signed by the xcodid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDataStr := args[1]
			xcodid, err := didtypes.Unmarshalxcodid(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(xcodid.Address())

			msg := types.NewMsgUpdateProjectDoc(iidtypes.DIDFragment(senderDid), json.RawMessage(projectDataStr), xcodid.Did, xcodid.Address().String())
			msg.ProjectAddress = xcodid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return xcotypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), xcodid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
