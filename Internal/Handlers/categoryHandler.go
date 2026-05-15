package Handlers

import (
	"LangBot/Internal/models"
	"LangBot/Internal/service"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

type CategoryHandler struct {
	cs  *service.CategoryService
	bot *tele.Bot
}

func NewCategoryHandler(cs *service.CategoryService, bot *tele.Bot) *CategoryHandler {
	return &CategoryHandler{cs: cs, bot: bot}
}

func (c *CategoryHandler) Register() {
	c.bot.Handle("/addcategory", c.CategoryAdd)
	c.bot.Handle("/deletecategory", c.DelCategory)
	c.bot.Handle("/categories", c.Categories)
}

func (c *CategoryHandler) CategoryAdd(ctx tele.Context) error {
	uId := ctx.Sender().ID
	category := &models.Category{UserId: int(uId), Name: ctx.Message().Text}
	err := c.cs.CreateCategory(category)
	if err != nil {
		return err
	}
	return ctx.Send("Категория успешно добавлена!")
}

func (c *CategoryHandler) Categories(ctx tele.Context) error {
	var res string
	uId := ctx.Sender().ID
	categories, err := c.cs.GetByUserIdService(int(uId))
	if err != nil {
		return err
	}
	for _, cat := range categories {
		res += fmt.Sprintf("%s", cat.Name)
	}
	return ctx.Send("Доступные категории:", res)

}

func (c *CategoryHandler) DelCategory(ctx tele.Context) error {
	catId := ctx.Message().Text
	trimmed := strings.TrimPrefix(catId, "/deletecategory ")
	parsed, err := strconv.Atoi(trimmed)
	if err != nil {
		return err
	}
	err = c.cs.DeleteCategoryService(parsed)
	if err != nil {
		return err
	}
	return ctx.Send("Успешно удалено!")

}
