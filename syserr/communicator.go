package syserr

type NoResponseError struct {

}

func (h NoResponseError) Error() string{
	return "No for the request.Maybe the target module is broken or not running."
}