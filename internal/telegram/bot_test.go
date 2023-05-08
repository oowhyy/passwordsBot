package telegram

import (
	"fmt"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/oowhyy/passwordbot/internal/storage"
	"github.com/oowhyy/passwordbot/internal/storage/redis"
	"github.com/stretchr/testify/suite"
)

var testUser1 = "testUser1"
var testUser2 = "testUser2"

func newBaseMessage(want string) *tgbotapi.Message {
	return &tgbotapi.Message{
		MessageID: 123,
		From: &tgbotapi.User{
			UserName: testUser1,
		},
		Chat: &tgbotapi.Chat{
			ID: 42,
		},
		Text: want,
	}
}

func newBaseCommand(command string) *tgbotapi.Message {
	return &tgbotapi.Message{
		MessageID: 123,
		From: &tgbotapi.User{
			UserName: testUser1,
		},
		Chat: &tgbotapi.Chat{
			ID: 42,
		},
		Text:     command,
		Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Length: len(command), Offset: 0}},
	}
}

// Integration tests

type BotTestSuite struct {
	suite.Suite
	bot *Bot
}

func TestBottTestSuite(t *testing.T) {
	suite.Run(t, &BotTestSuite{})
}

func (bts *BotTestSuite) SetupSuite() {
	// test DB: 1
	// test expire time: 1 second
	testdb, err := redis.NewRedisStorage("localhost:6379", "", 1)
	if err != nil {
		bts.FailNowf("unable to connect to database", err.Error())
	}
	bts.bot = NewBot(nil, testdb)
}

func (bts *BotTestSuite) SetupTest() {
	bts.bot.storage.Set(testUser1, &storage.Item{
		Service:  "predefined",
		Password: "password",
	})
}

func (bts *BotTestSuite) TearDownTest() {
	bts.bot.userStates = map[string]State{}
	bts.bot.storage.Delete(testUser1, "predefined")
}

func (bts *BotTestSuite) TestStart() {

	message := newBaseCommand("/start")
	res := bts.bot.HandleAny(message).Text
	want := fmt.Sprintf(msgStart, int(bts.bot.expire.Seconds()))
	bts.Require().Equal(want, res)
}

func (bts *BotTestSuite) TestUnknown() {
	message := newBaseMessage("hi")
	res := bts.bot.HandleAny(message).Text
	bts.Require().Equal(msgUnknownCommand, res)
}

func (bts *BotTestSuite) TestSetNew() {
	message := newBaseCommand("/set")
	res := bts.bot.HandleAny(message).Text
	want := msgNewSet
	bts.Require().Equal(want, res)
}

func (bts *BotTestSuite) TestSetDialogue() {
	message := newBaseCommand("/set")
	res := bts.bot.HandleAny(message).Text
	want := msgNewSet
	bts.Require().Equal(want, res)
	message = newBaseMessage("serv pass")
	res = bts.bot.HandleAny(message).Text
	want = fmt.Sprintf(msgOKSet, "pass", "serv")
	bts.Require().Equal(want, res)
}

func (bts *BotTestSuite) TestSetGet() {
	// set step 1
	message := newBaseCommand("/set")
	res := bts.bot.HandleAny(message).Text
	want := msgNewSet
	bts.Require().Equal(want, res)
	// set step 2
	message = newBaseMessage("serv pass")
	res = bts.bot.HandleAny(message).Text
	want = fmt.Sprintf(msgOKSet, "pass", "serv")
	bts.Require().Equal(want, res)
	// get step 1
	message = newBaseCommand("/get")
	res = bts.bot.HandleAny(message).Text
	want = msgNewGet
	bts.Require().Equal(want, res)
	// get step 2
	message = newBaseMessage("serv")
	res2 := bts.bot.HandleAny(message)
	want = "pass"
	bts.Require().Equal(want, res2.Text)
	bts.Require().Equal(res2.ReplyToMessageID, message.MessageID)
}

func (bts *BotTestSuite) TestDelDialogue() {
	// del step 2
	message := newBaseCommand("/del")
	res := bts.bot.HandleAny(message).Text
	want := msgNewDel
	bts.Require().Equal(want, res)
	// del step 2
	message = newBaseMessage("predefined")
	res = bts.bot.HandleAny(message).Text
	want = fmt.Sprintf(msgOKDel, "predefined")
	bts.Require().Equal(want, res)
}

func (bts *BotTestSuite) TestTwoUsers() {
	message := newBaseCommand("/get")
	res := bts.bot.HandleAny(message).Text
	want := msgNewGet
	bts.Require().Equal(want, res)
	message2 := newBaseCommand("/get")
	message2.From.UserName = testUser2
	res = bts.bot.HandleAny(message2).Text
	bts.Require().Equal(want, res)
}
