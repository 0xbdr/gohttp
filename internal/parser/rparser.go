package main

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

type request struct {
	Method     string
	Route      string
	Protocol   string
	Parameters map[string]string
	Headers    map[string]string
	Body       []byte
}

func parserequest(client net.Conn) (request, error) {
	var reader = bufio.NewReader(client)
	var rrequest = request{"", "", "", nil, nil, nil}
	var srequest string
	var err error
	var line []byte
	for {
		line, _, err = reader.ReadLine()
		check(err)
		if string(line) != "" {
			srequest += string(line) + "\n"
		} else {
			srequest = strings.TrimSuffix(srequest, "\n")
			break
		}
	}
	rrequest.Parameters = make(map[string]string)
	rrequest.Headers = make(map[string]string)
	var lines = strings.Split(srequest, "\n")
	var rawrequest string = lines[0]
	var rparts = strings.Split(rawrequest, " ")
	if len(rparts) != 3 {
		return request{}, errors.New("Bad request")
	}
	rrequest.Method = rparts[0]
	rrequest.Protocol = rparts[2]
	var qp []string = strings.Split(rparts[1], "?")
	if len(qp) > 2 {
		for i, v := range qp {
			if i > 1 {
				qp[1] += v
			}
		}
	}
	rrequest.Route = qp[0]
	if len(qp) != 1 {
		var params = strings.Split(qp[1], "&")
		for _, v := range params {
			var key strings.Builder
			var value strings.Builder
			var f = false
			for _, vv := range v {
				if vv == '=' && !f {
					f = true
				} else {
					if !f {
						key.WriteString(string(vv))
					} else {
						value.WriteString(string(vv))
					}
				}
			}
			rrequest.Parameters[key.String()] = value.String()
		}

	}
	for i, line := range lines {
		if i != 0 {
			var ss = strings.Split(line, ":")
			var header = ss[0]
			var value, done = strings.CutPrefix(ss[1], " ")
			if done {
				rrequest.Headers[header] = value
			} else {
				if len(ss) == 2 {
					rrequest.Headers[header] = ss[1]
				} else {
					return request{}, errors.New("error reading headers")
				}
			}
		}
	}
	//TODO REQUEST BODY
	return rrequest, err
}
