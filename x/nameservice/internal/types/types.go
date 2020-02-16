package types

type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

func NeWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}
