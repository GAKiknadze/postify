package posters

import (
	"context"
	"strconv"

	"github.com/GAKiknadze/postify/internal/social"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const TelegramPlatform social.PlatformType = "telegram"

type TelegramConfig struct {
	BotToken string
	ChatID   int64
}

type telegramPoster struct {
	config TelegramConfig
	client *bot.Bot
}

func newTelegramPoster(config TelegramConfig) (*telegramPoster, error) {
	b, err := bot.New(config.BotToken)
	if err != nil {
		return nil, err
	}

	return &telegramPoster{
		config: config,
		client: b,
	}, nil
}

func (t *telegramPoster) Post(ctx context.Context, post social.Post) (string, error) {
	var result *models.Message
	var err error
	if len(post.Media) == 0 {
		result, err = t.client.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: t.config.ChatID,
			Text:   post.Caption,
		})
		if err != nil {
			return "", err
		}
	} else {
		var media []models.InputMedia

		for i, elem := range post.Media {
			var caption *string
			if i == 0 {
				caption = &post.Caption
			}
			switch elem.Type {
			case social.Image:
				media = append(media, &models.InputMediaPhoto{
					Caption:         *caption,
					MediaAttachment: elem.Data,
				})
			case social.Video:
				media = append(media, &models.InputMediaVideo{
					Caption:         *caption,
					MediaAttachment: elem.Data,
				})
			}
		}

		params := &bot.SendMediaGroupParams{
			ChatID: t.config.ChatID,
			Media:  media,
		}

		results, err := t.client.SendMediaGroup(ctx, params)

		if err != nil {
			return "", err
		}
		result = results[0]
	}
	return strconv.Itoa(result.ID), nil
}

func (t *telegramPoster) Validate(ctx context.Context) error {
	var err error
	_, err = t.client.GetMe(ctx)
	if err != nil {
		return err
	}
	_, err = t.client.GetChat(ctx, &bot.GetChatParams{
		ChatID: t.config.ChatID,
	})
	return err
}
