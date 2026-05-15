package Handlers

import (
	"LangBot/Internal/models"
	"LangBot/Internal/service"

	tele "gopkg.in/telebot.v3"
)

type UserHandler struct {
	sv  *service.UserService
	ct  *service.CategoryService
	bot *tele.Bot
}

func NewUserHandler(sv *service.UserService, ct *service.CategoryService, bot *tele.Bot) *UserHandler {
	return &UserHandler{sv: sv, ct: ct, bot: bot}

}

func (u *UserHandler) Register() {
	u.bot.Handle("/start", u.Start)
}

func (u *UserHandler) Start(c tele.Context) error {
	uId := c.Sender().ID
	_, err := u.sv.GetUserService(uId)
	if err != nil {
		err = u.sv.CreateUserService(&models.User{TgId: uId, Username: c.Sender().Username})
		if err != nil {
			return err
		}
		def := u.ct.DefaultCtService(int(uId))
		if def != nil {
			return err
		}
		return c.Send("Привет! Я бот математик расходов,помогу тебе посчитать и автоматизировать твои расходы!")
	}
	return c.Send("Привет! Я бот математик расходов,помогу тебе посчитать и автоматизировать твои расходы!")

}
