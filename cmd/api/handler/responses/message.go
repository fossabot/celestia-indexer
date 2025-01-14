package responses

import (
	"time"

	"github.com/dipdup-io/celestia-indexer/internal/storage"
	"github.com/dipdup-io/celestia-indexer/internal/storage/types"
)

type Message struct {
	Id       uint64    `example:"321"                       format:"int64"     json:"id"              swaggertype:"integer"`
	Height   uint64    `example:"100"                       format:"int64"     json:"height"          swaggertype:"integer"`
	Time     time.Time `example:"2023-07-04T03:10:57+00:00" format:"date-time" json:"time"            swaggertype:"string"`
	Position uint64    `example:"2"                         format:"int64"     json:"position"        swaggertype:"integer"`
	TxId     uint64    `example:"11"                        format:"int64"     json:"tx_id,omitempty" swaggertype:"integer"`

	Type types.MsgType `example:"MsgCreatePeriodicVestingAccount" json:"type"`

	Data map[string]any `json:"data"`
}

func NewMessage(msg storage.Message) Message {
	return Message{
		Id:       msg.Id,
		Height:   uint64(msg.Height),
		Time:     msg.Time,
		Position: msg.Position,
		Type:     msg.Type,
		TxId:     msg.TxId,
		Data:     msg.Data,
	}
}
