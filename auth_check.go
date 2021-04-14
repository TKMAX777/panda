package panda

func (p *Handler) sustainAuth() (err error) {
	c := p.GetOwnInfo()

	if c.Author == "" {
		return p.Login(p.auth.ID, p.auth.Pass)
	}

	return
}
