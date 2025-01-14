package responses

import "github.com/dipdup-io/celestia-indexer/internal/storage"

// Address model info
//
//	@Description	Celestia address information
type Address struct {
	Id      uint64 `example:"321"                                             json:"id"           swaggertype:"integer"`
	Height  uint64 `example:"100"                                             json:"first_height" swaggertype:"integer"`
	Balance string `example:"10000000000"                                     json:"balance"      swaggertype:"string"`
	Hash    string `example:"celestia1jc92qdnty48pafummfr8ava2tjtuhfdw774w60" json:"hash"         swaggertype:"string"`
}

func NewAddress(addr storage.Address) Address {
	return Address{
		Id:      addr.Id,
		Height:  uint64(addr.Height),
		Hash:    addr.Address,
		Balance: addr.Balance.Total.String(),
	}
}

func (Address) SearchType() string {
	return "address"
}
