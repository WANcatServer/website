package main

import (
	"database/sql"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

func getUser(w http.ResponseWriter, r *http.Request) *User {
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return nil
	}
	session := cookie.Value
	q := `
	SELECT U.ID
	FROM Sessions S, Users U
	WHERE S.ID = ? AND U.ID = S.User
	`
	raw := db.QueryRow(q, session)
	user := new(User)
	var userID int
	err = raw.Scan(&userID)
	if err == sql.ErrNoRows {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		return nil
	} else if err != nil {
		log.Fatal("raw scan fatal: ", err)
	}
	user.ID = userID
	return user
}

func join(w http.ResponseWriter, r *http.Request) {
	if getUser(w, r) != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if r.Method != "POST" {
		err := tpl.ExecuteTemplate(w, "join.html", "")
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	username := r.FormValue("uname")
	email := r.FormValue("email")
	password := r.FormValue("psw")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Password generate error: ", err)
	}
	result, err := db.Exec(`
	INSERT INTO Users (Name, Email, Password, JoinDate)
	VALUES (?, ?, ?, CURRENT_TIMESTAMP());
	`, username, email, hashPassword)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	userID, _ := result.LastInsertId()
	sessionUUID := uuid.Must(uuid.NewV4())
	session := sessionUUID.String()
	_, err = db.Exec(`
	INSERT INTO Sessions 
	VALUES (?, ?, CURRENT_TIMESTAMP())	
	`, session, userID)
	if err != nil {
		log.Fatalln("Insert new session fatal: ", err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: session,
	})
	http.Redirect(w, r, "/dashboard", 303)
}

func login(w http.ResponseWriter, r *http.Request) {
	if getUser(w, r) != nil {
		// already login
		http.Redirect(w, r, "/", 303)
		return
	}
	if r.Method != "POST" {
		tpl.ExecuteTemplate(w, "login.html", "")
		return
	}
	uname := r.FormValue("uname")
	password := r.FormValue("psw")
	q := `
	SELECT ID, Password
	FROM Users
	WHERE Email = ? OR Name = ?
	`
	row := db.QueryRow(q, uname, uname)
	var id int
	var hashPassword []byte
	err := row.Scan(&id, &hashPassword)
	if err == sql.ErrNoRows {
		log.Println("Email is not found")
		tpl.ExecuteTemplate(w, "login.html", "Fatal: Email or password is uncorrect")
		return
	}
	if bcrypt.CompareHashAndPassword(hashPassword, []byte(password)) != nil {
		log.Println("Password is not correct")
		tpl.ExecuteTemplate(w, "login.html", "Fatal: Email or password is uncorrect")
		return
	}
	sessionUUID, err := uuid.NewV4()
	session := sessionUUID.String()
	if err != nil {
		log.Fatalln("uuid init fatal: ", err)
	}
	stmt, err := db.Prepare(`
	INSERT INTO Sessions
	VALUES (?, ?, CURRENT_TIMESTAMP())
	`)
	if err != nil {
		log.Fatalln("db prepare fatal: ", err)
	}
	result, err := stmt.Exec(session, id)
	if r, _ := result.RowsAffected(); err != nil || r != 1 {
		log.Fatalln("Database insert fatal: ", err, result)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: session,
	})
	http.Redirect(w, r, "/dashboard", 303)
}

func logout(w http.ResponseWriter, r *http.Request) {
	user := getUser(w, r)
	if user == nil { // Didn't login
		http.Redirect(w, r, "/", 303)
	}
	cookie, _ := r.Cookie("session")
	session := cookie.Value
	stmt, err := db.Prepare(`
	DELETE FROM Sessions WHERE ID = ?
	`)
	if err != nil {
		log.Fatalln("db prepare fatal: ", err)
	}
	result, err := stmt.Exec(session)
	if err != nil {
		log.Fatal(err)
	} else {
		raws, _ := result.RowsAffected()
		if raws != 1 {
			log.Fatalln("Effect row not only one: ", result)
		}
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 303)
}
