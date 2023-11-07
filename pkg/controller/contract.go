package controller

import "context"

type service interface {
	ShowLink(ctx context.Context,hash string) (string, error)
	SaveShortURL(url string) (string,error)
}
