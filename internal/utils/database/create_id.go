package database_utils

import "github.com/lithammer/shortuuid/v3"

func CreateID() string {
	return shortuuid.New()

}
