package syserr


//errors for ws conn
type WriteFailError struct {

}

func (h WriteFailError) Error() string{
	return "the ws conn is closed."
}

type SendOnClosedWsConnError struct {

}

func (h SendOnClosedWsConnError) Error() string{
	return "tried to write to closed a ws conn."
}


//errors for hub
type NoMatchedPlayerError struct {

}

func (h NoMatchedPlayerError) Error() string{
	return "no matched player."
}

type ClosedWsConnError struct {

}

func (h ClosedWsConnError) Error() string{
	return "the conn is closed."
}


//errors for gate
type ClosedGate struct {

}

func (h ClosedGate) Error() string{
	return "gate instance is already closed."
}
