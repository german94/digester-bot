package tbot

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

type MockFileHelper struct {
	Content []byte
	Path    string
}

func (mfh *MockFileHelper) GetFileFromURL(url string) ([]byte, error) {
	return mfh.Content, nil
}

func (mfh *MockFileHelper) StoreFileLocally(data []byte, chatID int64, fileName string) (string, error) {
	return mfh.Path, nil
}

type MockBot struct {
}

func (mt *MockBot) Handle(endpoint interface{}, h tele.HandlerFunc, m ...tele.MiddlewareFunc) {
}

func (mt *MockBot) Start() {
}

func (mt *MockBot) Token() string {
	return ""
}

func (mt *MockBot) FileByID(fileID string) (tele.File, error) {
	return tele.File{}, nil
}

type MockContext struct {
	UserID int64
	ChatID int64
	Msg    *tele.Message

	ReceivedMessage string
}

func (mc *MockContext) Chat() *tele.Chat {
	return &tele.Chat{
		ID: mc.ChatID,
	}
}

func (mc *MockContext) Sender() *tele.User {
	return &tele.User{
		ID: mc.UserID,
	}
}

func (mc *MockContext) Send(what interface{}, opts ...interface{}) error {
	text, ok := what.(string)
	if ok {
		mc.ReceivedMessage = text
	}
	return nil
}

func (mc *MockContext) Message() *tele.Message {
	return mc.Msg
}

// ////////////////////
// Unused functions

func (mc *MockContext) Bot() *tele.Bot {
	return nil
}
func (mc *MockContext) Update() tele.Update {
	return tele.Update{}
}

func (mc *MockContext) Callback() *tele.Callback {
	return nil
}
func (mc *MockContext) Query() *tele.Query {
	return nil
}
func (mc *MockContext) InlineResult() *tele.InlineResult {
	return nil
}
func (mc *MockContext) ShippingQuery() *tele.ShippingQuery {
	return nil
}
func (mc *MockContext) PreCheckoutQuery() *tele.PreCheckoutQuery {
	return nil
}
func (mc *MockContext) Poll() *tele.Poll {
	return nil
}
func (mc *MockContext) PollAnswer() *tele.PollAnswer {
	return nil
}
func (mc *MockContext) ChatMember() *tele.ChatMemberUpdate {
	return nil
}
func (mc *MockContext) ChatJoinRequest() *tele.ChatJoinRequest {
	return nil
}
func (mc *MockContext) Migration() (int64, int64) {
	return 0, 0
}
func (mc *MockContext) Topic() *tele.Topic {
	return nil
}

func (mc *MockContext) Recipient() tele.Recipient {
	return nil
}
func (mc *MockContext) Text() string {
	return ""
}
func (mc *MockContext) Entities() tele.Entities {
	return nil
}
func (mc *MockContext) Data() string {
	return ""
}
func (mc *MockContext) Args() []string {
	return nil
}

func (mc *MockContext) SendAlbum(a tele.Album, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) Reply(what interface{}, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) Forward(msg tele.Editable, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) ForwardTo(to tele.Recipient, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) Edit(what interface{}, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) EditCaption(caption string, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) EditOrSend(what interface{}, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) EditOrReply(what interface{}, opts ...interface{}) error {
	return nil
}
func (mc *MockContext) Delete() error {
	return nil
}
func (mc *MockContext) DeleteAfter(d time.Duration) *time.Timer {
	return nil
}
func (mc *MockContext) Notify(action tele.ChatAction) error {
	return nil
}
func (mc *MockContext) Ship(what ...interface{}) error {
	return nil
}
func (mc *MockContext) Accept(errorMessage ...string) error {
	return nil
}
func (mc *MockContext) Answer(resp *tele.QueryResponse) error {
	return nil
}
func (mc *MockContext) Respond(resp ...*tele.CallbackResponse) error {
	return nil
}
func (mc *MockContext) Get(key string) interface{} {
	return nil
}
func (mc *MockContext) Set(key string, val interface{}) {
}
