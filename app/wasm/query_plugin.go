package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Fury-Labs/fury/app/wasm/bindings"
)

func CustomQuerier(queryPlugin *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var furyQuery bindings.FuryQuery
		if err := json.Unmarshal(request, &furyQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "app query")
		}
		if furyQuery.AppData != nil {
			appID := furyQuery.AppData.AppID
			MinGovDeposit, GovTimeInSeconds, assetID, _ := queryPlugin.GetAppInfo(ctx, appID)
			res := bindings.AppDataResponse{
				MinGovDeposit:    MinGovDeposit.String(),
				GovTimeInSeconds: GovTimeInSeconds,
				AssetID:          assetID,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "App data query response")
			}
			return bz, nil
		} else if furyQuery.AssetData != nil {
			assetID := furyQuery.AssetData.AssetID
			denom, _ := queryPlugin.GetAssetInfo(ctx, assetID)
			res := bindings.AssetDataResponse{
				Denom: denom,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "App data query response")
			}
			return bz, nil
		} else if furyQuery.MintedToken != nil {
			appID := furyQuery.MintedToken.AppID
			assetID := furyQuery.MintedToken.AssetID
			MintedToken, _ := queryPlugin.GetTokenMint(ctx, appID, assetID)
			res := bindings.MintedTokenResponse{
				MintedTokens: MintedToken,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "tokenMint query response")
			}
			return bz, nil
		} else if furyQuery.RemoveWhiteListAssetLocker != nil {
			appID := furyQuery.RemoveWhiteListAssetLocker.AppID
			assetID := furyQuery.RemoveWhiteListAssetLocker.AssetIDs

			found, errormsg := queryPlugin.GetRemoveWhitelistAppIDLockerRewardsCheck(ctx, appID, assetID)
			res := bindings.RemoveWhiteListAssetResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhiteListAssetLocker query response")
			}
			return bz, nil
		} else if furyQuery.WhitelistAppIDLockerRewards != nil {
			appID := furyQuery.WhitelistAppIDLockerRewards.AppID
			assetID := furyQuery.WhitelistAppIDLockerRewards.AssetID

			found, errormsg := queryPlugin.GetWhitelistAppIDLockerRewardsCheck(ctx, appID, assetID)
			res := bindings.WhitelistAppIDLockerRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhitelistAppIdLockerRewards query response")
			}
			return bz, nil
		} else if furyQuery.WhitelistAppIDVaultInterest != nil {
			appID := furyQuery.WhitelistAppIDVaultInterest.AppID

			found, errormsg := queryPlugin.GetWhitelistAppIDVaultInterestCheck(ctx, appID)
			res := bindings.WhitelistAppIDLockerRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhitelistAppIdVaultInterest query response")
			}
			return bz, nil
		} else if furyQuery.ExternalLockerRewards != nil {
			appID := furyQuery.ExternalLockerRewards.AppID
			assetID := furyQuery.ExternalLockerRewards.AssetID

			found, errormsg := queryPlugin.GetExternalLockerRewardsCheck(ctx, appID, assetID)
			res := bindings.WhitelistAppIDLockerRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "GetExternalLockerRewardsCheck query response")
			}
			return bz, nil
		} else if furyQuery.ExternalVaultRewards != nil {
			appID := furyQuery.ExternalVaultRewards.AppID
			assetID := furyQuery.ExternalVaultRewards.AssetID

			found, errormsg := queryPlugin.GetExternalVaultRewardsCheck(ctx, appID, assetID)
			res := bindings.ExternalVaultRewardsResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExternalVaultRewards query response")
			}
			return bz, nil
		} else if furyQuery.CollectorLookupTableQuery != nil {
			appID := furyQuery.CollectorLookupTableQuery.AppID
			collectorAssetID := furyQuery.CollectorLookupTableQuery.CollectorAssetID
			secondaryAssetID := furyQuery.CollectorLookupTableQuery.SecondaryAssetID
			found, errormsg := queryPlugin.CollectorLookupTableQueryCheck(ctx, appID, collectorAssetID, secondaryAssetID)
			res := bindings.CollectorLookupTableQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExternalVaultRewards query response")
			}
			return bz, nil
		} else if furyQuery.ExtendedPairsVaultRecordsQuery != nil {
			appID := furyQuery.ExtendedPairsVaultRecordsQuery.AppID
			pairID := furyQuery.ExtendedPairsVaultRecordsQuery.PairID
			StabilityFee := furyQuery.ExtendedPairsVaultRecordsQuery.StabilityFee
			ClosingFee := furyQuery.ExtendedPairsVaultRecordsQuery.ClosingFee
			DrawDownFee := furyQuery.ExtendedPairsVaultRecordsQuery.DrawDownFee
			DebtCeiling := furyQuery.ExtendedPairsVaultRecordsQuery.DebtCeiling
			DebtFloor := furyQuery.ExtendedPairsVaultRecordsQuery.DebtFloor
			PairName := furyQuery.ExtendedPairsVaultRecordsQuery.PairName

			found, errorMsg := queryPlugin.ExtendedPairsVaultRecordsQueryCheck(ctx, appID, pairID, StabilityFee, ClosingFee, DrawDownFee, DebtCeiling, DebtFloor, PairName)
			res := bindings.ExtendedPairsVaultRecordsQueryResponse{
				Found: found,
				Err:   errorMsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExternalVaultRewards query response")
			}
			return bz, nil
		} else if furyQuery.AuctionMappingForAppQuery != nil {
			appID := furyQuery.AuctionMappingForAppQuery.AppID
			found, errormsg := queryPlugin.AuctionMappingForAppQueryCheck(ctx, appID)
			res := bindings.AuctionMappingForAppQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "AuctionMappingForAppQuery query response")
			}
			return bz, nil
		} else if furyQuery.WhiteListedAssetQuery != nil {
			appID := furyQuery.WhiteListedAssetQuery.AppID
			assetID := furyQuery.WhiteListedAssetQuery.AssetID
			found, errormsg := queryPlugin.WhiteListedAssetQueryCheck(ctx, appID, assetID)
			res := bindings.WhiteListedAssetQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhiteListedAssetQueryCheck query response")
			}
			return bz, nil
		} else if furyQuery.UpdatePairsVaultQuery != nil {
			appID := furyQuery.UpdatePairsVaultQuery.AppID
			extPairID := furyQuery.UpdatePairsVaultQuery.ExtPairID
			found, errormsg := queryPlugin.UpdatePairsVaultQueryCheck(ctx, appID, extPairID)
			res := bindings.UpdatePairsVaultQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "UpdatePairsVaultQuery query response")
			}
			return bz, nil
		} else if furyQuery.UpdateCollectorLookupTableQuery != nil {
			appID := furyQuery.UpdateCollectorLookupTableQuery.AppID
			assetID := furyQuery.UpdateCollectorLookupTableQuery.AssetID
			found, errormsg := queryPlugin.UpdateCollectorLookupTableQueryCheck(ctx, appID, assetID)
			res := bindings.UpdateCollectorLookupTableQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "UpdatePairsVaultQuery query response")
			}
			return bz, nil
		} else if furyQuery.RemoveWhitelistAppIDVaultInterestQuery != nil {
			appID := furyQuery.RemoveWhitelistAppIDVaultInterestQuery.AppID
			found, errormsg := queryPlugin.WasmRemoveWhitelistAppIDVaultInterestQueryCheck(ctx, appID)
			res := bindings.RemoveWhitelistAppIDVaultInterestQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhitelistAppIdVaultInterestQuery query response")
			}
			return bz, nil
		} else if furyQuery.RemoveWhitelistAssetLockerQuery != nil {
			appID := furyQuery.RemoveWhitelistAssetLockerQuery.AppID
			assetID := furyQuery.RemoveWhitelistAssetLockerQuery.AssetID

			found, errormsg := queryPlugin.WasmRemoveWhitelistAssetLockerQueryCheck(ctx, appID, assetID)
			res := bindings.RemoveWhitelistAssetLockerQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhitelistAssetLockerQuery query response")
			}
			return bz, nil
		} else if furyQuery.WhitelistAppIDLiquidationQuery != nil {
			AppID := furyQuery.WhitelistAppIDLiquidationQuery.AppID

			found, errormsg := queryPlugin.WasmWhitelistAppIDLiquidationQueryCheck(ctx, AppID)
			res := bindings.WhitelistAppIDLiquidationQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "WhitelistAppIDLiquidationQuery query response")
			}
			return bz, nil
		} else if furyQuery.RemoveWhitelistAppIDLiquidationQuery != nil {
			AppID := furyQuery.RemoveWhitelistAppIDLiquidationQuery.AppID

			found, errormsg := queryPlugin.WasmRemoveWhitelistAppIDLiquidationQueryCheck(ctx, AppID)
			res := bindings.RemoveWhitelistAppIDLiquidationQueryResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "RemoveWhitelistAppIDLiquidationQuery query response")
			}
			return bz, nil
		} else if furyQuery.AddESMTriggerParamsForAppQuery != nil {
			AppID := furyQuery.AddESMTriggerParamsForAppQuery.AppID

			found, errormsg := queryPlugin.WasmAddESMTriggerParamsQueryCheck(ctx, AppID)
			res := bindings.AddESMTriggerParamsForAppResponse{
				Found: found,
				Err:   errormsg,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "AddESMTriggerParamsForAppResponse query response")
			}
			return bz, nil
		} else if furyQuery.ExtendedPairByApp != nil {
			AppID := furyQuery.ExtendedPairByApp.AppID

			extendedPair, _ := queryPlugin.WasmExtendedPairByApp(ctx, AppID)
			res := bindings.ExtendedPairByAppResponse{
				ExtendedPair: extendedPair,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "ExtendedPairByAppResponse query response")
			}
			return bz, nil
		} else if furyQuery.CheckSurplusReward != nil {
			AppID := furyQuery.CheckSurplusReward.AppID
			AssetID := furyQuery.CheckSurplusReward.AssetID
			amount := queryPlugin.WasmCheckSurplusReward(ctx, AppID, AssetID)
			res := bindings.CheckSurplusRewardResponse{
				Amount: amount,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "CheckSurplusRewardResponse query response")
			}
			return bz, nil
		} else if furyQuery.CheckWhitelistedAsset != nil {
			Denom := furyQuery.CheckWhitelistedAsset.Denom

			found := queryPlugin.WasmCheckWhitelistedAsset(ctx, Denom)
			res := bindings.CheckWhitelistedAssetResponse{
				Found: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "CheckWhitelistedAssetResponse query response")
			}
			return bz, nil
		} else if furyQuery.CheckVaultCreated != nil {
			Address := furyQuery.CheckVaultCreated.Address
			AppID := furyQuery.CheckVaultCreated.AppID
			found := queryPlugin.WasmCheckVaultCreated(ctx, Address, AppID)
			res := bindings.VaultCreatedResponse{
				IsCompleted: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "VaultCreatedResponse query response")
			}
			return bz, nil
		} else if furyQuery.CheckBorrowed != nil {
			AssetID := furyQuery.CheckBorrowed.AssetID
			Address := furyQuery.CheckBorrowed.Address
			found := queryPlugin.WasmCheckBorrowed(ctx, AssetID, Address)
			res := bindings.BorrowedResponse{
				IsCompleted: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "BorrowedResponse query response")
			}
			return bz, nil
		} else if furyQuery.CheckLiquidityProvided != nil {
			AppID := furyQuery.CheckLiquidityProvided.AppID
			PoolID := furyQuery.CheckLiquidityProvided.PoolID
			Address := furyQuery.CheckLiquidityProvided.Address
			found := queryPlugin.WasmCheckLiquidityProvided(ctx, AppID, PoolID, Address)
			res := bindings.LiquidityProvidedResponse{
				IsCompleted: found,
			}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "LiquidityProvidedResponse query response")
			}
			return bz, nil
		}
		return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown App Data query variant"}
	}
}
