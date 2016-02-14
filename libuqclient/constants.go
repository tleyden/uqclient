package libuqclient

type PushServiceType int

const (
	APNS = iota // Apple Push Notification Service
	GCM         // Android Google Cloud Messaging
	ADM         // Kindle Tablets
	C2DM        // Android (Deprecated by google. Should use GCM instead)
)

func (ps PushServiceType) String() string {
	switch ps {
	case APNS:
		return "apns"
	case GCM:
		return "gcm"
	case ADM:
		return "adm"
	case C2DM:
		return "c2dm"
	default:
		return "error - unknown type"
	}
}
