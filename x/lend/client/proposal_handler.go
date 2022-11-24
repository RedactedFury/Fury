package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/Fury-Labs/fury/x/lend/client/cli"
	"github.com/Fury-Labs/fury/x/lend/client/rest"
)

var (
	AddLendPairsHandler        = govclient.NewProposalHandler(cli.CmdAddNewLendPairsProposal, rest.AddNewPairsProposalRESTHandler)
	AddPoolHandler             = govclient.NewProposalHandler(cli.CmdAddPoolProposal, rest.AddPoolProposalRESTHandler)
	AddAssetToPairHandler      = govclient.NewProposalHandler(cli.CmdAddAssetToPairProposal, rest.AddAssetToPairProposalRESTHandler)
	AddAssetRatesParamsHandler = govclient.NewProposalHandler(cli.CmdAddNewAssetRatesParamsProposal, rest.AddNewAssetRatesParamsProposalRESTHandler)
	AddAuctionParamsHandler    = govclient.NewProposalHandler(cli.CmdAddNewAuctionParamsProposal, rest.AddNewAuctionParamsProposalRESTHandler)
)
