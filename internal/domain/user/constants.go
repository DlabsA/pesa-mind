package user

type ChannelType string

const (
	ChannelTypeCash        ChannelType = "Cash"
	ChannelTypeMobileMoney ChannelType = "MobileMoney"
	ChannelTypeBank        ChannelType = "Bank"
)

var ValidChannelTypes = []ChannelType{
	ChannelTypeCash,
	ChannelTypeMobileMoney,
	ChannelTypeBank,
}
