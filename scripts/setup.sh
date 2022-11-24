make all

rm -rf ~/.fury

mkdir ~/.fury

fury init --chain-id test test
fury keys add test --recover --keyring-backend test<<<"y
wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please"
fury add-genesis-account test 100000000000000stake --keyring-backend test
fury gentx test 1000000000stake --chain-id test --keyring-backend test
fury collect-gentxs
fury start