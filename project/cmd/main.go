package main

import (
	"backend/AI"
	verifykey "backend/AI/verifyKey"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//  "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//	"os"
//	"strconv"
//	"path/filepath"

//	"mime/multipart"
//	"io"
//	"time"
//	"log"
//		"gonum.org/v1/plot"
//	"gonum.org/v1/plot/plotter"
//"gonum.org/v1/plot/plotutil"
//"gonum.org/v1/plot/vg"
//"gonum.org/v1/plot/vg/draw"

//	"sort"
//"log"
//	"strings"

type User struct {
	Id, Name, Surname, Speciality, Email, Password string
	YearOfBirth, Expirience                        int
	Is_that_authorized_user                        bool
	Is_it_admin                                    bool
}

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func populateUserFromSession(r *http.Request, sessionName string) (User, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		panic(err)
	}

	user := User{
		Id:                      getSessionValueAsString(session, "curret_user_id"),
		Name:                    getSessionValueAsString(session, "name"),
		Surname:                 getSessionValueAsString(session, "surname"),
		Email:                   getSessionValueAsString(session, "email"),
		Password:                getSessionValueAsString(session, "password"),
		Is_that_authorized_user: getSessionValueAsBool(session, "is_authorized"),
		Is_it_admin:             getSessionValueAsBool(session, "adminAuth"),
	}

	return user, nil
}

func saveUserToSession(r *http.Request, w http.ResponseWriter, sessionName string, user User) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values["curret_user_id"] = user.Id
	session.Values["name"] = user.Name
	session.Values["surname"] = user.Surname
	session.Values["email"] = user.Email
	session.Values["password"] = user.Password
	session.Values["is_authorized"] = user.Is_that_authorized_user
	session.Values["adminAuth"] = user.Is_it_admin

	fmt.Println("Данные о сохранении:  : : :")
	fmt.Println("------------")
	return session.Save(r, w)
}

func returnToLastPage(r *http.Request, w http.ResponseWriter) {
	referer := r.Referer()

	// Если Referer пустой, используем URL по умолчанию
	if referer == "" {
		referer = "buzhor13.ru" // Замените на ваш URL по умолчанию
	}

	// Выполняем редирект
	http.Redirect(w, r, referer, http.StatusFound)
}

func getSessionValueAsString(session *sessions.Session, key string) string {
	if val, ok := session.Values[key].(string); ok {
		return val
	}
	return ""
}

func getSessionValueAsBool(session *sessions.Session, key string) bool {
	if val, ok := session.Values[key].(bool); ok {
		return val
	}
	return false
}

func home_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../frontend/templates/index.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	fmt.Println("GO")
	//var data Data_for_personal_page

	//authorized_user, err  := populateUserFromSession(r, sessionName)
	//data.Authorized_user_data  = authorized_user
	t.ExecuteTemplate(w, "index", nil)
}

// Либо эта функция сюда чисто параметры передавать
func AiResponse(nameRecruit, specialityRecruit string, birthDateRecruit, expirienceRecruit int, nameEmployee, specialityEmployee string, birthDateEmployee, expirienceEmployee int) (string, error) {
	accessToken, err := verifykey.GetAccessToken()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}

	request := fmt.Sprintf("Возвращает ТОЛЬКО число с плавающей точкой (в формате 0.XX) (округленное до двух знаков) - совместимость между двумя людьми на основе их данных."+
		"Имя первого человека (того, кто устраивается на работу) - %s, Год рождения первого человека %d, Опыт работы первого человека (в годах) %d, Специальность первого человека в компании %s \n"+
		"Имя второго человека (уже работающего) - %s, Год рождения второго человека %d, Опыт работы второго человека (в годах) %d, Специальность второго человека в компании %s", nameRecruit, birthDateRecruit, expirienceRecruit, specialityRecruit,
		nameEmployee, birthDateEmployee, expirienceEmployee, specialityEmployee)

	content, err := getChatResponse(accessToken, request)
	if err != nil {
		return "0", err
	}

	return content, nil
}

func getChatResponse(accessToken string, answer string) (string, error) {
	url := "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "GigaChat",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": answer,
			},
		},
		"function_call": "auto",
		"functions": []map[string]interface{}{
			{
				"name":        "Совместимость",
				"description": "Возвращает ТОЛЬКО число с плавающей точкой (в формате 0.XX) (округленное до двух знаков) - совместимость между двумя людьми на основе их данных.",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"first_person_name": map[string]interface{}{
							"type":        "string",
							"description": "Имя первого человека (того, кто устраивается на работу)",
						},
						"second_person_name": map[string]interface{}{
							"type":        "string",
							"description": "Имя второго человека (уже работающего)",
						},
						"first_person_specialty": map[string]interface{}{
							"type":        "string",
							"description": "Специальность первого человека в компании",
						},
						"second_person_specialty": map[string]interface{}{
							"type":        "string",
							"description": "Специальность второго человека в компании",
						},
						"first_person_birth_year": map[string]interface{}{
							"type":        "number",
							"description": "Год рождения первого человека",
						},
						"second_person_birth_year": map[string]interface{}{
							"type":        "number",
							"description": "Год рождения второго человека",
						},
						"first_person_experience": map[string]interface{}{
							"type":        "number",
							"description": "Опыт работы первого человека (в годах)",
						},
						"second_person_experience": map[string]interface{}{
							"type":        "number",
							"description": "Опыт работы второго человека (в годах)",
						},
					},
					"required": []interface{}{
						"first_person_name",
						"second_person_name",
						"first_person_specialty",
						"second_person_specialty",
						"first_person_birth_year",
						"second_person_birth_year",
						"first_person_experience",
						"second_person_experience",
					},
				},
			},
		},
		"stream":             false,
		"repetition_penalty": 1,
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytePayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Отключение проверки сертификата
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", err
	}

	choices, ok := responseData["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no choices found in response")
	}

	choice := choices[0].(map[string]interface{})
	message := choice["message"].(map[string]interface{})
	content := message["content"].(string)

	return content, nil
}

// Либо эта функция сюда просто типы
func AiCheck(recruit, employee User) (float64, error) {
	perCentStr, err := AI.AiResponse(recruit.Name, recruit.Speciality, recruit.YearOfBirth, recruit.Expirience,
		employee.Name, employee.Speciality, employee.YearOfBirth, employee.Expirience)
	perCent, _ := strconv.ParseFloat(perCentStr, 64)
	if err != nil {
		return 0, err
	}
	return perCent, nil
}

func main() {
	store.Options = &sessions.Options{
		Path:     "/",       // Путь для куков
		MaxAge:   86400 * 7, // Время жизни куков - 7 дней
		HttpOnly: true,      // Запрет доступа к кукам через JavaScript
		Secure:   false,     // Для HTTP должен быть false
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r := mux.NewRouter()
	fmt.Println("Start")

	r.HandleFunc("/", home_page)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
