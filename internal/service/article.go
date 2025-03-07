// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IArticle interface {
		Create(ctx context.Context, in model.ArticleCreateInput) (out model.ArticleCreateOutput, err error)
		Delete(ctx context.Context, id uint) (err error)
		Update(ctx context.Context, in model.ArticleUpdateInput) error
		GetList(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
