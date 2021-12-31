package main

import "fmt"

const (
	FAIL = iota
	ALERT
	WARN
	SUCCESS
)

type NotificationBuilder struct {
	Title            string
	Subtitle         string
	Message          string
	Image            string
	Icon             string
	Priority         int
	NotificationType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubtitle(subtitle string) {
	nb.Subtitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotificationType(notificationType string) {
	nb.NotificationType = notificationType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	// TODO: Add some validations
	if nb.Priority < FAIL || nb.Priority > SUCCESS {
		return nil, fmt.Errorf("`Priority` cannot be less than %d and more than %d", FAIL, SUCCESS)
	}

	return &Notification{
		title:            nb.Title,
		subtitle:         nb.Subtitle,
		message:          nb.Message,
		image:            nb.Image,
		icon:             nb.Icon,
		priority:         nb.Priority,
		notificationType: nb.NotificationType,
	}, nil
}
