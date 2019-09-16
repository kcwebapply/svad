package common

const (
	SERVICE_NAME_HEADER_NAME = "service_name"
	REQUEST_TYPE_HEADER_NAME = "request_type"
)

func GetSvadHeaders() []string {
	return []string{
		SERVICE_NAME_HEADER_NAME, REQUEST_TYPE_HEADER_NAME,
	}
}
