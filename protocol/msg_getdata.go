package protocol

// MsgGetData represents 'getdata' message.
type MsgGetData struct {
	Count     uint8 // TODO: Change to var_int
	Inventory []InvVector
}
