package terminal

import (
	"io"
	"os/exec"

	"github.com/kr/pty"
	"github.com/sirupsen/logrus"
)

// Service serve the terminal session
type Service struct{}

type SessionWriter struct {
	session Terminal_SessionServer
}

func (s *SessionWriter) Write(p []byte) (int, error) {
	res := &SessionResponse{Message: string(p)}
	err := s.session.Send(res)
	if err != nil {
		return 0, err
	}

	return len(p), err
}

// Session rpc manage streaming session between client and server
func (*Service) Session(session Terminal_SessionServer) error {
	logrus.Info("Session created")

	c := exec.Command("bash")
	c.Env = append([]string{})
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	defer func() {
		err = ptmx.Close()
		if err != nil {
			logrus.Errorf("Failed to close ptmx %v", err)
		} else {
			logrus.Info("Closing ptmx")
		}
		err := c.Wait()
		if err != nil {
			logrus.Errorf("Command wait fail : %v", err)
		}
		logrus.Info("Closed command")
	}()

	sWriter := &SessionWriter{session: session}
	go func() { io.Copy(sWriter, ptmx) }()

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
				_, err := ptmx.Write([]byte(msg))
				if err != nil {
					return err
				}
			}
		case *SessionRequest_Resize:
			{
				resize := command.Resize
				logrus.Infof("Request to resize columns %v, rows %v", resize.Columns, resize.Rows)
				ws := &pty.Winsize{Cols: uint16(resize.Columns), Rows: uint16(resize.Rows)}
				pty.Setsize(ptmx, ws)
			}
		case nil:
		default:
			logrus.Warn("Empty SessionRequest command")
		}

	}

}
