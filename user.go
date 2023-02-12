package topazsdk

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string
	Admin    bool   `json:"admin"`
	Email    string `json:"email"`
	Data     string `json:"data"`
	Banned   bool   `json:"banned"`
}

func (m *Manager) PullUser(uid int, password ...string) (*User, error) {
}
