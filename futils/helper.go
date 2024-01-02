package futils

type FileHelper interface {
	GetFileFromURL(url string) ([]byte, error)
	StoreFileLocally(data []byte, chatID int64, fileName string) (string, error)
}

type Service struct {
}
