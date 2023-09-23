package domain

type User struct {
	Name string
	Id   string
}

type UserCollection []User

func (uc UserCollection) Contains(u User) bool {
	for _, v := range uc {
		if v.Alike(u) {
			return true
		}
	}
	return false
}

func (uc UserCollection) Len() int {
	return len(uc)
}

func (u User) Alike(subject User) bool {
	return u.Id == subject.Id
}
