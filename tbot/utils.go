package tbot

import (
	"errors"
	"fmt"
	"time"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) Send(c tele.Context, msg, caption string, menu *tele.ReplyMarkup) error {
	err := c.Send(msg, menu)
	if err == nil {
		return nil
	}
	if errors.Is(err, tele.ErrTooLongMessage) {
		filePath, err := tb.StoreFileLocally([]byte(msg), c.Chat().ID, fmt.Sprintf("%v_%v", caption, time.Now().Unix()))
		if err != nil {
			return err
		}
		return c.Send(&tele.Document{
			File:     tele.FromDisk(filePath),
			Caption:  caption,
			MIME:     "text/plain",
			FileName: caption + ".txt",
		})
	}
	return err
}
