package protocol

const (
	// Version ...
	Version = 70015

	// SrvNodeNetwork This node can be asked for full blocks instead of just headers.
	SrvNodeNetwork = 1
	// SrvNodeGetUTXO See BIP 0064
	SrvNodeGetUTXO = 2
	// SrvNodeBloom See BIP 0111
	SrvNodeBloom = 4
	// SrvNodeWitness See BIP 0144
	SrvNodeWitness = 8
	// SrvNodeNetworkLimited See BIP 0159
	SrvNodeNetworkLimited = 1024
)

// NewUserAgent ...
func NewUserAgent(userAgent string) VarStr {
	return newVarStr(userAgent)
}
