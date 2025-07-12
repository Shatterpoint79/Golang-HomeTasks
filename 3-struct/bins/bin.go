package bins

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// Bin представляет один бинарный объект
type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []*Bin `json:"bins"`
}

func NewBin(name string, isPrivate bool) *Bin {
	return &Bin{
		ID:        generateID(),
		Name:      name,
		Private:   isPrivate,
		CreatedAt: time.Now().UTC(),
	}
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]*Bin, 0),
	}
}

func (bl *BinList) AddBin(b *Bin) {
	bl.Bins = append(bl.Bins, b)
}

// От себя решил добавить генерацию уникального ID
func generateID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(bytes)
}
