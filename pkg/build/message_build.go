package build

import . "github.com/felipe-brito/splunk-go-log/pkg/model"

// 	MessageBuilder :
// 	@WithSource
// 	@WithSourceType
// 	@WithIndex
// 	@WithEvent
// 	@Build
type MessageBuilder interface {
	WithSource(source string) MessageBuilder
	WithSourceType(sourceType string) MessageBuilder
	WithIndex(index string) MessageBuilder
	WithEvent(event interface{}) MessageBuilder
	Build() *Message
}

// 	ConcreteMessageBuilder :
// 	Message
type ConcreteMessageBuilder struct {
	Message *Message
}

// 	MessageBuild :
func MessageBuild() ConcreteMessageBuilder {
	return ConcreteMessageBuilder{
		Message: &Message{},
	}
}

func (c ConcreteMessageBuilder) WithSource(source string) MessageBuilder {
	c.Message.Source = source
	return c
}

func (c ConcreteMessageBuilder) WithSourceType(sourceType string) MessageBuilder {
	c.Message.SourceType = sourceType
	return c
}

func (c ConcreteMessageBuilder) WithIndex(index string) MessageBuilder {
	c.Message.Index = index
	return c
}

func (c ConcreteMessageBuilder) WithEvent(event interface{}) MessageBuilder {
	c.Message.Event = event
	return c
}

func (c ConcreteMessageBuilder) Build() *Message {
	c.Message.SetTime()
	c.Message.SetHost()
	return c.Message
}
