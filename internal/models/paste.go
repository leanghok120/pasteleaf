package models

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"
)

const idCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Paste struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var pastes = make(map[string]Paste)
var mu sync.RWMutex

func GenerateID(length int) (string, error) {
	id := make([]byte, length)
	for i := range id {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(idCharset))))
		if err != nil {
			return "", err
		}
		id[i] = idCharset[n.Int64()]
	}
	return string(id), nil
}

func SavePaste(paste Paste) {
	mu.Lock()
	pastes[paste.ID] = paste
	mu.Unlock()
}

func GetPaste(id string) (Paste, bool) {
	mu.RLock()
	paste, ok := pastes[id]
	mu.RUnlock()
	return paste, ok
}

func GetPastes() []Paste {
	mu.RLock()

	list := make([]Paste, 0, len(pastes))
	for _, p := range pastes {
		list = append(list, p)
	}

	mu.RUnlock()

	return list
}
