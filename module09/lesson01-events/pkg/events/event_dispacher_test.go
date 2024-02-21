package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func (h *TestEventHander) Handler(event EventInterface) {}

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
	assert.Equal(suite.T(),&suite.handler, suite.eventDispacher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(),&suite.handler2, suite.eventDispacher.handlers[suite.event.GetName()][1])

	}	
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispacherTestSuite))
}
