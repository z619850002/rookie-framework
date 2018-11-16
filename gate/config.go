package gate

// Config melody configuration struct.
type Config struct {
	Port			  	int				// Open Port of the service.
	WriteWait         	int  			// Milliseconds until write times out.
	PongWait          	int 			// Timeout for waiting on pong.
	PingPeriod        	int 			// Milliseconds between pings.
	MaxMessageSize    	int64         	// Maximum size in bytes of a message.
	MessageBufferSize 	int           	// The max amount of messages that can be in a WsConn buffer before it starts dropping them.
}
