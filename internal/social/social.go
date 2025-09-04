package social

import (
	"bytes"
	"context"
)

// MediaType определяет тип медиаконтента
type MediaType int

const (
	Image MediaType = iota
	Video
)

// MediaItem представляет медиафайл для публикации
type MediaItem struct {
	Type MediaType
	Data *bytes.Buffer
}

// Post представляет структуру публикуемого контента
type Post struct {
	Caption string
	Media   []MediaItem
	Extra   map[string]interface{}
}

// SocialPoster интерфейс для публикации контента в соцсети
type SocialPoster interface {
	// Post публикует контент в социальную сеть
	// Возвращает идентификатор публикации и ошибку при неудаче
	Post(ctx context.Context, post Post) (string, error)

	// Validate проверяет корректность конфигурации клиента
	Validate(ctx context.Context) error
}

// PlatformType идентифицирует тип социальной сети
type PlatformType string

// PosterFactory создает клиенты для разных платформ
type PosterFactory interface {
	Create(platform PlatformType, config interface{}) (*SocialPoster, error)
}
