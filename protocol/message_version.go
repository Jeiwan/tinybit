package protocol

// MsgVersion ...
type MsgVersion struct {
	Version     int32
	Services    uint64
	Timestamp   int64
	AddrRecv    VersionNetAddr
	AddrFrom    VersionNetAddr
	Nonce       uint64
	UserAgent   VarStr
	StartHeight int32
	Relay       bool
}
