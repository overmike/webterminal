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

		switch command := req.Command.(type) {
		case *SessionRequest_Message:
			{
				msg := command.Message
				logrus.Debugf("Request string : %v", msg)
				res := &SessionResponse{Message: msg}
				sendErr := session.Send(res)
				if sendErr == io.EOF {
					logrus.Info("Client closed the session")
					return nil
				} else if sendErr != nil {
					logrus.Errorf("Sending error :%v", sendErr)
					return sendErr
				}
			}
		case *SessionRequest_Resize:
			{
				resize := command.Resize
				logrus.Infof("Request to resize columns %v, rows %v", resize.Columns, resize.Rows)
			}
		case nil:
		default:
			logrus.Warn("Empty SessionRequest command")
		}

	}

}
