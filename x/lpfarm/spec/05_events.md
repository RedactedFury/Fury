<!-- order: 5 -->

# Events

## Handlers

### MsgCreatePrivatePlan

| Type                                           | Attribute Key        | Attribute Value                                |
|------------------------------------------------|----------------------|------------------------------------------------|
| message                                        | action               | /fury.lpfarm.v1beta1.Msg/CreatePrivatePlan |
| fury.lpfarm.v1beta1.EventCreatePrivatePlan | creator              | {planCreatorAddress}                           |
| fury.lpfarm.v1beta1.EventCreatePrivatePlan | plan_id              | {planId}                                       |
| fury.lpfarm.v1beta1.EventCreatePrivatePlan | farming_pool_address | {farmingPoolAddress}                           |

### MsgFarm

| Type                              | Attribute Key     | Attribute Value                   |
|-----------------------------------|-------------------|-----------------------------------|
| message                           | action            | /fury.lpfarm.v1beta1.Msg/Farm |
| fury.lpfarm.v1beta1.EventFarm | farmer            | {farmerAddress}                   |
| fury.lpfarm.v1beta1.EventFarm | coin              | {coin}                            |
| fury.lpfarm.v1beta1.EventFarm | withdrawn_rewards | {withdrawnRewards}                |

### MsgUnfarm

| Type                                | Attribute Key     | Attribute Value                     |
|-------------------------------------|-------------------|-------------------------------------|
| message                             | action            | /fury.lpfarm.v1beta1.Msg/Unfarm |
| fury.lpfarm.v1beta1.EventUnfarm | farmer            | {farmerAddress}                     |
| fury.lpfarm.v1beta1.EventUnfarm | coin              | {coin}                              |
| fury.lpfarm.v1beta1.EventUnfarm | withdrawn_rewards | {withdrawnRewards}                  |

### MsgHarvest

| Type                                 | Attribute Key     | Attribute Value                      |
|--------------------------------------|-------------------|--------------------------------------|
| message                              | action            | /fury.lpfarm.v1beta1.Msg/Harvest |
| fury.lpfarm.v1beta1.EventHarvest | farmer            | {farmerAddress}                      |
| fury.lpfarm.v1beta1.EventHarvest | denom             | {farmingAssetDenom}                  |
| fury.lpfarm.v1beta1.EventHarvest | withdrawn_rewards | {withdrawnRewards}                   |
