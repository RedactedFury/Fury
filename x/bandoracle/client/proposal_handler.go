package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/redactedfury/fury/x/bandoracle/client/cli"
	"github.com/redactedfury/fury/x/bandoracle/client/rest"
)

var AddFetchPriceHandler = govclient.NewProposalHandler(cli.NewCmdSubmitFetchPriceProposal, rest.SubmitFetchPriceProposalRESTHandler)
