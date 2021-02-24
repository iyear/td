// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// MessagesCreateChatRequest represents TL type `messages.createChat#9cb126e`.
// Creates a new chat.
//
// See https://core.telegram.org/method/messages.createChat for reference.
type MessagesCreateChatRequest struct {
	// List of user IDs to be invited
	Users []InputUserClass `tl:"users"`
	// Chat name
	Title string `tl:"title"`
}

// MessagesCreateChatRequestTypeID is TL type id of MessagesCreateChatRequest.
const MessagesCreateChatRequestTypeID = 0x9cb126e

func (c *MessagesCreateChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Users == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCreateChatRequest) String() string {
	if c == nil {
		return "MessagesCreateChatRequest(nil)"
	}
	type Alias MessagesCreateChatRequest
	return fmt.Sprintf("MessagesCreateChatRequest%+v", Alias(*c))
}

// FillFrom fills MessagesCreateChatRequest from given interface.
func (c *MessagesCreateChatRequest) FillFrom(from interface {
	GetUsers() (value []InputUserClass)
	GetTitle() (value string)
}) {
	c.Users = from.GetUsers()
	c.Title = from.GetTitle()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *MessagesCreateChatRequest) TypeID() uint32 {
	return MessagesCreateChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (c *MessagesCreateChatRequest) TypeName() string {
	return "messages.createChat"
}

// Encode implements bin.Encoder.
func (c *MessagesCreateChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.createChat#9cb126e as nil")
	}
	b.PutID(MessagesCreateChatRequestTypeID)
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.createChat#9cb126e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.createChat#9cb126e: field users element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.Title)
	return nil
}

// GetUsers returns value of Users field.
func (c *MessagesCreateChatRequest) GetUsers() (value []InputUserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in InputUserClassSlice helper.
func (c *MessagesCreateChatRequest) MapUsers() (value InputUserClassSlice) {
	return InputUserClassSlice(c.Users)
}

// GetTitle returns value of Title field.
func (c *MessagesCreateChatRequest) GetTitle() (value string) {
	return c.Title
}

// Decode implements bin.Decoder.
func (c *MessagesCreateChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.createChat#9cb126e to nil")
	}
	if err := b.ConsumeID(MessagesCreateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.createChat#9cb126e: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#9cb126e: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.createChat#9cb126e: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#9cb126e: field title: %w", err)
		}
		c.Title = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesCreateChatRequest.
var (
	_ bin.Encoder = &MessagesCreateChatRequest{}
	_ bin.Decoder = &MessagesCreateChatRequest{}
)

// MessagesCreateChat invokes method messages.createChat#9cb126e returning error if any.
// Creates a new chat.
//
// Possible errors:
//  400 CHAT_INVALID: Invalid chat
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 USERS_TOO_FEW: Not enough users (to create a chat, for example)
//  403 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/messages.createChat for reference.
func (c *Client) MessagesCreateChat(ctx context.Context, request *MessagesCreateChatRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
