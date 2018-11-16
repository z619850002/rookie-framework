package syserr

type RequestError struct {}

func (h RequestError) Error() string{
	return "No such request in the module."
}


type EmptyMessageError struct {

}

func (h EmptyMessageError) Error() string{
	return "The message received is empty."
}