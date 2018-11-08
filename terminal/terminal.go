package terminal

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Service serve the terminal session
type Service struct{}

// Session rpc manage streaming session between client and server
func (*Service) Session(session Terminal_SessionServer) error {
	logrus.Info("Session created")

	for {
		req, err := session.Recv()
		if err == io.EOF {
			logrus.Info("Session closed from client")
			return nil
		} else if err != nil {
			logrus.Errorf("Receive error : %v", err)
			return err
		}

		logrus.Debugf("Request string : %v", req.Message)
		res := &SessionResponse{Message: req.Message}
		sendErr := session.Send(res)
		if sendErr == io.EOF {
			logrus.Info("Client closed the session")
			return nil
		} else if sendErr != nil {
			logrus.Errorf("Sending error :%v", sendErr)
			return sendErr
		}
	}

}
