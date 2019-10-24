package rcss

import (
	"bytes"
	"fmt"
)

type Message struct {
	name        string
	values      []string
	submessages []Message
}

func (c *Message) AddValues(values ...string) {
	c.values = append(c.values, values...)
}

func (c *Message) AddSubmessages(messages ...Message) {
	c.submessages = append(c.submessages, messages...)
}

func (c Message) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer

	b.WriteRune('(')
	b.WriteString(c.name)

	for _, v := range c.values {
		b.WriteString(fmt.Sprintf(" %v", v))
	}

	for _, msg := range c.submessages {
		if mb, err := msg.MarshalBinary(); err != nil {
			return nil, err
		} else {
			b.WriteRune(' ')
			b.Write(mb)
		}
	}

	b.WriteRune(')')

	return b.Bytes(), nil
}

func (c *Message) UnmarshalBinary(data []byte) error {
	c.name = ""
	c.values = make([]string, 0)
	c.submessages = make([]Message, 0)

	r := bytes.NewBuffer(data)

	if ch, _, err := r.ReadRune(); err != nil {
		return err
	} else if '(' != ch {
		return fmt.Errorf("corrupted message")
	}

	var name bytes.Buffer
name:
	for ch, _, err := r.ReadRune(); ; ch, _, err = r.ReadRune() {
		if err != nil {
			return err
		}

		switch ch {
		case ' ':
			break name

		case ')':
			r.UnreadByte()
			break name

		default:
			name.WriteRune(ch)
		}
	}
	c.name = name.String()

	var value bytes.Buffer
values:
	for ch, _, err := r.ReadRune(); ; ch, _, err = r.ReadRune() {
		if err != nil {
			return err
		}

		switch ch {
		case '(':
			r.UnreadRune()
			break values

		case ' ':
			c.AddValues(value.String())
			value.Reset()

		case ')':
			if value.Len() > 0 {
				c.AddValues(value.String())
				value.Reset()
			}
			r.UnreadRune()
			break values

		default:
			value.WriteRune(ch)
		}
	}

messages:
	for ch, _, err := r.ReadRune(); ; ch, _, err = r.ReadRune() {
		if err != nil {
			return err
		}

		switch ch {
		case '(':
			var subbuffer bytes.Buffer
			subbuffer.WriteRune(ch)

			for hops := 1; hops > 0; {
				if ch, _, err := r.ReadRune(); err != nil {
					return err
				} else {
					subbuffer.WriteRune(ch)

					switch ch {
					case '(':
						hops++

					case ')':
						hops--
					}
				}
			}

			var submessage Message
			if err := submessage.UnmarshalBinary(subbuffer.Bytes()); err != nil {
				return err
			} else {
				c.AddSubmessages(submessage)
			}

		case ')':
			r.UnreadRune()
			break messages
		}
	}

	if ch, _, err := r.ReadRune(); err != nil {
		return err
	} else if ')' != ch {
		return fmt.Errorf("corrupted message")
	}

	return nil
}

type Marshaler interface {
	MarshalRcss() (Message, error)
}

type Unmarshaler interface {
	UnmarshalRcss(Message) error
}

func NewMessage(name string) *Message {
	return &Message{
		name:        name,
		values:      make([]string, 0),
		submessages: make([]Message, 0),
	}
}

type buffer struct {
	title  string
	value  interface{}
	name   string
	format string

	vars []buffer
}

func (buf buffer) MarshalRcss() (Message, error) {
	msg := NewMessage(buf.name)
	for _, v := range buf.vars {
		var value string
		if len(v.format) > 0 {
			value = fmt.Sprintf(v.format, v.value)
		} else {
			value = fmt.Sprint(v.value)
		}

		if len(v.name) > 0 {
			submsg := NewMessage(v.name)
			submsg.AddValues(value)
			msg.AddSubmessages(*submsg)
		} else {
			msg.AddValues(value)
		}
	}

	return *msg, nil
}

func (buf *buffer) UnmarshalRcss(msg Message) error {
	for k, v := range buf.vars {
		if len(v.format) > 0 {
			if _, err := fmt.Sscanf(msg.values[k], v.name, &v.value); err != nil {
				return fmt.Errorf("error on parsing %s: %s", v.name, err)
			}
		} else {
			if _, err := fmt.Sscan(msg.values[k], &v.value); err != nil {
				return fmt.Errorf("error on parsing %s: %s", v.name, err)
			}
		}
	}

	return nil
}
