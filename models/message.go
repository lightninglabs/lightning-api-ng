package models

import (
	"fmt"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

type Message struct {
	Package  *Package
	FileName string

	Name        string
	LongName    string
	FullName    string
	Description string
	Fields      []*Field

	// Source is the github url to the exact line in the proto file where
	// this message is declared. This field is set in Method.Request and
	// Method.Response
	Source string
}

// NewMessage creates a new Message model from a Message definition
func NewMessage(messageDef *gendoc.Message, pkg *Package,
	fileName string) *Message {

	fmt.Printf("Creating message %s with %d fields\n", messageDef.Name,
		len(messageDef.Fields))

	m := &Message{
		Package:     pkg,
		FileName:    fileName,
		Name:        messageDef.Name,
		LongName:    messageDef.LongName,
		FullName:    messageDef.FullName,
		Description: messageDef.Description,

		Fields: make([]*Field, 0),
	}
	for _, fieldDef := range messageDef.Fields {
		field := NewField(*fieldDef, m)
		m.Fields = append(m.Fields, field)
	}

	return m
}

// HasFields returns true if the message has fields
func (m *Message) HasFields() bool {
	return len(m.Fields) > 0
}
