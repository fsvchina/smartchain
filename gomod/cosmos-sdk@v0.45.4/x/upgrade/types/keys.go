package types

import "fmt"

const (

	ModuleName = "upgrade"


	RouterKey = ModuleName


	StoreKey = ModuleName


	QuerierKey = ModuleName
)

const (

	PlanByte = 0x0

	DoneByte = 0x1


	VersionMapByte = 0x2


	ProtocolVersionByte = 0x3


	KeyUpgradedIBCState = "upgradedIBCState"


	KeyUpgradedClient = "upgradedClient"


	KeyUpgradedConsState = "upgradedConsState"
)



func PlanKey() []byte {
	return []byte{PlanByte}
}




func UpgradedClientKey(height int64) []byte {
	return []byte(fmt.Sprintf("%s/%d/%s", KeyUpgradedIBCState, height, KeyUpgradedClient))
}




func UpgradedConsStateKey(height int64) []byte {
	return []byte(fmt.Sprintf("%s/%d/%s", KeyUpgradedIBCState, height, KeyUpgradedConsState))
}
