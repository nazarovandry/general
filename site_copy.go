package main

import (
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"sync"
	"strings"
	"strconv"

	"os"
	_ "github.com/heroku/x/hmetrics/onload"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie

	if logged {
		w.Write([]byte(`
		<!doctype html>
		<html>
		<body>
		<form action="/action" method="post" class="reg-form">
		<div class="form-row">
			<p>Hi, ` + session.Value + `)  <p>
			<input type="submit" name="action" value="Logout">
			<input type="submit" name="action" value="Change password">
			<input type="submit" name="action" value="Show/hide cards">
		</div>
		</form>
		</body>
		</html>`))
	} else {
		w.Write([]byte(`
		<!doctype html>
		<html>
		<body>
		<form action="/login" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_name">Name: </label>
    		<input type="text" id="form_name" name="name">
  		</div>
		<div class="form-row">
			<label for="form_pw">Password: </label>
			<input type="password" id="form_pw" name="password">
		</div>
		<div class="form-row">
			<input type="submit" value="Oke">
		</div>
		</form>
		</body>
		</html>`))
	}
	if logged && session.Value == "andry" {
		w.Write([]byte(`
		<!doctype html>
		<html>
		<body>
		<form action="/users" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_name">Name: </label>
			<input type="text" id="form_name" name="name">
		</div>
		<div class="form-row">
			<label for="form_name">Value: </label>
    		<input type="text" id="form_value" name="password">
  		</div>
		<div class="form-row">
			<input type="submit" name="but" value="Add man">
			<input type="submit" name="but" value="Change password">
			<input type="submit" name="but" value="Change name">
			<input type="submit" name="but" value="Delete man">
		</div>
		</form>
		</body>
		<body>
		<form action="/addcard" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_cars">Card: </label>
			<input type="text" id="form_card" name="card">
		</div>
		<div class="form-row">
			<label for="form_name">For man: </label>
    		<input type="text" id="form_name" name="name">
			<label for="form_show">Shown: </label>
    		<input type="checkbox" id="form_show" name="shown">
  		</div>
		<div class="form-row">
			<input type="submit" value="Add card">
		</div>
		</form>
		</body>
		<body>
		<form action="/opercard" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_card">Card number: </label>
			<input type="text" id="form_card" name="num">
		</div>
		<div class="form-row">
			<label for="form_name">From man: </label>
    		<input type="text" id="form_name" name="name">
  		</div>
		<div class="form-row">
			<input type="submit" name="card_oper" value="Delete card">
			<input type="submit" name="card_oper" value="Make card shown">
			<input type="submit" name="card_oper" value="Make card hidden">
		</div>
		</form>
		</body>
		<body>
		<form action="/setpictures" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_name">Picture name: </label>
			<input type="text" id="form_name" name="name">
		</div>
		<div class="form-row">
			<label for="form_url">Url: </label>
    		<input type="text" id="form_name" name="url">
  		</div>
		<div class="form-row">
			<input type="submit" name="pic_oper" value="Delete pic">
			<input type="submit" name="pic_oper" value="Create/edit pic">
		</div>
		</form>
		</body>
		<body>
		<form action="/reload" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_list">list.txt </label>
			<textarea rows="3" cols="30" name="list"></textarea>
		</div>
		<div class="form-row">
			<label for="form_cards">cards.txt </label>
			<textarea rows="3" cols="30" name="cards"></textarea>
  		</div>
		<div class="form-row">
			<input type="submit" name="load" value="Reload">
		</div>
		</form>
		</body>
		<body>
		<form action="/download" method="post" class="reg-form">
		<div class="form-row">
			<input type="submit" name="load" value="Download">
		</div>
		</form>
		</body>
		</html>`))
	}
	w.Write([]byte(`
		<!doctype html>
		<html>
		<table border="1">
			<tr>
				<th>Name</th>
				<th>Cards</th>
			</tr>`))
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("list.txt")
	mu.Unlock()
	hiddenUrl := hiddenPic()
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		elems := strings.Split(line, " ")
		w.Write([]byte(
			`<tr>
				<td>` + elems[0] + `</td><td>`))
		if len(elems) > 2 {
			links := strings.Split(elems[2], "&")
			for _, link := range links {
				pic := strings.Split(link, "?")
				if len(pic) > 1 && (pic[1] == "show" ||
					(logged && elems[0] == session.Value) ||
					(logged && session.Value == "andry")){
					w.Write([]byte(`
					<img src=` + pic[0] + `>`))
				} else {
					w.Write([]byte(`
					<img src=` + hiddenUrl + `>`))
				}
			}
		}
		w.Write([]byte(`</td></tr>`))
	}
	w.Write([]byte(`
		</table>
		</p><p>
		</p>&copy;AndrY 2019<p>
		</html>
	`))
}

func hiddenPic() (string) {
	mu := &sync.Mutex{}
	mu.Lock()
	cards, _ := ioutil.ReadFile("cards.txt")
	mu.Unlock()
	array := strings.Split(string(cards), "\n")
	for _, ar := range array {
		first := strings.Split(string(ar), " ")
		if len(first) > 1 && first[0] == "hide" {
			return first[1]
		}
	}
	return ""
}

func setPictures(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie
	if !logged || session.Value != "andry" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	name := r.FormValue("name")
	url := r.FormValue("url")
	button := r.FormValue("pic_oper")
	if wrong(name) || wrong(url) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("cards.txt")
	found := false
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) == 2 && elems[0] == name {
			if button == "Delete pic" {
				lines[i] = ""
			} else {
				found = true
				lines[i] = elems[0] + " " + url
			}
			continue
		}
	}
	if found == false && button == "Create/edit pic" {
		lines = append(lines, name + " " + url)
	}
	finlines := strings.Join(lines, "\n")
	finlines = strings.Replace(finlines, "\n\n", "\n", 1)
	if len(finlines) > 0 && finlines[0] == '\n' {
		finlines = finlines[1:]
	}
	_ = ioutil.WriteFile("cards.txt", []byte(finlines), 0644)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusFound)
}

func getUrl(card string) (string) {
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("cards.txt")
	mu.Unlock()
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) == 2 && elems[0] == card {
			return elems[1]
		}
	}
	return ""
}

func wrong(s string) (bool) {
	return s == "" || strings.Contains(s, " ") ||
		strings.Contains(s, "\n") || strings.Contains(s, "&") ||
		strings.Contains(s, "?") || strings.Contains(s, "\t")
}

func users(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie
	name := r.FormValue("name")
	pass := r.FormValue("password")
	pass2 := r.FormValue("password2")
	button := r.FormValue("but")
	if wrong(pass) || (pass2 != "" && (wrong(pass2) || pass2 != pass)) ||
		button == "No, I dont want" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	log.Println(button)
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("list.txt")
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		elems := strings.Split(line, " ")
		if elems[0] == name ||
			(button == "Oke" && logged && session.Value == elems[0]) {
			if button == "Delete man" && pass == "delete" {
				lines[i] = ""
			} else if button == "Change name" && len(elems) > 1 {
				lines[i] = pass + " " + elems[1]
				if len(elems) > 2 {
					lines[i] += " " + elems[2]
				}
			} else if (button == "Change password" || button == "Oke") &&
				len(elems) > 1 {
				lines[i] = elems[0] + " " + pass
				if len(elems) > 2 {
					lines[i] += " " + elems[2]
				}
			} else if button == "Add man" {
				log.Println("Name is already exists.")
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
			continue
		}
	}
	if button == "Add man" {
		lines = append(lines, name + " " + pass)
	}
	finlines := strings.Join(lines, "\n")
	finlines = strings.Replace(finlines, "\n\n", "\n", 1)
	if len(finlines) > 0 && finlines[0] == '\n' {
		finlines = finlines[1:]
	}
	_ = ioutil.WriteFile("list.txt", []byte(finlines), 0644)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusFound)
}

func operCard(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie
	name := r.FormValue("name")
	num := r.FormValue("num")
	button := r.FormValue("card_oper")
	number := -1
	if name != "" || button == "No, I dont want"{
		numberTmp, err := strconv.Atoi(num)
		if err != nil || button == "No, I dont want" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		number = numberTmp-1
	}
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("list.txt")
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) > 2 && (elems[0] == name ||
			(button == "Oke" && logged && session.Value == elems[0])) {
			newline := ""
			links := strings.Split(elems[2], "&")
			if len(links) > number || number == -1 {
				for j, link := range links {
					if j != number && number != -1 {
						newline += "&" + link
					} else if button == "Make card shown" {
						newline += "&" + strings.Split(link, "?")[0] + "?show"
					} else if button == "Make card hidden" {
						newline += "&" + strings.Split(link, "?")[0] + "?hide"
					} else if button == "Oke" {
						check := r.FormValue("shown" + strconv.Itoa(j))
						newline += "&" + strings.Split(link, "?")[0] +
							"?" + getShow(check)
					}
				}
				lines[i] = elems[0] + " " + elems[1]
				if len(newline) > 0 {
					lines[i] += " " + newline[1:]
				}
			}
		}
	}
	_ = ioutil.WriteFile("list.txt", []byte(strings.Join(lines, "\n")), 0644)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusFound)
}

func getShow (s string) string {
	if s == "on" {
		return "show"
	}
	return "hide"
}

func addCard(w http.ResponseWriter, r *http.Request) {
	card := r.FormValue("card")
	name := r.FormValue("name")
	shown := r.FormValue("shown")
	url := getUrl(card)
	if url == "" || wrong(card) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("list.txt")
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) > 2 && elems[0] == name {
			lines[i] += "&" + getUrl(card) + "?" + getShow(shown)
		} else if len(elems) == 2 && elems[0] == name {
			lines[i] += " " + getUrl(card) + "?" + getShow(shown)
		}
	}
	_ = ioutil.WriteFile("list.txt", []byte(strings.Join(lines, "\n")), 0644)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusFound)
}

func download(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie
	if logged && session.Value == "andry" {
		mu := &sync.Mutex{}
		mu.Lock()
		list, _ := ioutil.ReadFile("list.txt")
		cards, _ := ioutil.ReadFile("cards.txt")
		mu.Unlock()
		w.Write([]byte(list))
		w.Write([]byte("===="))
		w.Write([]byte(cards))
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func reload(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie
	if !logged || session.Value != "andry" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	list := r.FormValue("list")
	cards := r.FormValue("cards")
	mu := &sync.Mutex{}
	mu.Lock()
	list = strings.Replace(list, "\r", "", -1)
	cards = strings.Replace(cards, "\r", "", -1)
	_ = ioutil.WriteFile("list.txt", []byte(list), 0644)
	_ = ioutil.WriteFile("cards.txt", []byte(cards), 0644)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusFound)
}

func allRight(name string, pass string) (bool) {
	if wrong(name) || wrong(pass) {
		return false
	}
	mu := &sync.Mutex{}
	mu.Lock()
	data, _ := ioutil.ReadFile("list.txt")
	mu.Unlock()
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		elem := strings.Split(l, " ")
		if len(elem) >= 2 && elem[0] == name {
			return elem[1] == pass
		}
	}
	return false
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(10 * time.Hour)
	name := r.FormValue("name")
	pass := r.FormValue("password")
	if allRight(name, pass) {
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   name,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func actionPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	logged := err != http.ErrNoCookie
	button := r.FormValue("action")
	if button == "Logout" {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else if button == "Show/hide cards" {
		w.Write([]byte(`
		<!doctype html>
		<html>
		<form action="/opercard" method="post" class="reg-form">
		<table border="1">
			<tr>
				<th>Card</th>
				<th>Shown</th>
			</tr>`))
		mu := &sync.Mutex{}
		mu.Lock()
		data, _ := ioutil.ReadFile("list.txt")
		mu.Unlock()
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			elems := strings.Split(line, " ")
			if len(elems) > 2 && logged && elems[0] == session.Value {
				links := strings.Split(elems[2], "&")
				for i, link := range links {
					pic := strings.Split(link, "?")
					if len(pic) > 1 {
						w.Write([]byte(`<tr>
							<td><img src=` + pic[0] + `></td>
							<td><div class="form-row">`))
						is := strconv.Itoa(i)
						if pic[1] == "hide" {
							w.Write([]byte(`
								<input type="checkbox" id="check` + is +
									 `" name="shown` + is + `">`))
						} else {
							w.Write([]byte(`
								<input type="checkbox" id="check` + is +
									`" name="shown` +
									is + `" checked="checked">`))
						}
						w.Write([]byte(`</div></td></tr>`))
					}
				}
			}
		}
		w.Write([]byte(`
		</table>
		<div class="form-row">
			<input type="submit" name="card_oper" value="Oke">
			<input type="submit" name="card_oper" value="No, I dont want">
		</div>
		</form>
		</html>`))
		return
	}
	// else "Change Password"
	w.Write([]byte(`
		<!doctype html>
		<html>
		<body>
		<form action="/users" method="post" class="reg-form">
		<div class="form-row">
			<label for="form_passnew">New password: </label>
    		<input type="password" id="form_passnew" name="password">
  		</div>
		<div class="form-row">
			<label for="form_passnew2">New password again: </label>
    		<input type="password" id="form_passnew2" name="password2">
  		</div>
		<div class="form-row">
			<input type="submit" name="but" value="Oke">
			<input type="submit" name="but" value="No, I dont want">
		</div>
		</form>
		</body>
		</html>
	`))
}

func main() {
	list, _ := ioutil.ReadFile("list.txt")
	cards, _ := ioutil.ReadFile("cards.txt")
	_ = ioutil.WriteFile("list.txt",
		[]byte(strings.Replace(string(list), "\r", "", -1)), 0644)
	_ = ioutil.WriteFile("cards.txt",
		[]byte(strings.Replace(string(cards), "\r", "", -1)), 0644)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/action", actionPage)
	http.HandleFunc("/addcard", addCard)
	http.HandleFunc("/opercard", operCard)
	http.HandleFunc("/users", users)
	http.HandleFunc("/setpictures", setPictures)
	http.HandleFunc("/reload", reload)
	http.HandleFunc("/download", download)
	http.HandleFunc("/", mainPage)

	log.Println("starting server at :8080")
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":"+port, nil)
}
