package config

type SilentError struct{}

func (se SilentError) Error() string {
	return ""
}
