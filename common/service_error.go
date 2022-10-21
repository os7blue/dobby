package common

type ServiceError struct {
	Msg string
}

func (s *ServiceError) Error() string {

	return s.Msg

}
