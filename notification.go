package fofiwano

type Notifcation interface {
	Notify(event string, path string) error
}

