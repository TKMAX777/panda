package panda

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"regexp"
)

// Login logins panda with put ID / PASSWORD
func (p *Handler) Login(id, pass string) error {
	var uri = BaseURI + "/cas/login"
	var loginQue = url.Values{}
	loginQue.Add("service", BaseURI+"/sakai-login-tool/container")

	res, err := p.get(uri, loginQue)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)
	io.Copy(b, res.Body)
	res.Body.Close()

	var jsessionID string

	{
		var cookies = p.jar.Cookies(&url.URL{
			Host:   "panda.ecs.kyoto-u.ac.jp",
			Path:   "/cas",
			Scheme: "https",
		})

		for _, c := range cookies {
			if c.Name == "JSESSIONID" {
				jsessionID = c.Value
			}
		}
	}

	if jsessionID == "" {
		return fmt.Errorf("not found jsessionid")
	}

	var ltReg = regexp.MustCompile(`<input\s+type="hidden"\s+name="lt"\s+value="(\S+)"\s+/>`)
	if !ltReg.Match(b.Bytes()) {
		return fmt.Errorf("Not found lt")
	}

	var loginOpt = url.Values{}
	var mt = ltReg.FindAllStringSubmatch(b.String(), 1)

	loginOpt.Add("lt", mt[0][1])
	loginOpt.Add("execution", "e1s1")
	loginOpt.Add("username", id)
	loginOpt.Add("password", pass)
	loginOpt.Add("_eventId", "submit")
	loginOpt.Add("submit", "ログイン")

	res, err = p.post(uri+";jsessionid="+jsessionID, loginQue, loginOpt)
	if err != nil {
		return err
	}

	res.Body.Close()

	p.auth.ID = id
	p.auth.Pass = pass

	p.authOK = true

	return nil
}
