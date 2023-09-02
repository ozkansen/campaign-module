package campaign

import (
	"errors"
)

var (
	ErrInvalidValue         = errors.New("invalid value")
	ErrCampaignAlreadyExist = errors.New("campaign already exists")
	ErrCampaignNotFound     = errors.New("campaign not found")
)
