package notification

type Options map[string]string

type Notification interface {
	Notify(event string, path string) error
}
