package export

// Protocols
const (
	ProtoHTTP  = "HTTP"
	ProtoTCP   = "TCP"
	ProtoMAC   = "MAC"
	ProtoZMQ   = "ZMQ"
	ProtoOther = "OTHER"
)

// Methods
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// Addressable - address for reaching the service
type Addressable struct {
	Name      string
	Method    string
	Protocol  string
	Address   string
	Port      int
	Path      string
	Publisher string
	User      string
	Password  string
	Topic     string
}
