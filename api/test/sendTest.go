package main

import (
	"chatplus/handler/chatimpl"
	"context"
	"testing"

	"chatplus/core/types"
	"chatplus/store/model"
	"chatplus/store/vo"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestChatHandler_sendMessage(t *testing.T) {
	var chatHandler = newChatHandler()
	ctx := context.Background()
	var session = types.ChatSession{
		SessionId: "1",
		UserId:    100,
		Model: types.ChatModel{
			Id:          1,
			Name:        "gpt3.5",
			Value:       "text-davinci-003",
			Power:       100,
			MaxTokens:   1024,
			MaxContext:  1000,
			Temperature: 0.7,
			Platform:    types.Azure,
		},
	}
	var role = model.ChatRole{
		Name:    "用户",
		Context: "[]",
	}
	var userVo = vo.User{
		Username: "test",
		Power:    100,
	}
	err := chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	userVo.Power = 1
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.NotNil(t, err)
	userVo.Status = false
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.NotNil(t, err)
	userVo.Status = true
	userVo.Power = 1
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.NotNil(t, err)
	userVo.Power = 100
	session.Model.Platform = types.OpenAI
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	session.Model.Platform = types.ChatGLM
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	session.Model.Platform = types.Baidu
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	session.Model.Platform = types.XunFei
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
	session.Model.Platform = types.QWen
	err = chatHandler.SendMessage(ctx, &session, role, "你好！", nil)
	assert.Nil(t, err)
}

func newChatHandler() *chatimpl.ChatHandler {
	chatHandler := chatimpl.NewChatHandler(NewFakeApp(), nil, redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}), nil)
	return chatHandler
}
