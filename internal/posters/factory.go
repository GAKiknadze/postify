package posters

import (
	"errors"

	"github.com/GAKiknadze/postify/internal/social"
)

var (
	ErrInvalidConfig       error = errors.New("invalid configuration for the specified platform")
	ErrUnsupportedPlatform error = errors.New("unsupported platform type")
)

var AllowedProviders = []social.PlatformType{TelegramPlatform}

type SocialFactory struct{}

func NewSocialFactory() *SocialFactory {
	return &SocialFactory{}
}

func (f *SocialFactory) Create(platform social.PlatformType, config interface{}) (social.SocialPoster, error) {
	switch platform {
	case TelegramPlatform:
		telegramConfig, ok := config.(TelegramConfig)
		if !ok {
			return nil, ErrInvalidConfig
		}
		poster, err := newTelegramPoster(telegramConfig)
		return poster, err
	default:
		return nil, ErrUnsupportedPlatform
	}
}
