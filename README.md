
A uniqush client library in Go.

At the moment I'm just implementing what I need for the [backend](https://github.com/tleyden/deepstyle) of [deepstyle-ios](https://github.com/tleyden/deepstyle-ios).

## Usage Example

```

uniqushClient := libuqclient.NewUniqushClient(f.UniqushURL)
uniqushService := uniqushClient.NewService("deepstyle", libuqclient.APNS)
subscriber := uniqushService.NewSubscriber(jobDoc.Owner, jobDoc.OwnerDeviceToken)

_, err := subscriber.Create()
if err != nil {
	return err
}

_, err = subscriber.Push(message)
if err != nil {
	return err
}

```
