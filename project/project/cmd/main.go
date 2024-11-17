package main

//import "backend/internal/server"

import ("fmt";"net/http";
	"html/template"
)
import( "github.com/gorilla/mux"
    "database/sql"
		_ "github.com/go-sql-driver/mysql"

		"project/AI"
		verifykey "project/AI/verifyKey"
		aiKand "project/AiKandinsky"
		"bytes"
		"crypto/tls"
		"encoding/json"

		"os/exec"
		"io/ioutil"
		"log"

		"strconv"



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
		"math"
		"github.com/gorilla/sessions"
		)

type Auth_data struct{
	Is_that_authorized bool

	Is_that_employee bool
	Employee User
	Corporation Corporation_data
}




type Seraching_page_data struct{
		Employees []User
		Style string
}



type User struct{
   Id, Firstname, Surname,Secondname, Email, Password, Date_birth_day, Time_birth, City_of_birth, Work_experience, Speciality string
   Is_that_authorized_user bool
   Is_it_admin bool
	 Suitable_rating string
	 Icon_link string
}

type Corporation_data struct{
	Id, Email, Password, Name string
	Employees []User
}

type Search_corp_data struct{
	Corporations []Corporation_data
	Style string
}
var store = sessions.NewCookieStore([]byte("super-secret-key"))

// var adress_web = "http://localhost:8080"
// //var authorized_user User
// var sessionName = "name_session"
// var adress_sql = "root:@tcp(127.127.126.50)/"
// var adress_data_base_test = adress_sql + "naimix"




//"root:@tcp(127.127.126.50)/test"
//"user:password@tcp(147.45.163.58:3306)/test"
//"user:password@tcp(147.45.163.58:3306)/test"

//http://buzhor13.ru
//"http://147.45.163.58:8080"
//http://localhost:8080
var adress_web = "http://buzhor13.ru"
//var authorized_user User
var sessionName = "name_session"
var adress_sql = "user:password@tcp(147.45.163.58:3306)/"
var adress_data_base_test = adress_sql + "naimix"
var adress_data_base_store = adress_sql + "/store"




func populateUserFromSession(r *http.Request, sessionName string) (Auth_data, error) {
    session, err := store.Get(r, sessionName)
    if err != nil{
      panic(err)
    }

    user := User{
        Id:                    getSessionValueAsString(session, "Curret_user_id"),
        Firstname:                  getSessionValueAsString(session, "Firstname"),
        Surname:               getSessionValueAsString(session, "Surname"),
        Secondname:                 getSessionValueAsString(session, "Secondname"),
				Email:  getSessionValueAsString(session, "Email"),
        Password:              getSessionValueAsString(session, "Password"),

				Date_birth_day:              getSessionValueAsString(session, "Date_birth_day"),
				Time_birth:              getSessionValueAsString(session, "Time_birth"),
				City_of_birth:              getSessionValueAsString(session, "City_of_birth"),
				Work_experience:              getSessionValueAsString(session, "Work_experience"),


        Is_that_authorized_user: getSessionValueAsBool(session, "Is_authorized"),


    }
		user2 := Corporation_data{
  			Id:                    getSessionValueAsString(session, "Id_corporation"),
		}

		var auth_data Auth_data
		auth_data.Employee = user
		auth_data.Corporation = user2
    return auth_data, nil
}




func saveUserToSession(r *http.Request, w http.ResponseWriter, sessionName string, user Auth_data) error{
    session, err := store.Get(r, sessionName)
    if err != nil {
        return err
    }


		session.Values["Id_corporation"] = user.Corporation.Id


    session.Values["Curret_user_id"] = user.Employee.Id
    session.Values["Firstname"] = user.Employee.Firstname
    session.Values["Secondname"] = user.Employee.Secondname
    session.Values["Surname"] = user.Employee.Surname

		session.Values["Date_birth_day"] = user.Employee.Date_birth_day
		session.Values["Time_birth"] = user.Employee.Time_birth
		session.Values["City_of_birth"] = user.Employee.City_of_birth
		session.Values["Work_experience"] = user.Employee.Work_experience


    session.Values["Email"] = user.Employee.Email
		session.Values["Password"] = user.Employee.Password
    session.Values["Is_authorized"] = user.Employee.Is_that_authorized_user



    fmt.Println("Данные о сохранении:  : : :")
    fmt.Println("------------")
    return session.Save(r, w)
}















//---------------------AI----------------------------


func AiResponse(nameRecruit, specialityRecruit string, birthDateRecruit, expirienceRecruit string, nameEmployee, specialityEmployee string, birthDateEmployee, expirienceEmployee string) (string, error) {
	accessToken, err := verifykey.GetAccessToken()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}

	request := fmt.Sprintf("Возвращает ТОЛЬКО число с плавающей точкой (в формате 0.XX) (округленное до двух знаков) - совместимость между двумя людьми на основе их данных."+
		"Имя первого человека (того, кто устраивается на работу) - %s, Год рождения первого человека %d, Опыт работы первого человека (в годах) %d, Специальность первого человека в компании %s \n"+
		"Имя второго человека (уже работающего) - %s, Год рождения второго человека %s, Опыт работы второго человека (в годах) %s, Специальность второго человека в компании %s", nameRecruit, birthDateRecruit, expirienceRecruit, specialityRecruit,
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

func AiCheck(recruit, employee User) (float64, error) {

	perCentStr, err := AI.AiResponse(recruit.Firstname, recruit.Speciality, recruit.Date_birth_day, recruit.Work_experience,
		employee.Firstname, employee.Speciality, employee.Date_birth_day, employee.Work_experience)
	perCent, _ := strconv.ParseFloat(perCentStr, 64)
	if err != nil {
		return 0, err
	}
	return perCent, nil
}




//---------------------AI----------------------------

















func  get_user_employee_by_id(id string) (User){

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()


	var zapros = fmt.Sprintf("SELECT id, firstname, secondname, surname, Work_experience, City_of_birth, Time_birth, Date_birth_day, job_title FROM `users` WHERE id = %s", id)
	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var user User

	for res.Next(){
		err = res.Scan(&user.Id, &user.Firstname,&user.Secondname,&user.Surname, &user.Work_experience,  &user.City_of_birth,  &user.Time_birth, &user.Date_birth_day, &user.Speciality)
	}
	user.Icon_link = "../static2/" + user.Id + "/icon.jpg"
	return user
}



func  get_user_corporation_by_id(id string) (Corporation_data){

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()


	var zapros = fmt.Sprintf("SELECT id, name FROM `corporations` WHERE id = %s", id)
	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var user Corporation_data

	for res.Next(){
		err = res.Scan(&user.Id,&user.Name)
	}
	return user
}


func returnToLastPage(r *http.Request, w http.ResponseWriter){
  referer := r.Referer()

    // Если Referer пустой, используем URL по умолчанию
    if referer == "" {
        referer = "buzhor13.ru" // Замените на ваш URL по умолчанию
    }

    // Выполняем редирект
    http.Redirect(w, r, "/", http.StatusFound)
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

func home_page(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/index.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }

			fmt.Println("GO")
      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "index", nil)
}


func employee_profile(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var current_user_id = vars["id_employee"]

	fmt.Println("prof")
	fmt.Println(current_user_id)
	user := get_user_employee_by_id(current_user_id)

	//user, _  := populateUserFromSession(r, sessionName)
	fmt.Println(user)
}

func employee_authorization_page(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("../frontend/templates/employee_authorization_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())

	}
		db, err := sql.Open("mysql", adress_data_base_test)
		if err != nil{
			panic(err)
		}
		defer db.Close()


		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			password := r.FormValue("password")
			var zapros = fmt.Sprintf("SELECT id, firstname, secondname, surname FROM `users` WHERE email = '%s' AND password = '%s'", email, password)
			res,_ := db.Query(zapros)
			fmt.Println(zapros)
			var user Auth_data
			user.Is_that_employee = true
			for res.Next(){
				err = res.Scan(&user.Employee.Id,&user.Employee.Firstname,&user.Employee.Secondname,&user.Employee.Surname)
			}

			user.Employee.Is_that_authorized_user = true
			saveUserToSession(r, w, sessionName, user)
			http.Redirect(w, r, "/for_recruts", http.StatusSeeOther)
		}


	t.ExecuteTemplate(w, "employee_authorization_page", nil)
}

func employee_registration_page(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/employee_registration_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }


				db, err := sql.Open("mysql", adress_data_base_test)
			  if err != nil{
			    panic(err)
			  }
			  defer db.Close()


			if r.Method == http.MethodPost {
				name := r.FormValue("name")
				surname := r.FormValue("surname")
				//weight := r.FormValue("weight")
				secondname := r.FormValue("secondname")
				email := r.FormValue("email")

				Date_birth_day := r.FormValue("Date_birth_day")
				Time_birth := r.FormValue("Time_birth")
				City_of_birth := r.FormValue("City_of_birth")
				Work_experience := r.FormValue("Work_experience")

				password := r.FormValue("password")
				job_title := r.FormValue("dropdown2")






					result, err := db.Exec("insert into naimix.users ( `firstname`,	`surname`, `secondname`, `job_title`, `email`, `password`, `Date_birth_day`, `Time_birth`, `City_of_birth`, `Work_experience`) values (?, ?,?,?, ?, ?, ?,?,?,?)",name, surname, secondname, job_title, email, password, Date_birth_day, Time_birth, City_of_birth, Work_experience)

					fmt.Println(result)
					if(err != nil){
						fmt.Println(err)
					}
					var zapros = fmt.Sprintf("SELECT id FROM `users` ORDER BY id DESC LIMIT 1;")
					var newU User
					res,_ := db.Query(zapros)
					fmt.Println(zapros)
					for res.Next(){
						err = res.Scan(&newU.Id)
					}
					newU = get_user_employee_by_id(newU.Id)


					var aiUser aiKand.User
					aiUser.Id = newU.Id
					aiUser.Name = newU.Firstname
					aiUser.Surname = newU.Surname
					aiUser.Expirience = newU.Work_experience


					aiKand.GenAvatar(aiUser)


			}

      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "employee_registration_page", nil)
}




func corporation_registration_page(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/corporation_registration_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }


				db, err := sql.Open("mysql", adress_data_base_test)
			  if err != nil{
			    panic(err)
			  }
			  defer db.Close()


			if r.Method == http.MethodPost {
				name := r.FormValue("name")
				email := r.FormValue("email")
				//weight := r.FormValue("weight")
				password := r.FormValue("password")

					result, err := db.Exec("insert into naimix.corporations ( `name`,	`email`, `password`) values (?, ?,?)",name, email, password)

					fmt.Println(result)
					if(err != nil){
						fmt.Println(err)
					}



			}

      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "corporation_registration_page", nil)
}

func corporation_profile(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var current_user_id = vars["id_corporation"]

	fmt.Println("prof")
	fmt.Println(current_user_id)
	user := get_user_corporation_by_id(current_user_id)

	//user, _  := populateUserFromSession(r, sessionName)
	fmt.Println(user)
}

func corporation_authorization_page(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("../frontend/templates/corporation_authorization_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())

	}




		db, err := sql.Open("mysql", adress_data_base_test)
		if err != nil{
			panic(err)
		}
		defer db.Close()


		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			password := r.FormValue("password")
			var zapros = fmt.Sprintf("SELECT id, name FROM `corporations` WHERE email = '%s' AND password = '%s'", email, password)
			res,_ := db.Query(zapros)
			fmt.Println(zapros)
			var user Auth_data
			user.Is_that_employee = true
			for res.Next(){
				err = res.Scan(&user.Corporation.Id,&user.Corporation.Name)
			}


			saveUserToSession(r, w, sessionName, user)
			//http.Redirect(w, r, "/" , http.StatusSeeOther)
			http.Redirect(w, r, "/for_company", http.StatusSeeOther)
		}


	t.ExecuteTemplate(w, "corporation_authorization_page", nil)
}




func search_employees_page (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("../frontend/templates/search_employees_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")


	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var style = vars["style"]
	where_zapr := ""
	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Corporation = get_user_corporation_by_id(authorized_user.Corporation.Id)
	if(style == "0"){
		where_zapr = "SELECT id FROM `users` "
	}else if(style == "1"){

		where_zapr = fmt.Sprintf("SELECT DISTINCT  id_employee FROM `Employees_of_companies` WHERE id_company = %s", authorized_user.Corporation.Id)
	}else if(style == "3"){

		where_zapr = fmt.Sprintf("SELECT DISTINCT  id_employee FROM `favourites` WHERE id_corporation = %s", authorized_user.Corporation.Id)
	}else{

		where_zapr = fmt.Sprintf("SELECT DISTINCT  id_employee FROM `ask_to_corp` WHERE id_corporation = %s", authorized_user.Corporation.Id)
	}

	var data Seraching_page_data
	data.Style = style
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()


	var zapros = fmt.Sprintf("%s", where_zapr)
	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var user User

	for res.Next(){
		err = res.Scan(&user.Id)
		fmt.Print("This:   ")
		fmt.Println(user.Id)

		user = get_user_employee_by_id(user.Id)
		res_avg := compare_employees(user, r)
		user.Suitable_rating = strconv.FormatFloat(res_avg, 'f', -1, 64)
		data.Employees = append(data.Employees, user)
		fmt.Print("This:   ")
		fmt.Println(user.Id)
	}

	fmt.Println("55")
	fmt.Println(data.Employees)
	t.ExecuteTemplate(w, "search_employees_page", data)


}

func compare_employees (userMayBe User , r *http.Request) (float64){

	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Corporation = get_user_corporation_by_id(authorized_user.Corporation.Id)
	where_zapr := fmt.Sprintf("SELECT DISTINCT  id_employee FROM `Employees_of_companies` WHERE id_company = %s", authorized_user.Corporation.Id)
	var summEmpl float64
	var countEmpl float64
	var zapros = fmt.Sprintf("%s", where_zapr)


	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()



	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var user User

	for res.Next(){
		err = res.Scan(&user.Id)



		user = get_user_employee_by_id(user.Id)
		fmt.Print("This2 :   ")
		fmt.Print(user)
		fmt.Print("     ")
		fmt.Println(userMayBe)
		res, _ := AiCheck(user, userMayBe)
		summEmpl += res
		countEmpl += 1
		fmt.Print("check   ")
		fmt.Println(res)


		fmt.Println("----------------------------------------------")
		fmt.Println()
	}
	if(countEmpl == 0){
		return -1
	}
	return math.Round(summEmpl/countEmpl*100)/100


}


func add_employee_to_your_company(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var current_employee_id = vars["id_employee"]

	fmt.Println()

	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Corporation = get_user_corporation_by_id(authorized_user.Corporation.Id)



	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into naimix.ask_to_employees ( `id_employee`,	`id_company`) values (?, ?)",current_employee_id, authorized_user.Corporation.Id)

	fmt.Println(result)
	if(err != nil){
		fmt.Println(err)
	}


	fmt.Println(authorized_user.Corporation)

	http.Redirect(w, r, "/for_company", http.StatusSeeOther)

}

func asked_to_employee(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("../frontend/templates/asked_to_employee.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()


	var data Search_corp_data
	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Employee = get_user_employee_by_id(authorized_user.Employee.Id)

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
	var style = vars["style"]
	where_zapr := ""
	data.Style = style
	if(style == "0"){
		where_zapr = fmt.Sprintf("SELECT 	id FROM `corporations`")
		//where_zapr = "SELECT id FROM `users` "
	}else if(style == "1"){

		where_zapr = fmt.Sprintf("SELECT DISTINCT 	id_company FROM `ask_to_employees` WHERE id_employee = %s", authorized_user.Employee.Id)
	}



	var zapros = fmt.Sprintf(where_zapr)


	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var corp_now Corporation_data

	for res.Next(){
		err = res.Scan(&corp_now.Id)
		corp_now = get_user_corporation_by_id(corp_now.Id)
		data.Corporations = append(data.Corporations, corp_now)
	}


	t.ExecuteTemplate(w, "asked_to_employee", data)
}


func add_company_to_wanted_company(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var current_company_id = vars["id_company"]

	fmt.Println()

	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Employee = get_user_employee_by_id(authorized_user.Employee.Id)



	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into naimix.ask_to_corp ( `id_employee`,	`id_corporation`) values (?, ?)",authorized_user.Employee.Id, current_company_id)

	fmt.Println(result)
	if(err != nil){
		fmt.Println(err)
	}


	fmt.Println(authorized_user.Corporation)
	http.Redirect(w, r, "/for_recruts", http.StatusSeeOther)


}



func add_or_delete_from_favourites(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
	var corp = vars["id_empl"]

	var rej_apr = vars["rej_apr"]

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Corporation = get_user_corporation_by_id(authorized_user.Corporation.Id)

	query := ""
	if(rej_apr == "1"){
		fmt.Println("C>LFFFLC>LF")
		query = "insert into naimix.favourites ( `id_employee`,	`id_corporation`) values (?, ?)"
	}else{
		fmt.Println("DELETETETE")
		query = "DELETE FROM favourites WHERE id_employee = ? AND id_corporation=?"
	}

	_, _ = db.Exec(query, corp ,authorized_user.Corporation.Id)

  if err != nil {
      log.Fatal(err)
  }

	http.Redirect(w, r, "/", http.StatusSeeOther)
}



func reject_apruvd_ask(w http.ResponseWriter, r *http.Request){


	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
	var corp = vars["id_corp"]

	var rej_apr = vars["rej_apr"]

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Employee = get_user_employee_by_id(authorized_user.Employee.Id)


	if(rej_apr == "1"){
		_, _ = db.Exec("insert into naimix.Employees_of_companies ( `id_employee`,	`id_company`) values (?, ?)",authorized_user.Employee.Id, corp)
	}

	query := "DELETE FROM ask_to_employees WHERE id_company = ? AND id_employee = ?"

  _, err = db.Exec(query, corp, authorized_user.Employee.Id)
  if err != nil {
      log.Fatal(err)
  }

	http.Redirect(w, r, "/", http.StatusSeeOther)
}





func reject_apruvd_ask_to_empl(w http.ResponseWriter, r *http.Request){


	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
	var corp = vars["id_empl"]

	var rej_apr = vars["rej_apr"]

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	authorized_user, _  := populateUserFromSession(r, sessionName)
	authorized_user.Corporation = get_user_corporation_by_id(authorized_user.Corporation.Id)
	fmt.Print("auCORP")
	fmt.Println(authorized_user.Corporation)


	if(rej_apr == "1"){
		_, _ = db.Exec("insert into naimix.Employees_of_companies ( `id_employee`,	`id_company`) values (?, ?)", corp, authorized_user.Corporation.Id)

	}

	query := fmt.Sprintf("DELETE FROM ask_to_corp WHERE id_corporation = %s AND id_employee = %s",  authorized_user.Corporation.Id, corp)


  _, err = db.Exec(query)
  if err != nil {
      log.Fatal(err)
  }

	http.Redirect(w, r, "/", http.StatusSeeOther)
}





func for_company(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/for_company.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }
			authorized_user, _  := populateUserFromSession(r, sessionName)
			//authorized_user.Corporation = get_user_employee_by_id(authorized_user.Corporation.Id)
			if(authorized_user.Corporation.Id == ""){
					http.Redirect(w, r, "/for_company_input", http.StatusFound)
			}
			fmt.Println("for_company_input ")
      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "for_company", nil)
}

func for_recruts(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/for_recruts.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }
			authorized_user, _  := populateUserFromSession(r, sessionName)
			//authorized_user.Corporation = get_user_employee_by_id(authorized_user.Corporation.Id)
			if(authorized_user.Employee.Id == ""){
					http.Redirect(w, r, "/for_recruts_input", http.StatusFound)
			}
			fmt.Println("for_recruts ")
      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "for_recruts", nil)
}



func for_company_input(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/for_company_input.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }

			fmt.Println("for_company_input ")
      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "for_company_input", nil)
}



func for_recruts_input(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/for_recruts_input.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }

			fmt.Println("for_recruts_input ")
      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "for_recruts_input", nil)
}

func GenAvatar(user User) {
	prompt := fmt.Sprintf("Нарисуй абстрактную картинку используя знания из нумерологии"+
		"космограммы и таро, зная такие данные: %s, %s, "+
		"%s года рождения, профессия - %s, %s опыт работы",
		user.Surname, user.Firstname, user.Date_birth_day, user.Speciality, user.Work_experience)

	cmd := exec.Command("python3", "../AiKandinsky/main.py", prompt, user.Id, user.Surname, user.Firstname)
	fmt.Println("9 8 ---- 8743 -")
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error running script: %v", err)
	}

	log.Println("Image generated and saved successfully.")
}

func main() {

	store.Options = &sessions.Options{
        Path:     "/",       // Путь для куков
        MaxAge:   86400 * 7, // Время жизни куков - 7 дней
        HttpOnly: true,      // Запрет доступа к кукам через JavaScript
        Secure:   false,     // Для HTTP должен быть false
    }


 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))
	http.Handle("/static2/", http.StripPrefix("/static2/", http.FileServer(http.Dir("../AiKandinsky/users"))))
	r := mux.NewRouter()
	fmt.Println("Start")



	r.HandleFunc("/", home_page)
	r.HandleFunc("/employee_registration_page", employee_registration_page)
	r.HandleFunc("/employee_authorization_page", employee_authorization_page)
	r.HandleFunc("/employee_profile/{id_employee}", employee_profile)
	r.HandleFunc("/asked_to_employee/{style}", asked_to_employee)
	r.HandleFunc("/reject_apruvd_ask/{id_corp}/{rej_apr}", reject_apruvd_ask)
	r.HandleFunc("/add_company_to_wanted_company/{id_company}", add_company_to_wanted_company)


	r.HandleFunc("/corporation_registration_page", corporation_registration_page)
	r.HandleFunc("/corporation_authorization_page", corporation_authorization_page)
	r.HandleFunc("/corporation_profile/{id_corporation}", corporation_profile)
	r.HandleFunc("/search_employees_page/{style}", search_employees_page)
	r.HandleFunc("/reject_apruvd_ask_to_empl/{id_empl}/{rej_apr}", reject_apruvd_ask_to_empl)
	r.HandleFunc("/add_or_delete_from_favourites/{id_empl}/{rej_apr}", add_or_delete_from_favourites)




	r.HandleFunc("/add_employee_to_your_company/{id_employee}", add_employee_to_your_company)




//	r.HandleFunc("/input_page", input_page)

	r.HandleFunc("/for_recruts", for_recruts)
 	r.HandleFunc("/for_company", for_company)

	r.HandleFunc("/for_recruts_input", for_recruts_input)
	r.HandleFunc("/for_company_input", for_company_input)
//	r.HandleFunc("/check_ai", check_ai)
	http.DefaultTransport.(*http.Transport).MaxIdleConns = 1000
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000

  http.Handle ("/", r)
  http.ListenAndServe(":8080", nil)


}
