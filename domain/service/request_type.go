package service

var (
	RANDOM = RequestTypeEnum{"rand", 1}
	ALL    = RequestTypeEnum{"all", 2}
)

// RequestType defines how svad server send http-requests to registered servers.
type RequestTypeEnum struct {
	StringValue string
	ID          int
}

// GetRequestTypeEnum return RequestTypeEnum from http-header's string value
func GetRequestTypeEnum(value string) *RequestTypeEnum {
	switch value {
	case RANDOM.StringValue:
		return &RANDOM
	case ALL.StringValue:
		return &ALL
	default:
		return &RANDOM
	}
}
