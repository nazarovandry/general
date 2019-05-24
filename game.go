package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
)

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные перменные ( тут можно ), функции
*/

type World struct {
    where   *Room
    rooms   [](*Room)
    things  [](*Thing)
    haveBag bool
    bag     *Place
    aims    [](*Aim)
}

type Aim struct {
    toGo    *Room
    infGo   string
    toHave  [](*Thing)
    infHave string
    badHave string
    toDo    *Thing
    infDo   string
    badDo   string
    status  bool
}

type Room struct {
    name    string
    info    string
    hello   string
    enter   string
    empty   string
    places  [](*Place)
    sayAim  bool
    jump    bool
    toGo    [](*Room)
}

type Place struct {
    name    string
    things  [](*Thing)
    bag     bool
}

type Thing struct {
    name    string
    secret  bool
    mobile  bool
    taken   bool
    used    bool
    useFor  [](*Thing)
    okInfo  string
}

func (p Place) nothingHere() (bool) {
    for _, th := range p.things {
        if !th.taken && !th.secret {
            return false
        }
    }
    if p.bag {
        return false
    }
    return true
}

func (r Room) nothing() (bool) {
    for _, pl := range r.places {
        if !pl.nothingHere() {
            return false
        }
    }
    return true
}

func (r Room) aboutPlaces() (out string) {
    flagComma := false
    for plIndex, pl := range r.places {
        if pl.nothingHere() {
            continue
        }
        flagComma = true
        if flagComma && plIndex > 0 {
            out += ", "
        }
        out += (pl.name + ": ")
        flagComma = false
        for ind, th := range pl.things {
            if !th.taken && !th.secret{
                if flagComma && ind > 0 {
                    out += ", "
                }
                out += th.name
                flagComma = true
            }
        }
        if pl.bag {
            out += "рюкзак"
        }
    }
    if r.nothing() {
        if r.empty == "" {
            out += "здесь ничего нет"
        } else {
            out += r.empty
        }
    }
    return
}

func (r Room) whereToGo() (out string) {
    if r.toGo != nil {
        out += ". можно пройти - "
        for ind, to := range r.toGo {
            if ind > 0 {
                out += ", "
            }
            out += to.name
        }
    }
    return
}

func (ai Aim) CheckStatusSee(invThings [](*Thing), out *string) {
    *out += "надо "
    for _, th := range ai.toHave {
        flag := true
        for _, your := range invThings {
            if (your == th) {
                flag = false
                continue
            }
        }
        if flag {
            if ai.infHave != "" {
                *out += (ai.infHave + " и ")
            }
        }
    }
    if ai.toDo != nil {
        if !ai.toDo.used {
            if ai.infDo != "" {
                *out += ai.infDo
                if ai.toGo != nil {
                    *out += ", затем "
                }
            }
        }
    }
    if ai.toGo != nil {
        *out += ai.infGo
    }
}

func (ai Aim) CheckStatusGo(invThings [](*Thing), out *string) (bool) {
    for _, th := range ai.toHave {
        flag := true
        for _, your := range invThings {
            if your == th {
                flag = false
                continue
            }
        }
        if flag {
            *out = ai.badHave
            return false
        }
    }
    if ai.toDo != nil {
        if !ai.toDo.used {
            *out = ai.badDo
            return false
        }
    }
    return true
}

func (y *World) UpdateStatus() {
    for _, ai := range y.aims {
        if !ai.status {
            out := ""
            if ai.CheckStatusGo(y.things, &out) {
                ai.status = true
            }
        }
    }
}

func (y World) tooMany() (string) {
    return "слишком много аргументов"
}

func (y World) See(arg []string) (out string) {
    if len(arg) > 1 {
        out = y.tooMany()
        return
    }
    if y.where.info != "" {
        out = (y.where.info + ", ")
    }
    out += y.where.aboutPlaces()
    if y.where.sayAim {
        flagAllDone := true
        for _, ai := range y.aims {
            if !ai.status {
                flagAllDone = false
                if (out != "") {
                    out += ", "
                }
                ai.CheckStatusSee(y.things, &out)
                break
            }
        }
        if flagAllDone {
            if (out != "") {
                out += ", "
            }
            out += "вы выполнили все задания"
        }
    }
    out += y.where.whereToGo()
    return
}

func (y *World) SearchRoom(arg string, out *string) (bool) {
    for _, to := range y.where.toGo {
        if arg == to.name {
            for _, ai := range y.aims {
                if ai.toGo == to {
                    if !ai.status {
                        if !ai.CheckStatusGo(y.things, out) {
                            return true
                        }
                        ai.status = true
                        break
                    }
                }
            }
            if to.jump {
                if to.toGo == nil {
                    *out = ("нет входа")
                    return true
                }
                to = (to.toGo)[0]
            }
            y.where = to
            *out = (y.where.hello + y.where.whereToGo())
            return true
        }
    }
    return false
}


func (y *World) Go(arg []string) (out string) {
    if len(arg) == 1 {
        out = "идти куда? выбирай"
        out += y.where.whereToGo()
    } else if len(arg) > 2 {
        out = y.tooMany()
        return
    } else {
        if arg[1] == y.where.name {
            out = "вы и так здесь"
            return
        }
        if y.SearchRoom(arg[1], &out) {
            return
        }
        for _, ro := range y.rooms {
            if ro.name == arg[1] {
                out = ("нет пути в " + ro.name)
                return
            }
        }
        out = "нет такого места"
    }
    return
}

func (y *World) Wear(arg []string) (out string) {
    if len(arg) == 1 {
        out = "надеть что? осмотритесь сначала"
        return
    }
    if len(arg) > 2 {
        out = y.tooMany()
        return
    }
    for _, pl := range y.where.places {
        if pl.bag == true {
            if arg[1] == "рюкзак" {
                out = "вы надели: рюкзак"
                pl.bag = false
                y.haveBag = true
                return
            }
        }
    }
    out = "нет такого"
    return
}

func (y *World) Take(arg []string) (out string) {
    if len(arg) == 1 {
        out = "взять что? осмотритесь сначала"
        return
    }
    if len(arg) > 2 {
        out = y.tooMany()
        return
    }
    for _, pl := range y.where.places {
        for _, th := range pl.things {
            if arg[1] == th.name && !th.taken {
                if !th.mobile {
                    out += ("вы не можете взять " + th.name)
                    return
                }
                if y.haveBag {
                    out += ("предмет добавлен в инвентарь: " + th.name)
                    y.UpdateStatus()
                    th.taken = true
                    y.things = append(y.things, th)
                    return
                }
                out = "некуда класть"
                return
            }
        }
    }
    out = "нет такого"
    return
}

func (y *World) UsingOperation(th *Thing, arg2 string) (string) {
    for _, pl := range y.where.places {
        for _, to := range pl.things {
            if arg2 == to.name {
                if th.used {
                    return "уже использовано"
                }
                for _, us := range th.useFor {
                    if us == to {
                        to.used = true
                        y.UpdateStatus()
                        return th.okInfo
                    }
                }
                return "нельзя применить"
            }
        }
    }
    return "не к чему применить"
}


func (y *World) Use(arg []string) (out string) {
    if len(arg) == 1 {
        out = "применить что? посмотрите инвентарь"
        return
    }
    if len(arg) > 3 {
        out = y.tooMany()
        return
    }
    for _, th := range y.things {
        if arg[1] == th.name {
            if len(arg) == 3 {
                if th.used {
                    out = "уже использовано"
                    return
                }
                return y.UsingOperation(th, arg[2])
            }
            out = ("вы не указали к чему вы применяете " + arg[1])
            return
        }
    }
    out = ("нет предмета в инвентаре - " + arg[1])
    return
}

func (y World) Inv(arg []string) (out string) {
    if len(arg) > 1 {
        out = y.tooMany()
        return
    }
    for ind, th := range y.things {
        if ind > 0 {
            out += ", "
        } else {
            out = "у вас есть "
        }
        out += th.name
    }
    if len(y.things) == 0 {
        out = "у вас ничего нет"
    }
    return
}

var you World

func main() {
	/*
		в этой функции можно ничего не писать
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
    initGame()
    for {
        reader := bufio.NewReader(os.Stdin)
        s, _ := reader.ReadString('\n')
        answer := handleCommand(s[:len(s)-1])
        if answer == "@" {
            return
        }
        fmt.Println(answer)
    }
    return
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все команты
		если что-то было - оно корректно перезатирается
	*/

    tea := Thing{
        name:   "чай",
    }

    door := Thing{
        name:   "дверь",
        secret: true,
    }

    konsp := Thing{
        name:   "конспекты",
        mobile: true,
    }

    keys := Thing{
        name:   "ключи",
        mobile: true,
        useFor: [](*Thing){&door},
        okInfo: "дверь открыта",
    }

    spoon := Thing{
        name:   "ложка",
        secret: true,
        mobile: true,
        useFor: [](*Thing){&tea},
        okInfo: "чай перемешан",
    }

    chair := Place{
        name:   "на стуле",
        bag:    true,
    }

    kitchen := Room{
        name:   "кухня",
        info:   "ты находишься на кухне",
        hello:  "кухня, ничего интересного",
        empty:  "пустая комната",
        places: [](*Place){
            &Place{
                name:   "на столе",
                things: [](*Thing){&tea},
            },
        },
        sayAim: true,
    }

    hall := Room{
        name:   "коридор",
        info:   "ты в коридоре",
        hello:  "ничего интересного",
        empty:  "пустая комната",
        places: [](*Place){
            &Place{
                name:   "вдоль стены",
                things: [](*Thing){&door},
            },
        },
    }

    room := Room{
        name:   "комната",
        hello:  "ты в своей комнате",
        empty:  "пустая комната",
        places: [](*Place){
            &Place{
                name:   "на столе",
                things: [](*Thing){&keys, &konsp},
            },
            &chair,
        },
    }

    out := Room{
        name:   "улица",
        info:   "ты на улице",
        hello:  "на улице весна",
        places: [](*Place){
            &Place{
                name:   "на земле",
                things: [](*Thing){&spoon},
            },
        },
    }

    home := Room{
        name:   "домой",
        jump:   true,
        toGo:   [](*Room){&hall},
    }

    kitchen.toGo = append(kitchen.toGo, &hall)
    hall.toGo = append(hall.toGo, &kitchen, &room, &out)
    room.toGo = append(room.toGo, &hall)
    out.toGo = append(out.toGo, &home)

    walk := Aim{
        toGo:       &out,
        infGo:      "идти в универ",
        toHave:     [](*Thing){&konsp},
        infHave:    "собрать рюкзак",
        badHave:    "вы еще не собрали рюкзак",
        toDo:       &door,
        badDo:      "дверь закрыта",
    }

    drink := Aim{
        toDo:   &tea,
        infDo:  "перемешать чай при помощи ложки",
    }

    you = World{
        where:  &kitchen,
        rooms:  [](*Room){&kitchen, &hall, &room, &out},
        bag:    &chair,
        aims:   [](*Aim){&walk, &drink},
    }
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
    list := strings.Split(command, " ")
    if len(list) >= 1 {
        switch list[0] {
        case "осмотреться", "о":
            return you.See(list)
	    case "идти", "и":
            return you.Go(list)
	    case "надеть", "н":
            return you.Wear(list)
	    case "взять", "в":
            return you.Take(list)
	    case "применить", "п":
            return you.Use(list)
        case "инвентарь", "ин":
            return you.Inv(list)
        case "выход", "exit", "x", "escape":
            return "@"
        default:
            return "неизвестная команда"
        }
    }
    return ""
}
