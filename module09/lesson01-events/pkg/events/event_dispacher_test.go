package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}
func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHander struct {
	ID int
}

func (h *TestEventHander) Handler(event EventInterface, wg *sync.WaitGroup) {
}

type EventDispacherTestSuite struct {
	suite.Suite
	event          TestEvent
	event2         TestEvent
	handler        TestEventHander
	handler2       TestEventHander
	handler3       TestEventHander
	eventDispacher *EventDispacher
}

func (suite *EventDispacherTestSuite) SetupTest() {
	suite.eventDispacher = NewEventDispacher()
	suite.handler = TestEventHander{
		ID: 1,
	}
	suite.handler2 = TestEventHander{
		ID: 2,
	}
	suite.handler3 = TestEventHander{
		ID: 3,
	}
	suite.event = TestEvent{Name: "test", Payload: "Test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "Test2"}
}
func (suite *EventDispacherTestSuite) TestEventDispacher_Register() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler, suite.eventDispacher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispacher.handlers[suite.event.GetName()][1])
}
func (suite *EventDispacherTestSuite) TestEventDispacher_Register_WithSameHandler() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	suite.eventDispacher.Clear()
	suite.Equal(0, len(suite.eventDispacher.handlers))
}
func (suite *EventDispacherTestSuite) TestEventDispacher_Clear() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event2.GetName()]))
}

func (suite *EventDispacherTestSuite) TestEventDispacher_Has() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispacher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispacher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispacher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (suite *EventDispacherTestSuite) TestEventDispach_Remove() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event2.GetName()]))

	suite.eventDispacher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispacher.handlers[suite.event.GetName()][0])

	suite.eventDispacher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	suite.eventDispacher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispacher.handlers[suite.event2.GetName()]))

}

func (m *MockHandler) Handler(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}
func (suite *EventDispacherTestSuite) TestEventDispach_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handler", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handler", &suite.event)

	suite.eventDispacher.Register(suite.event.GetName(), eh)
	suite.eventDispacher.Register(suite.event.GetName(), eh2)

	suite.eventDispacher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handler", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handler", 1)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispacherTestSuite))
}
