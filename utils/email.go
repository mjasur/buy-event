package utils


const (
	sender = "example@gmail.com"
	password = "password"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendByEmail(data map[string] interface{}) error {
	var receivers []string
	receivers = append(receivers, data ["email"].(string))

	//smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	//
	//message := []byte(fmt.Sprintf("Your purchase:\nProduct:%s\nPrice:%s", data["product"], data["price"]))
	//
	//auth := smtp.PlainAuth("", sender, password, smtpServer.host)
	//err := smtp.SendMail(smtpServer.Address(), auth, sender, receivers, message)
	//if err != nil {
	//	return err
	//}
	return nil
}
