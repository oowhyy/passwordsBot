package telegram

import (
	"fmt"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/oowhyy/passwordbot/internal/storage/redis"
	"github.com/stretchr/testify/suite"
)

func newBaseMessage(text string) *tgbotapi.Message {
	return &tgbotapi.Message{
		MessageID: 123,
		From: &tgbotapi.User{
			UserName: "testUser",
		},
		Chat: &tgbotapi.Chat{
			ID: 42,
		},
		Text: text,
	}
}

func newBaseCommand(command string) *tgbotapi.Message {
	return &tgbotapi.Message{
		MessageID: 123,
		From: &tgbotapi.User{
			UserName: "testUser",
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
	testdb, err := redis.NewRedisStorage("localhost:6379", "", 1, 1)
	if err != nil {
		bts.FailNowf("unable to connect to database", err.Error())
	}
	bts.bot = NewBot(nil, testdb)
}

// func (its *IntTestSuite) BeforeTest(suiteName, testName string) {
// 	if testName == "TestCalculate_Error" {
// 		return
// 	}
// 	seedTestTable(its, its.db) // ts -> price=1, ts+1min -> price=2
// }

// func (its *IntTestSuite) TearDownSuite() {
// 	tearDownDatabase(its)
// }

// func (its *IntTestSuite) TearDownTest() {
// 	cleanTable(its)
// }

func (bts *BotTestSuite) TestStart() {

	message := newBaseCommand("/start")
	res := bts.bot.HandleAny(message).Text
	text := fmt.Sprintf(msgStart, bts.bot.storage.Expire())
	bts.Require().Equal(text, res)
}

func (bts *BotTestSuite) TestUnknown() {
	message := newBaseMessage("hi")
	res := bts.bot.HandleAny(message).Text
	bts.Require().Equal(msgUnknownCommand, res)
}

// func (its *IntTestSuite) TestCalculate() {

// 	actual, err := its.calculator.PriceIncrease()

// 	its.Nil(err)
// 	its.Equal(100.0, actual)

// }

// Helper functions

// func seedTestTable(its *IntTestSuite, db *sql.DB) {
// 	its.T().Log("seeding test table")

// 	for i := 1; i <= 2; i++ {
// 		_, err := db.Exec("INSERT INTO stockprices (timestamp, price) VALUES ($1,$2)", time.Now().Add(time.Duration(i)*time.Minute), float64(i))
// 		if err != nil {
// 			its.FailNowf("unable to seed table", err.Error())
// 		}
// 	}
// }

// func cleanTable(its *IntTestSuite) {
// 	its.T().Log("cleaning database")

// 	_, err := its.db.Exec(`DELETE FROM stockprices`)
// 	if err != nil {
// 		its.FailNowf("unable to clean table", err.Error())
// 	}
// }

// func tearDownDatabase(its *IntTestSuite) {
// 	its.T().Log("tearing down database")

// 	_, err := its.db.Exec(`DROP TABLE stockprices`)
// 	if err != nil {
// 		its.FailNowf("unable to drop table", err.Error())
// 	}

// 	_, err = its.db.Exec(`DROP DATABASE stockprices_test`)
// 	if err != nil {
// 		its.FailNowf("unable to drop database", err.Error())
// 	}

// 	err = its.db.Close()
// 	if err != nil {
// 		its.FailNowf("unable to close database", err.Error())
// 	}
// }
