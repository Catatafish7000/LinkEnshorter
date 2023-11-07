package generator

import(
	"math/rand"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type Generator struct{}

func NewGenerator () *Generator{
	return &Generator{}
}

func (g *Generator) GenerateHash() string{
	hash := make([]byte, 10)
	alphLen:=len(alphabet)
	for i := range hash {
		hash[i] = alphabet[rand.Intn(alphLen)]
	}
	return string(hash)
}

