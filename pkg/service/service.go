package service

import (
	"math/rand"
	"time"
	"strings"
	"fmt"
	"log"
)

const Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const AlphLen = 63

type Repo interface{
	GetURL(hash string) (string,error)
	SaveHashByURL(url, hash string) error
	Clear()
}

type Service struct{
	Repo Repo
	generator Generator
}

func NewService(Repo Repo, generator Generator) *Service{
	return &Service{
		Repo,
		generator,
	}
}
type Generator interface{
	GenerateHash() (string,error)
}

func (s *Service) ShowLink(hash string) (string,error){
	url,err:=s.Repo.GetURL(hash)
	return url,err
}

func(s *Service) SaveShortURL(url string) (string,error){
	rand.Seed(time.Now().UTC().UnixNano())
	var hash string
	var errGen error
	for {
		hash,errGen = s.generator.GenerateHash()
		if errGen!=nil{
			return "",errGen
		}
		err := s.Repo.SaveHashByURL(url, hash)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			continue
		} else {
			log.Println(fmt.Sprintf("Failed to create hash. Error: %v", err))
			return "",err
		}
	}
	return hash, nil
}

