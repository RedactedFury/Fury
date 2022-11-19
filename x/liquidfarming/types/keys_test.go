package types_test

import (
	fmt "fmt"
	"testing"
	time "time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	utils "github.com/redactedfury/fury/types"
	"github.com/redactedfury/fury/x/liquidfarming/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetLastRewardsAuctionIdKey() {
	s.Require().Equal([]byte{0xe2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLastRewardsAuctionIdKey(0))
	s.Require().Equal([]byte{0xe2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetLastRewardsAuctionIdKey(9))
	s.Require().Equal([]byte{0xe2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetLastRewardsAuctionIdKey(10))
}

func (s *keysTestSuite) TestGetLiquidFarmKey() {
	s.Require().Equal([]byte{0xe3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidFarmKey(0))
	s.Require().Equal([]byte{0xe3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetLiquidFarmKey(9))
	s.Require().Equal([]byte{0xe3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetLiquidFarmKey(10))
}

func (s *keysTestSuite) TestGetCompundingRewardsKey() {
	s.Require().Equal([]byte{0xe4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetCompoundingRewardsKey(0))
	s.Require().Equal([]byte{0xe4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetCompoundingRewardsKey(9))
	s.Require().Equal([]byte{0xe4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetCompoundingRewardsKey(10))
}

func (s *keysTestSuite) TestGetRewardsAuctionKey() {
	testCases := []struct {
		poolId    uint64
		auctionId uint64
		expected  []byte
	}{
		{
			1,
			1,
			[]byte{0xe5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			1,
			5,
			[]byte{0xe5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			5,
			5,
			[]byte{0xe5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetRewardsAuctionKey(tc.auctionId, tc.poolId)
			s.Require().Equal(tc.expected, key)
		})
	}
}

func (s *keysTestSuite) TestGetBidKey() {
	testCases := []struct {
		poolId   uint64
		bidder   sdk.AccAddress
		expected []byte
	}{
		{
			1,
			sdk.AccAddress(crypto.AddressHash([]byte("bidder1"))),
			[]byte{0xe6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14,
				0x20, 0x5c, 0xa, 0x82, 0xa, 0xf1, 0xed, 0x98, 0x39, 0x6a,
				0x35, 0xfe, 0xe3, 0x5d, 0x5, 0x2c, 0xd7, 0x96, 0x5a, 0x37},
		},
		{
			1,
			sdk.AccAddress(crypto.AddressHash([]byte("bidder3"))),
			[]byte{0xe6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14,
				0xe, 0x99, 0x7b, 0x9b, 0x5c, 0xef, 0x81, 0x2f, 0xc6, 0x3f,
				0xb6, 0x8b, 0x27, 0x42, 0x8a, 0xab, 0x7a, 0x58, 0xbc, 0x5e},
		},
		{
			5,
			sdk.AccAddress(crypto.AddressHash([]byte("bidder22"))),
			[]byte{0xe6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x14,
				0x4c, 0xf1, 0xbd, 0x90, 0x1, 0x70, 0x78, 0xfb, 0xfc, 0x87,
				0x51, 0x9d, 0x40, 0x4, 0x39, 0x9f, 0x4d, 0xe3, 0xc9, 0x43},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetBidKey(tc.poolId, tc.bidder)
			s.Require().Equal(tc.expected, key)
		})
	}
}

func (s *keysTestSuite) TestGetBidByPoolIdPrefix() {
	s.Require().Equal([]byte{0xe6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetBidByPoolIdPrefix(0))
	s.Require().Equal([]byte{0xe6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetBidByPoolIdPrefix(9))
	s.Require().Equal([]byte{0xe6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetBidByPoolIdPrefix(10))
}

func (s *keysTestSuite) TestGetWinningBidKey() {
	testCases := []struct {
		poolId    uint64
		auctionId uint64
		expected  []byte
	}{
		{
			1,
			1,
			[]byte{0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			1,
			5,
			[]byte{0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
		{
			5,
			5,
			[]byte{0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetWinningBidKey(tc.poolId, tc.auctionId)
			s.Require().Equal(tc.expected, key)
		})
	}
}

func (s *keysTestSuite) TestLengthPrefixTimeBytes() {
	sampleTime1 := utils.ParseTime("2022-07-01T00:00:00Z")
	sampleTime2 := utils.ParseTime("2022-08-01T00:00:00Z")

	testCases := []struct {
		sampleTime time.Time
		length     int
		expected   []byte
	}{
		{
			sampleTime1,
			len(sdk.FormatTimeBytes(sampleTime1)),
			[]byte{0x1d, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x37, 0x2d, 0x30,
				0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30,
				0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30},
		},
		{
			sampleTime2,
			len(sdk.FormatTimeBytes(sampleTime2)),
			[]byte{0x1d, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x38, 0x2d, 0x30,
				0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30,
				0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30},
		},
	}

	for _, tc := range testCases {
		bz := types.LengthPrefixTimeBytes(tc.sampleTime)
		s.Require().Equal(tc.length, int(bz[0]))
		s.Require().Equal(tc.expected, bz)
	}
}
