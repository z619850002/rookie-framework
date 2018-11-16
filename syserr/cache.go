package syserr

type CacheConnectionError struct {
	Name  	string
}

func (h CacheConnectionError) Error()string{
	return "Cache " + h.Name + " can`t connect to the source."
}

