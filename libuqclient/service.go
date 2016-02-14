package libuqclient

type UniqushService struct {
	UniqushClient   UniqushClient
	PushServiceType PushServiceType
	Name            string
}

func (svc UniqushService) NewSubscriber(subscriberName, deviceToken string) *UniqushSubscriber {
	return &UniqushSubscriber{
		UniqushService: svc,
		Name:           subscriberName,
		DeviceToken:    deviceToken,
	}

}
