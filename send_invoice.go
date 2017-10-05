package telegram

import json "github.com/pquerna/ffjson/ffjson"

type SendInvoiceParameters struct {
	// Unique identifier for the target private chat
	ChatID int64 `json:"chat_id"` // required

	// Product name, 1-32 characters
	Title string `json:"title"` // required

	// Product description, 1-255 characters
	Description string `json:"description"` // required

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	Payload string `json:"payload"` // required

	// Payments provider token, obtained via Botfather
	ProviderToken string `json:"provider_token"` // required

	// Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
	StartParameter string `json:"start_parameter"` // required

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"` // required

	// Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"` // required

	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"` // optional

	// Photo size
	PhotoSize int `json:"photo_size,omitempty"` // optional

	// Photo width
	PhotoWidth int `json:"photo_width,omitempty"` // optional

	// Photo height
	PhotoHeight int `json:"photo_height,omitempty"` // optional

	// Pass True, if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"` // optional

	// Pass True, if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"` // optional

	// Pass True, if you require the user's email to complete the order
	NeedEmail bool `json:"need_email,omitempty"` // optional

	// Pass True, if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"` // optional

	// Pass True, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"` // optional

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"` // optional

	// If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"` // optional

	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // optional
}

// SendInvoice send invoices. On success, the sent Message is returned.
func (bot *Bot) SendInvoice(params *SendInvoiceParameters) (*Message, error) {
	dst, err := json.Marshal(*params)
	if err != nil {
		return nil, err
	}

	resp, err := bot.request(dst, "sendInvoice", nil)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(*resp.Result, &data)
	return &data, err
}