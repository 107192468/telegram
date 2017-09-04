package telegram

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// UnpinChatMessage unpin a message in a supergroup chat. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (bot *Bot) UnpinChatMessage(chat interface{}) (bool, error) {
	var args http.Args
	switch id := chatID.(type) {
	case int64: // Unique identifier for the target chat...
		args.Add("chat_id", strconv.FormatInt(id, 10))
	case string: // ...or username of the target supergroup or channel (in the format @supergroupusername)
		args.Add("chat_id", id)
	default:
		return nil, errors.New(errorInt64OrString)
	}

	resp, err := bot.post("unpinChatMessage", &args)
	if err != nil {
		return nil, err
	}

	var data bool
	err = json.Unmarshal(*resp.Result, &data)
	return data, err
}
