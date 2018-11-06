package terminal

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Service serve the terminal session
type Service struct{}

func (*Service) Session(session Terminal_SessionServer) error {
	logrus.Info("Session created")
	err := session.Send(&SessionResponse{Message: "Terminal Init"})
	if err == io.EOF {
		logrus.Info("Session closed from client")
		return nil
	}
	if err != nil {
		return err
	}

	for {
		req, err := session.Recv()
		if err == io.EOF {
			logrus.Info("Session closed from client")
			return nil
		} else if err != nil {
			logrus.Errorf("Receive error : %v", err)
			return err
		}

		if req.Message == "close" {
			logrus.Info("Client request to close the session, closing from server...")
			return nil
		}
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

	return nil
}
