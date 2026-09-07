package model

type User struct {
	userUUID            string
	login               string
	email               string
	notificationMethods []NotificationMethod
}
type NotificationMethod struct {
	providerName string
	target       string
}
