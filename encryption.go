package export

// Encryption types
const (
	EncNone = iota
	EncAes
)

// EncryptionDetails - Provides details for encryption
// of export data per client request
type EncryptionDetails struct {
	Algo       int
	Key        string
	InitVector string
}
