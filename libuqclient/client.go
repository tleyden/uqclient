package uqclient

type UniqushClient struct {
	UniqushURL string
}

func NewUniqushClient(url string) *UniqushClient {
	return &UniqushClient{
		UniqushURL: url,
	}
}

func (c UniqushClient) NewService(serviceName string, pst PushServiceType) *UniqushService {
	return &UniqushService{
		UniqushClient:   c,
		Name:            serviceName,
		PushServiceType: pst,
	}
}
