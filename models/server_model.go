package models

import (
	"errors"
	"fmt"
	"strings"
)

type Server struct {
	Ip   string `json:"ip"`   // ip
	Port string `json:"port"` // 端口
}

func StringToServer(str string) (server *Server, err error) {
	list := strings.Split(str, ":")
	if len(list) != 2 {

		return nil, errors.New("err")
	}

	server = &Server{
		Ip:   list[0],
		Port: list[1],
	}
	return
}

func (s *Server) String() (str string) {
	if s == nil {
		return
	}

	str = fmt.Sprintf("%s:%s", s.Ip, s.Port)

	return
}
