package cache

import (
	"errors"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

type link struct {
	url       string
	createdAt time.Time
}
type repo struct {
	data map[string]link
	mx   sync.Mutex
}

func NewRepo() *repo {
	data := make(map[string]link)
	return &repo{
		data: data,
		mx:   sync.Mutex{},
	}
}

func (r *repo) SaveHashByURL(url, hash string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	_, ok := r.data[hash]
	if ok {
		err := errors.New("Error: duplicate key value violates unique constraint")
		return err
	}
	r.data[hash] = link{
		url:       url,
		createdAt: time.Now(),
	}
	return nil
}

func (r *repo) GetURL(hash string) (string, error) {

	ans, ok := r.data[hash]
	if !ok {
		err := errors.New("no such hash in cache")
		return ans.url, err
	}
	return ans.url, nil
}

func (r *repo) Clear() {
	current := time.Now()
	for i := range r.data {
		if current.Sub(r.data[i].createdAt) >= time.Second {
			delete(r.data, i)
		}
	}
}
