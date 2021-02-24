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

// MessagesGetEmojiKeywordsLanguagesRequest represents TL type `messages.getEmojiKeywordsLanguages#4e9963b2`.
// Get info about an emoji keyword localization
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsLanguages for reference.
type MessagesGetEmojiKeywordsLanguagesRequest struct {
	// Language codes
	LangCodes []string `tl:"lang_codes"`
}

// MessagesGetEmojiKeywordsLanguagesRequestTypeID is TL type id of MessagesGetEmojiKeywordsLanguagesRequest.
const MessagesGetEmojiKeywordsLanguagesRequestTypeID = 0x4e9963b2

func (g *MessagesGetEmojiKeywordsLanguagesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCodes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiKeywordsLanguagesRequest(nil)"
	}
	type Alias MessagesGetEmojiKeywordsLanguagesRequest
	return fmt.Sprintf("MessagesGetEmojiKeywordsLanguagesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetEmojiKeywordsLanguagesRequest from given interface.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) FillFrom(from interface {
	GetLangCodes() (value []string)
}) {
	g.LangCodes = from.GetLangCodes()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) TypeID() uint32 {
	return MessagesGetEmojiKeywordsLanguagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) TypeName() string {
	return "messages.getEmojiKeywordsLanguages"
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsLanguages#4e9963b2 as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsLanguagesRequestTypeID)
	b.PutVectorHeader(len(g.LangCodes))
	for _, v := range g.LangCodes {
		b.PutString(v)
	}
	return nil
}

// GetLangCodes returns value of LangCodes field.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) GetLangCodes() (value []string) {
	return g.LangCodes
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsLanguages#4e9963b2 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsLanguagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: field lang_codes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: field lang_codes: %w", err)
			}
			g.LangCodes = append(g.LangCodes, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsLanguagesRequest.
var (
	_ bin.Encoder = &MessagesGetEmojiKeywordsLanguagesRequest{}
	_ bin.Decoder = &MessagesGetEmojiKeywordsLanguagesRequest{}
)

// MessagesGetEmojiKeywordsLanguages invokes method messages.getEmojiKeywordsLanguages#4e9963b2 returning error if any.
// Get info about an emoji keyword localization
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsLanguages for reference.
func (c *Client) MessagesGetEmojiKeywordsLanguages(ctx context.Context, langcodes []string) ([]EmojiLanguage, error) {
	var result EmojiLanguageVector

	request := &MessagesGetEmojiKeywordsLanguagesRequest{
		LangCodes: langcodes,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
