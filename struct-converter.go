package main

func StudyHallToUser(stud *Stud, Users []*User) *User {
	for _, user := range Users {
		if user.Discord_Name == stud.Nickname {
			return user
		}
	}
	return nil
}
