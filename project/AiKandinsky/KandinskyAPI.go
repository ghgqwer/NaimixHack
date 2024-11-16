package AiKandinsky

//пример использования
// в main.go:
//AiKandinsky.GenAvatar(AiKandinsky.User(user1))
//сохраняется в AiKandinsky/users/:id

import (
	"fmt"
	"log"
	"os/exec"
)

type User struct {
	Id, Name, Surname, Speciality, Email, Password, YearOfBirth, Expirience string
	Is_that_authorized_user                                                 bool
	Is_it_admin                                                             bool
}

func GenAvatar(user User) {
	prompt := fmt.Sprintf("Нарисуй абстрактную картинку используя знания из нумерологии"+
		"космограммы и таро, зная такие данные: %s, %s, "+
		"%s года рождения, профессия - %s, %s опыт работы",
		user.Surname, user.Name, user.YearOfBirth, user.Speciality, user.Expirience)

	cmd := exec.Command("python", "../AiKandinsky/main.py", prompt, user.Id, user.Surname, user.Name)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running script: %v", err)
	}

	log.Println("Image generated and saved successfully.")
}
