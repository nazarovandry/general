package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Set struct {
	Gl    []map[string]string
	Such  []map[string]string
	Pril  []map[string]string
	Mest  map[string]map[string]string
	Nar   []map[string]string
	Srav  [][]string
	Predl [][]string
	Link  map[string][]string
}

type Tree struct {
	elem   string
	pars   map[string]*string
	points map[string]*Tree
	resp   map[*string]*[]*Tree
	links  []*Tree
	back   *Tree
	accept bool
	done   bool
	hero   Set
	heroL  Set
}

func many(s string, rod string) string {
	arr := []string{"г", "ж", "к", "х", "ч", "ш", "щ"}
	if rod == "ж" || rod == "м" {
		for _, i := range arr {
			if strings.HasSuffix(s, i) {
				return "и"
			}
		}
		return "ы"
	}
	if strings.HasSuffix(s, "ак") {
		return "а"
	}
	for _, i := range arr {
		if strings.HasSuffix(s, i) {
			return "и"
		}
	}
	return "а"
}

func createMO(s string, par string) map[string]string {
	if strings.HasSuffix(s, "ий") {
		ret := strings.TrimSuffix(s, "ий")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "ия",
			"дат":  ret + "ию",
			"вин":  ret + "ия",
			"тв":   ret + "ием",
			"о":    ret + "ии",
			"мим":  ret + "ии",
			"мрод": ret + "иев",
			"мдат": ret + "иям",
			"мвин": ret + "иев",
			"мтв":  ret + "иями",
			"мо":   ret + "иях",
		}
	} else if strings.HasSuffix(s, "а") {
		ret := strings.TrimSuffix(s, "а")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "ы",
			"дат":  ret + "е",
			"вин":  ret + "у",
			"тв":   ret + "ой",
			"о":    ret + "е",
			"мим":  ret + many(ret, "м"),
			"мрод": ret,
			"мдат": ret + "ам",
			"мвин": ret,
			"мтв":  ret + "ами",
			"мо":   ret + "ах",
		}
	} else if strings.HasSuffix(s, "я") {
		ret := strings.TrimSuffix(s, "я")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "и",
			"дат":  ret + "е",
			"вин":  ret + "ю",
			"тв":   ret + "ей",
			"о":    ret + "е",
			"мим":  ret + "и",
			"мрод": ret + "ь",
			"мдат": ret + "ям",
			"мвин": ret + "ь",
			"мтв":  ret + "ями",
			"мо":   ret + "ях",
		}
	} else if strings.HasSuffix(s, "ь") {
		ret := strings.TrimSuffix(s, "ь")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "я",
			"дат":  ret + "ю",
			"вин":  ret + "я",
			"тв":   ret + "ем",
			"о":    ret + "е",
			"мим":  ret + "и",
			"мрод": ret + "ей",
			"мдат": ret + "ям",
			"мвин": ret + "ей",
			"мтв":  ret + "ями",
			"мо":   ret + "ях",
		}
	} else if strings.HasSuffix(s, "уч") {
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  s + "а",
			"дат":  s + "у",
			"вин":  s + "а",
			"тв":   s + "ем",
			"о":    s + "е",
			"мим":  s + "и",
			"мрод": s + "ей",
			"мдат": s + "ам",
			"мвин": s + "ей",
			"мтв":  s + "ами",
			"мо":   s + "ах",
		}
	}
	return map[string]string{
		"par":  par,
		"им":   s,
		"род":  s + "а",
		"дат":  s + "у",
		"вин":  s + "а",
		"тв":   s + "ом",
		"о":    s + "е",
		"мим":  s + many(s, "м"),
		"мрод": s + "ов",
		"мдат": s + "ам",
		"мвин": s + "ов",
		"мтв":  s + "ами",
		"мо":   s + "ах",
	}
}

func createMN(s string, par string) map[string]string {
	ret := createMO(s, par)
	ret["вин"] = ret["им"]
	ret["мвин"] = ret["мим"]
	return ret
}

func createJO(s string, par string) map[string]string {
	if strings.HasSuffix(s, "ия") {
		ret := strings.TrimSuffix(s, "ия")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "ии",
			"дат":  ret + "ие",
			"вин":  ret + "ию",
			"тв":   ret + "ией",
			"о":    ret + "ии",
			"мим":  ret + "ии",
			"мрод": ret + "ий",
			"мдат": ret + "иям",
			"мвин": ret + "ий",
			"мтв":  ret + "иями",
			"мо":   ret + "иях",
		}
	} else if strings.HasSuffix(s, "а") {
		ret := strings.TrimSuffix(s, "а")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "ы",
			"дат":  ret + "е",
			"вин":  ret + "у",
			"тв":   ret + "ой",
			"о":    ret + "е",
			"мим":  ret + many(ret, "ж"),
			"мрод": ret,
			"мдат": ret + "ам",
			"мвин": ret,
			"мтв":  ret + "ами",
			"мо":   ret + "ах",
		}
	} else if strings.HasSuffix(s, "я") {
		ret := strings.TrimSuffix(s, "я")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "и",
			"дат":  ret + "е",
			"вин":  ret + "ю",
			"тв":   ret + "ей",
			"о":    ret + "е",
			"мим":  ret + "и",
			"мрод": ret + "ь",
			"мдат": ret + "ям",
			"мвин": ret + "ь",
			"мтв":  ret + "ями",
			"мо":   ret + "ях",
		}
	} else if strings.HasSuffix(s, "ь") {
		ret := strings.TrimSuffix(s, "ь")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "и",
			"дат":  ret + "и",
			"вин":  s,
			"тв":   ret + "ью",
			"о":    ret + "и",
			"мим":  ret + "и",
			"мрод": ret + "ей",
			"мдат": ret + "ям",
			"мвин": ret + "ей",
			"мтв":  ret + "ью",
			"мо":   ret + "и",
		}
	}
	return map[string]string{
		"par":  par,
		"им":   s,
		"род":  s,
		"дат":  s,
		"вин":  s,
		"тв":   s,
		"о":    s,
		"мим":  s,
		"мрод": s,
		"мдат": s,
		"мвин": s,
		"мтв":  s,
		"мо":   s,
	}
}

func createJN(s string, par string) map[string]string {
	return createJO(s, par)
}

func createS(s string, par string) map[string]string {
	if strings.HasSuffix(s, "е") {
		ret := strings.TrimSuffix(s, "е")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "я",
			"дат":  ret + "ю",
			"вин":  s,
			"тв":   ret + "ем",
			"о":    ret + "е",
			"мим":  ret + "я",
			"мрод": ret + "ей",
			"мдат": ret + "ям",
			"мвин": ret + "ей",
			"мтв":  ret + "ями",
			"мо":   ret + "ях",
		}
	} else if strings.HasSuffix(s, "я") {
		ret := strings.TrimSuffix(s, "я")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "ени",
			"дат":  ret + "ени",
			"вин":  s,
			"тв":   ret + "енем",
			"о":    ret + "ени",
			"мим":  ret + "ена",
			"мрод": ret + "ен",
			"мдат": ret + "енам",
			"мвин": ret + "ена",
			"мтв":  ret + "енами",
			"мо":   ret + "енах",
		}
	} else if strings.HasSuffix(s, "о") {
		ret := strings.TrimSuffix(s, "о")
		return map[string]string{
			"par":  par,
			"им":   s,
			"род":  ret + "а",
			"дат":  ret + "у",
			"вин":  s,
			"тв":   ret + "ом",
			"о":    ret + many(ret, "ср"),
			"мим":  ret + "а",
			"мрод": ret,
			"мдат": ret + "ам",
			"мвин": ret + "а",
			"мтв":  ret + "ами",
			"мо":   ret + "ах",
		}
	}
	return map[string]string{
		"par":  par,
		"им":   s,
		"род":  s,
		"дат":  s,
		"вин":  s,
		"тв":   s,
		"о":    s,
		"мим":  s,
		"мрод": s,
		"мдат": s,
		"мвин": s,
		"мтв":  s,
		"мо":   s,
	}
}

func createG(pr string, sov string, par string) map[string]string {
	if strings.HasSuffix(pr, "сть") {
		retPr := strings.TrimSuffix(pr, "сть")
		retSov := strings.TrimSuffix(sov, "сть")
		return map[string]string{
			"par":          par,
			"инф":          pr,
			"проц:прош:м":  retPr + "л",
			"проц:прош:ж":  retPr + "ла",
			"проц:прош:мн": retPr + "ли",
			"проц:прош:ср": retPr + "ло",
			"проц:я":       retPr + "м",
			"проц:мы":      retPr + "дим",
			"проц:ты":      retPr + "шь",
			"проц:вы":      retPr + "дите",
			"проц:он":      retPr + "ст",
			"проц:они":     retPr + "дят",
			"проц:дееп":    retPr + "дя",
			"сов:прош:м":   retSov + "л",
			"сов:прош:ж":   retSov + "ла",
			"сов:прош:мн":  retSov + "ли",
			"сов:прош:ср":  retSov + "ло",
			"сов:инф":      sov,
			"сов:буд:ед":   retSov + "ст",
			"сов:буд:мн":   retSov + "дят",
			"сов:буд:я":    retSov + "м",
			"сов:буд:мы":   retSov + "им",
			"сов:буд:ты":   retSov + "шь",
			"сов:буд:вы":   retSov + "дите",
			"сов:дееп":     retSov + "в",
		}
	}
	retPr := strings.TrimSuffix(pr, "ть")
	retSov := strings.TrimSuffix(sov, "ть")
	return map[string]string{
		"par":          par,
		"инф":          pr,
		"проц:прош:м":  retPr + "л",
		"проц:прош:ж":  retPr + "ла",
		"проц:прош:мн": retPr + "ли",
		"проц:прош:ср": retPr + "ло",
		"проц:я":       retPr + "ю",
		"проц:мы":      retPr + "ем",
		"проц:ты":      retPr + "ешь",
		"проц:вы":      retPr + "ете",
		"проц:он":      retPr + "ет",
		"проц:они":     retPr + "ют",
		"проц:дееп":    retPr + "я",
		"сов:прош:м":   retSov + "л",
		"сов:прош:ж":   retSov + "ла",
		"сов:прош:мн":  retSov + "ли",
		"сов:прош:ср":  retSov + "ло",
		"сов:инф":      sov,
		"сов:буд:ед":   retSov + "ит",
		"сов:буд:мн":   retSov + "ют",
		"сов:буд:я":    retSov + "ю",
		"сов:буд:мы":   retSov + "ем",
		"сов:буд:ты":   retSov + "шь",
		"сов:буд:вы":   retSov + "те",
		"сов:дееп":     retSov + "в",
	}
}

func createP(s string, par string) map[string]string {
	if strings.HasSuffix(s, "ий") {
		ret := strings.TrimSuffix(s, "ий")
		return map[string]string{
			"par":     par,
			"мим":     s,
			"мрод":    ret + "ем",
			"мдат":    ret + "ему",
			"мвин":    ret + "его",
			"мвинне":  s,
			"мтв":     ret + "им",
			"мо":      ret + "ем",
			"жим":     ret + "яя",
			"жрод":    ret + "юю",
			"ждат":    ret + "ей",
			"жвин":    ret + "юю",
			"жтв":     ret + "ей",
			"жо":      ret + "ей",
			"срим":    ret + "ее",
			"сррод":   ret + "его",
			"срдат":   ret + "ему",
			"срвин":   ret + "ее",
			"сртв":    ret + "им",
			"сро":     ret + "ем",
			"мним":    ret + "ие",
			"мнрод":   ret + "их",
			"мндат":   ret + "им",
			"мнвин":   ret + "их",
			"мнвинне": ret + "ие",
			"мнтв":    ret + "ими",
			"мно":     ret + "их",
		}
	}
	tmp := strings.TrimSuffix(s, "ый")
	ret := strings.TrimSuffix(tmp, "ой")
	return map[string]string{
		"par":     par,
		"мим":     s,
		"мрод":    ret + "ом",
		"мдат":    ret + "ому",
		"мвин":    ret + "ого",
		"мвинне":  s,
		"мтв":     ret + "м",
		"мо":      ret + "ом",
		"жим":     ret + "ая",
		"жрод":    ret + "ую",
		"ждат":    ret + "ой",
		"жвин":    ret + "ую",
		"жтв":     ret + "ой",
		"жо":      ret + "ой",
		"срим":    ret + "ое",
		"сррод":   ret + "ого",
		"срдат":   ret + "ому",
		"срвин":   ret + "ое",
		"сртв":    ret + "м",
		"сро":     ret + "ом",
		"мним":    ret + "ые",
		"мнрод":   ret + "ых",
		"мндат":   ret + "ым",
		"мнвин":   ret + "ых",
		"мнвинне": ret + "ые",
		"мнтв":    ret + "ыми",
		"мно":     ret + "ых",
	}
}

func createN(s string, par string) map[string]string {
	if strings.HasSuffix(s, "о") {
		ret := strings.TrimSuffix(s, "о")
		return map[string]string{
			"par": par,
			"об":  s,
			"ср":  ret + "ее",
		}
	}
	return map[string]string{
		"par": par,
		"об":  s,
	}
}

func random(a int) int {
	t := rand.NewSource(time.Now().UnixNano())
	r := rand.New(t)
	return r.Intn(a)
}

func randomArr(a []int) int {
	t := rand.NewSource(time.Now().UnixNano())
	r := rand.New(t)
	return a[r.Intn(len(a))]
}

func toHigh(s string, c string) string {
	arr := strings.Split(s, c)
	for ind, elem := range arr {
		if ind != 0 {
			pair := strings.SplitAfterN(elem, "", 2)
			if len(pair) > 0 {
				pair[0] = strings.ToUpper(pair[0])
			}
			arr[ind] = strings.Join(pair, "")
		}
	}
	return strings.Join(arr, c)
}

func printCompact(s string, n int) {
	blocks := strings.Split(s, "\n")
	for _, b := range blocks {
		arr := strings.Split(b, " ")
		for i := range arr {
			if ((i + 1) % n) == 0 {
				fmt.Println(arr[i] + " ")
			} else {
				fmt.Print(arr[i] + " ")
			}
		}
		fmt.Println()
	}
}

func check(s string) string {
	res := s
	res = strings.Replace(res, "  ", " ", -1)
	res = strings.Replace(res, " *", "*", -1)
	res = strings.Replace(res, "* ", "*", -1)
	res = strings.Replace(res, "*", "", -1)
	res = strings.Replace(res, " .", ".", -1)
	res = strings.Replace(res, " ?", "?", -1)
	res = strings.Replace(res, " !", "!", -1)
	res = strings.Replace(res, " ,", ",", -1)
	for {
		res = strings.Replace(res, ",,", ",", -1)
		if !strings.Contains(res, ",,") {
			break
		}
	}
	res = strings.Replace(res, ",?", "?", -1)
	res = strings.Replace(res, ",!", "!", -1)
	res = strings.Replace(res, ",.", ".", -1)
	res = strings.Replace(res, "жы", "жи", -1)
	res = strings.Replace(res, "шы", "ши", -1)
	res = strings.Replace(res, "чя", "ча", -1)
	res = strings.Replace(res, "щя", "ща", -1)
	res = toHigh(res, ". ")
	res = toHigh(res, ".\n")
	res = toHigh(res, "\t")
	res = toHigh(res, "\t ")
	res = toHigh(res, "\t- ")
	return res
}

func split(s string, c string) []string {
	arr := []string{}
	limit := len(c)
	cIter := 0
	stack := 0
	last := 0
	for sIter, i := range s {
		if stack == 0 && s[sIter] == c[cIter] {
			cIter++
			if cIter == limit {
				cIter = 0
				arr = append(arr, s[last:sIter+1-limit])
				last = sIter + 1
			}
		} else if i == '(' {
			stack++
		} else if i == ')' {
			stack--
			if stack < 0 {
				fmt.Printf("Stack error!")
				return []string{}
			}
		}
	}
	arr = append(arr, s[last:len(s)])
	return arr
}

func encode(theRule string, t *Tree, rule *map[string]string, def *map[string]*string, deep int) {
	cases := split(theRule, "|")
	follow := split(cases[random(len(cases))], "+")
	for _, oper := range follow {
		if oper == "" {
			continue
		}
		if len(oper) > 1 && oper[0] == '(' && oper[len(oper)-1] == ')' {
			without := strings.TrimSuffix(strings.TrimPrefix(oper, "("), ")")
			encode(without, t, rule, def, deep)
			continue
		}
		features := split(oper, ":")
		pars := make(map[string]*string)
		for key, value := range *def {
			tmp := *value
			pars[key] = &tmp
		}
		if len(features) == 2 {
			list := split(features[1], ".")
			for _, parameter := range list {
				sides := split(parameter, "=")
				if len(sides) == 1 {
					fmt.Println("Error: bad assign " + parameter)
					continue
				}
				pars[sides[0]] = &sides[1]
			}
		}
		newTree := makeTree(rule, features[0], &pars, def, deep+1)
		(*t).links = append((*t).links, newTree)
		(*(*t).links[len(t.links)-1]).back = t
		//fmt.Println("FIRST:", t == (*(*t).links[0]).back)
	}
}

func makeTree(rule *map[string]string, s string, pars *map[string]*string, def *map[string]*string, deep int) *Tree {
	t := Tree{
		elem:   s,
		pars:   *pars,
		points: map[string]*Tree{},
		resp:   map[*string]*[]*Tree{},
		links:  []*Tree{},
		accept: true,
		done:   false,
	}
	theRule := ""
	exists := true
	theRule, exists = (*rule)[s]
	if !exists || deep > 30 {
		if deep > 30 {
			fmt.Println("THIS! " + s)
		}
		return &t
	}
	encode(theRule, &t, rule, def, deep)
	//if len(t.links) > 0 {fmt.Println("FIRST:", &t == (*t.links[0]).back)}
	return &t
}

func addr(s string) *string {
	return &s
}

func setPointers(t *Tree) {
	//fmt.Println((*t).elem + " -- " + *(*t).pars["type"])
	/*if (*t).elem == "сущ" {
	    *(*t).pars["rod"] = "*"
		*(*t).pars["od"] = "*"
	} else*/if (*t).elem == "дееп" {
		*(*t).pars["time"] = "дееп"
	}
	for index, query := range (*t).pars {
		f := t
		up := split(*query, "^")
		if len(up) > 1 {
			from := split(up[1], "<")
			indexFrom := index
			if len(from) > 1 {
				up[1] = from[0]
				indexFrom = from[1]
			}
			f = (*t).back
			(*t).pars[index] = (*f).pars[indexFrom]
			(*t).points[index] = f
			if (*f).resp[(*t).pars[index]] == nil {
				slice := []*Tree{}
				(*f).resp[(*t).pars[index]] = &slice
			}
			*(*f).resp[(*t).pars[index]] = append(*(*f).resp[(*t).pars[index]], t)
			//fmt.Println((*t).elem  + " <-- " + (*f).elem + " ... " + *(*t).pars[index] + "  (" + index + ")")
			continue
		}
		pointer := split(*query, "*")
		if len(pointer) == 1 {
			f = (*t).back
		}
		to := split(*query, "&")
		if len(to) == 2 {
			if (*f).links == nil {
				fmt.Println("Error: no links from " + (*f).elem)
			}
			from := split(to[1], "<")
			indexFrom := index
			if len(from) > 1 {
				to[1] = from[0]
				indexFrom = from[1]
			}
			for link := range (*f).links {
				_, exists := (*(*f).links[link]).pars["key"]
				if !exists {
					fmt.Println("Error: no KEY field in " + (*(*f).links[link]).elem)
					break
				}
				if *(*f).links[link].pars["key"] == to[1] {
					(*t).pars[index] = (*(*f).links[link]).pars[indexFrom]
					(*t).points[index] = (*f).links[link]
					if (*f).links[link].resp[(*t).pars[index]] == nil {
						slice := []*Tree{}
						(*f).links[link].resp[(*t).pars[index]] = &slice
					}
					*(*f).links[link].resp[(*t).pars[index]] = append(*(*f).links[link].resp[(*t).pars[index]], t)
					//fmt.Println((*t).elem  + " <-- " + (*f).links[link].elem + " !!! " + *(*t).pars[index] + "  (" + index + ")")
					break
				}
			}
		}
	}
	if len((*t).links) > 0 {
		for branch := range (*t).links {
			setPointers((*t).links[branch])
		}
	}
}

func deletePoints(from *Tree, to *Tree) {
	for index, p := range (*from).points {
		if p == to {
			delete((*from).points, index)
			//fmt.Println("DELETE LINK:", (*from).elem, "----->", (*to).elem, "(", index, ")")
		}
	}
}

func checkPoints(rule *map[string]string, next *Tree, set *Set) {
	if !(*next).done {
		for _, p := range (*next).points {
			if p != nil {
				if (*p).back == next {
					fmt.Println((*next).elem, "----->", (*p).elem, "...HMM")
					continue
				}
				checkPoints(rule, p, set)
				deletePoints(next, p)
			}
		}
		makeText(rule, next, set)
		(*next).done = true
	}
}

func offZone(t *Tree, memoryMoveFrom *string, memoryMoveTo *string) {
	/*for try2, memory2 := range (*t).pars {
	    if (*t).resp[memory2] == nil {
	        continue
	    }
	    for _, prev := range *(*t).resp[memory2] {
	        fmt.Println("Prev:", (*prev).elem, "-->", (*t).elem, "#", try2)
	    }
	}*/
	if memoryMoveFrom == nil {
		for _ /*try*/, memory := range (*t).pars {
			if (*t).resp[memory] == nil {
				continue
			}
			for _, prev := range *(*t).resp[memory] {
				for index, memoryMoveFrom := range (*prev).pars {
					if memory == memoryMoveFrom {
						tmp := *(*prev).pars[index]
						(*prev).pars[index] = &tmp
						offZone(prev, memoryMoveFrom, &tmp)
						(*prev).resp[&tmp] = (*prev).resp[memory]
						delete((*prev).resp, memory)
					}
				}
				//fmt.Println("DELETE MEMORY:", prev.elem, "----->", (*t).elem, "(", try, ")")
			}
		}
		return
	}
	if (*t).resp[memoryMoveFrom] == nil {
		return
	}
	//try := ""
	for _, prev := range *(*t).resp[memoryMoveFrom] {
		for index, memory := range (*prev).pars {
			if memory == memoryMoveFrom {
				//try = index
				(*prev).pars[index] = memoryMoveTo
				offZone(prev, memoryMoveFrom, memoryMoveTo)
			}
		}
		(*prev).resp[memoryMoveTo] = (*prev).resp[memoryMoveFrom]
		delete((*prev).resp, memoryMoveFrom)
		//fmt.Println("CHECK MEMORY:", prev.elem, "----->", (*t).elem, "(", try, ")")
	}
}

func addSynonim(oldA string, oldC string, newA string, newC string, query string) string {
	pars := split(query, "|")
	choise := split(pars[random(len(pars))], ".")
	oldNeed := oldC
	newNeed := newC
	if len(choise) > 2 {
		if choise[2] == "^" {
			//oldNeed :=
		}
	}
	switch choise[1] {
	case "take":
		oldNeed = oldA
	case "leave":
		newNeed = newA
	}
	switch choise[0] {
	case "#":
		return newNeed
	case "<":
		return newNeed + " " + oldNeed
	case ">":
		return oldNeed + " " + newNeed
	case "{":
		return newNeed + "-" + oldNeed
	case "}":
		return oldNeed + "-" + newNeed
	}
	return "WRONG_SYNONIM " + newC
}

func searchSynonim(hero *Set, heroL *Set, word *map[string]string, query *map[string]*string) string {
	oldC := formSuch(word, query)
	oldA := (*word)["им"]
	for key, name := range (*hero).Link {
		if oldA == key {
			switch name[1] {
			case "сущ":
				for _, syn := range (*heroL).Such {
					if syn["им"] == name[0] {
						compl := formSuch(&syn, query)
						return addSynonim(oldA, oldC, syn["им"], compl, name[2])
					}
				}
			case "прил":
				for _, syn := range (*heroL).Pril {
					if syn["мим"] == name[0] {
						compl := formPril(&syn, query)
						return addSynonim(oldA, oldC, syn["мим"], compl, name[2])
					}
				}
			}
		}
	}
	return oldC
}

func makeText(rule *map[string]string, t *Tree, set *Set) {
	/*fmt.Println("<BEFORE " + (*t).elem + " >")
	  for v, w := range(*t).pars {fmt.Print(v + "=" + *w + " ")}
	  fmt.Println()*/
	if (*t).back != t {
		(*t).hero = (*(*t).back).hero
		(*t).heroL = (*(*t).back).heroL
	}
	if (*t).back == t {
		(*t).hero = Set{
			Such: []map[string]string{
				createMO("Василий", "м.од.ед.чел.имя"),
				createMO("враг", "м.од.мн.чел"),
				createMN("штаб", "м.не.ед.место"),
			},
			Link: map[string][]string{
				"Василий": []string{"генерал", "сущ", "<.both|#.take|}.both"},
				"штаб":    []string{"квартира", "сущ", "}.take"},
				"враг":    []string{"Центральный", "прил", "<.both.^"},
			},
		}
		(*t).heroL = Set{
			Such: []map[string]string{
				createMO("генерал", "м.од.ед.чел"),
				createJN("квартира", "ж.не.ед.место"),
			},
			Pril: []map[string]string{
				createP("Центральный", ""),
			},
		}
	} else if (*t).elem == "сущ" {
		//fmt.Println(*(*t).pars["rod"], *(*t).pars["od"], *(*t).pars["type"])
		word := findWord(&(*set).Such, []*string{(*t).pars["rod"], (*t).pars["od"], (*t).pars["type"]})
		spec := findWord(&(*t).hero.Such, []*string{(*t).pars["rod"], (*t).pars["od"], (*t).pars["type"]})
		toSyn := false
		if word == nil && spec == nil {
			(*t).accept = false
			return
		} else if word == nil {
			word = spec
			toSyn = true
		} else if word != nil && spec != nil {
			if random(2) == 1 {
				word = spec
				toSyn = true
			}
		}
		if toSyn && random(2) == 1 {
			(*t).elem = searchSynonim(&(*t).hero, &(*t).heroL, word, &(*t).pars)
		} else {
			(*t).elem = formSuch(word, &(*t).pars)
		}
	} else if (*t).elem == "гл" || (*t).elem == "дееп" {
		//fmt.Println(*(*t).pars["type"], *(*t).pars["who"], *(*t).pars["whom"])
		word := findGl(&(*set).Gl, &(*t).pars)
		if word == nil {
			(*t).accept = false
			return
		}
		(*t).elem = formGl(word, &(*t).pars)
	} else if (*t).elem == "прил" {
		//fmt.Println(*(*t).pars["rod"])
		word := findWord(&(*set).Pril, []*string{(*t).pars["type"]})
		if word == nil {
			(*t).accept = false
			return
		}
		(*t).elem = formPril(word, &(*t).pars)
	} else if (*t).elem == "нар" {
		word := findWord(&(*set).Nar, []*string{(*t).pars["type"]})
		if word == nil {
			(*t).accept = false
			return
		}
		(*t).elem = (*word)["об"]
	} else if (*t).elem == "предл" {
		word := findPredl(&(*set).Predl, &(*t).pars)
		if word == nil {
			(*t).accept = false
			return
		}
		(*t).elem = formPredl((*word)[0])
		*(*t).pars["pad"] = (*word)[1]
		*(*t).pars["type"] = (*word)[2]
	}
	/*fmt.Println("<AFTER " + (*t).elem + " >")
		for v, w := range(*t).pars {fmt.Print(v + "=" + *w + " ")}
	    fmt.Println()*/
	offZone(t, nil, nil)
	if len((*t).links) > 0 {
		//fmt.Println(t == (*(*t).links[0]).back)
		for branch := range (*t).links {
			deletePoints((*t).links[branch], t)
			//fmt.Println("!!!")
		}
		for branch := range (*t).links {
			next := (*t).links[branch]
			checkPoints(rule, next, set)
		}
	}
}

func printText(t *Tree) string {
	ret := ""
	if (*t).accept {
		if len((*t).links) > 0 {
			for branch := range (*t).links {
				ret += printText((*t).links[branch])
			}
			return ret
		}
		return " " + (*t).elem
	}
	return " @" + (*t).elem
}

func formPredl(query string) string {
	if query == "whom" || query == "кому" {
		return ""
	}
	return strings.TrimSuffix(query, "~")
}

func formSuch(such *map[string]string, query *map[string]*string) string {
	pad := *(*query)["pad"]
	num := *(*query)["num"]
	//fmt.Println(pad, num)
	*(*query)["type"] = (*such)["par"]
	*(*query)["rod"] = split((*such)["par"], ".")[0]
	*(*query)["od"] = split((*such)["par"], ".")[1]
	if num == "*" {
		num = "ед"
		if random(3) == 0 {
			num = "мн"
		}
		pars := split((*such)["par"], ".")
		for _, p := range pars {
			if p == "ед" {
				num = "ед"
				break
			}
			if p == "мн" {
				num = "мн"
				break
			}
		}
		*(*query)["num"] = num
	}
	if pad == "*" {
		pad = "им"
		*(*query)["pad"] = pad
	}
	if num == "мн" {
		return (*such)["м"+pad]
	}
	return (*such)[pad]
}

func formPril(pril *map[string]string, query *map[string]*string) string {
	pad := *(*query)["pad"]
	num := *(*query)["num"]
	rod := *(*query)["rod"]
	od := *(*query)["od"]
	//fmt.Println(pad, num, rod, od)
	if num == "мн" {
		if pad == "вин" && od == "не" {
			return (*pril)["мнвинне"]
		}
		return (*pril)["мн"+pad]
	}
	if rod == "м" && pad == "вин" && od == "не" {
		return (*pril)["мвинне"]
	}
	return (*pril)[rod+pad]
}

func formGl(gl *map[string]string, query *map[string]*string) string {
	time := *(*query)["time"]
	proc := *(*query)["proc"]
	rod := *(*query)["rod"]
	num := *(*query)["num"]
	features := split((*gl)["par"], "|")
	for _, f := range features {
		one := split(f, "=")
		*(*query)[one[0]] = one[1]
	}
	//fmt.Println(proc, time, rod, num)
	if time == "дееп" {
		return (*gl)[proc+":дееп"]
	}
	if time == "инф" {
		if proc == "проц" {
			return (*gl)["инф"]
		}
		return (*gl)["сов:инф"]
	}
	if time == "наст" {
		if num == "мн" {
			return (*gl)["проц:мн"]
		}
		switch rod {
		case "я":
			return (*gl)["проц:я"]
		case "ты":
			return (*gl)["проц:ты"]
		case "мы":
			return (*gl)["проц:мы"]
		case "вы":
			return (*gl)["проц:вы"]
		case "м", "ж", "ср":
			return (*gl)["проц:ед"]
		case "мн":
			return (*gl)["проц:мн"]
		}
	}
	if time == "прош" {
		if num == "мн" {
			return (*gl)[proc+":"+time+":мн"]
		}
		switch rod {
		case "м":
			return (*gl)[proc+":"+time+":м"]
		case "ж":
			return (*gl)[proc+":"+time+":ж"]
		case "ср":
			return (*gl)[proc+":"+time+":ср"]
		case "мн":
			return (*gl)[proc+":"+time+":мн"]
		}
	}
	if time == "буд" && proc == "сов" {
		if num == "мн" {
			return (*gl)["сов:буд:мн"]
		}
		switch rod {
		case "я":
			return (*gl)["сов:буд:я"]
		case "мы":
			return (*gl)["сов:буд:мы"]
		case "ты":
			return (*gl)["сов:буд:ты"]
		case "вы":
			return (*gl)["сов:буд:вы"]
		case "м", "ж", "ср":
			return (*gl)["сов:буд:ед"]
		case "мн":
			return (*gl)["сов:буд:мн"]
		}
	}
	if time == "буд" && proc == "проц" {
		return (*gl)["сов:инф"]
	}
	fmt.Println("НЕПОДХОДЯЩИЙ_ГЛАГОЛ")
	return "НЕПОДХОДЯЩИЙ_ГЛАГОЛ"
}

/*func changeToMest(mest map[string]map[string]string, pol string) (*map[string]string) {
	switch pol {
	case "м":
		return &mest["он"]
	case "ж":
		return &mest["она"]
	case "ср":
		return &mest["оно"]
	case "мн":
		return &mest["они"]
	}
}*/

func findWord(set *[]map[string]string, query []*string) *map[string]string {
	choise := []int{}
	/*for _, e := range query {
	    fmt.Print(*e, "|")
	}
	fmt.Println("")*/
	for index, i := range *set {
		accept := true
		pars := split(i["par"], ".")
		for _, fromQ := range query {
			qParts := split(*fromQ, ".")
			//fmt.Println("Query/Par", *fromQ, i["par"])
			if *fromQ == "*" {
				continue
			}
			exists := false
			for _, q := range qParts {
				for _, p := range pars {
					if q == p || p == "*" {
						exists = true
						break
					}
				}
			}
			if !exists {
				accept = false
				break
			}
		}
		if accept {
			choise = append(choise, index)
		}
	}
	if len(choise) == 0 {
		fmt.Println("Error: cant find word")
		return nil
	}
	ret := randomArr(choise)
	(*set)[ret]["used"] = "used"
	return &(*set)[ret]
}

func findGl(set *[]map[string]string, query *map[string]*string) *map[string]string {
	choise := []int{}
	for index, i := range *set {
		accept := true
		blocks := split(i["par"], "|")
		for _, b := range blocks {
			line := split(b, "=")
			//fmt.Println(b)
			fromQ := split(*(*query)[line[0]], ".")
			if fromQ[0] == "*" {
				continue
			}
			//fmt.Println(line[0], line[1], *(*query)[line[0]])
			if len(line) < 2 || line[1] == "" {
				accept = false
				break
			}
			pars := split(line[1], ".")
			exists := false
			for _, q := range fromQ {
				if q == "*" {
					exists = true
					break
				}
				for _, p := range pars {
					if q == p || p == "*" {
						exists = true
						break
					}
				}
			}
			if !exists {
				accept = false
				break
			}
		}
		if accept {
			choise = append(choise, index)
		}
	}
	if len(choise) == 0 {
		fmt.Println("Error: cant find gl")
		return nil
	}
	ret := randomArr(choise)
	(*set)[ret]["used"] = "used"
	return &(*set)[ret]
}

func findPredl(set *[][]string, query *map[string]*string) *[]string {
	choise := []int{}
	for index, i := range *set {
		if *(*query)[i[0]] == "" {
			continue
		}
		choise = append(choise, index)
	}
	if len(choise) == 0 {
		fmt.Println("Error: cant find predl")
		return nil
	}
	ret := randomArr(choise)
	sl := []string{(*set)[ret][0], (*set)[ret][1], *(*query)[(*set)[ret][0]]}
	return &sl
}

func print(t *Tree, deep int) {
	if len((*t).links) == 0 {
		fmt.Print("DEEP=" + strconv.Itoa(deep) + ": [" + (*t).elem + "]")
	} else {
		for i := range (*t).links {
			print((*t).links[i], deep+1)
		}
		fmt.Println("DEEP=" + strconv.Itoa(deep) + ": [" + (*t).elem + "]")
	}
}

func up() string {
	return all("^")
}

func all(a string) string {
	p := a
	if a != "^" {
		p = "&" + a
	}
	return "time=" + p + ".proc=" + p + ".rod=" + p + ".pad=" + p + ".who=" + p + ".whom=" + p + ".num=" + p + ".type=" + p + ".od=" + p +
		".кому=" + p + ".на=" + p + ".в=" + p + ".с=" + p + ".под=" + p + ".над=" + p + ".за=" + p + ".о=" + p + ".у=" + p +
		".к=" + p + ".на~=" + p + ".в~=" + p + ".с~=" + p + ".под~=" + p + ".над~=" + p + ".за~=" + p
}

func tail(tails *map[string]string, from string, points string) string {
	ret := (*tails)[from]
	arr := split(strings.Replace(points, " ", "", -1), ",")
	for i := range arr {
		beg := "&"
		if arr[i] == "^" {
			beg = ""
		}
		ret = strings.Replace(ret, "&"+strconv.Itoa(i), beg+arr[i], -1)
	}
	//fmt.Println(ret)
	return from + ":" + ret
}

func orSame(s string, put string) string {
	ret := strings.Replace(s, "|", ":"+put+"|", -1)
	ret += ":" + put
	return ret
}

func orR(amount []int, s ...string) string {
	ret := []string{}
	l := len(s)
	for index, kol := range amount {
		if index >= l {
			break
		}
		tmp := []string{}
		for i := 0; i < kol; i++ {
			tmp = append(tmp, s[index])
		}
		ret = append(ret, strings.Join(tmp, "|"))
	}
	return "(" + strings.Join(ret, "|") + ")"
}

func or(s ...string) string {
	return "(" + strings.Join(s, "|") + ")"
}

func plus(s ...string) string {
	return "(" + strings.Join(s, "+") + ")"
}

func predl(par ...string) string {
	p := map[string]string{
		"кому": "",
		"на":   "",
		"на~":  "",
		"в":    "",
		"в~":   "",
		"за":   "",
		"над":  "",
		"под":  "",
		"за~":  "",
		"над~": "",
		"под~": "",
		"у":    "",
		"о":    "",
		"к":    "",
		"с":    "",
		"с~":   "",
	}
	for _, i := range par {
		switch i {
		case "сост":
			p["на"] += ".пред.дер"
			p["в"] += ".вмест.место.мат"
			p["за"] = p["на"]
			p["над"] = p["на"]
			p["под"] = p["на"]
			p["у"] += p["на"] + ".чел"
			p["с"] += ".од.пред"
		case "мысль":
			p["о"] = "*"
		case "у":
			p["у"] += ".од"
		case "кому":
			p["кому"] += ".од"
		case "движ":
			p["на~"] += ".пред.напр"
			p["в~"] += ".вмест.место"
			p["за~"] = p["на~"]
			p["над~"] = p["на~"]
			p["под~"] = p["на~"]
			p["с~"] = p["на~"]
		}
	}
	ret := ""
	for key, elem := range p {
		ret += "|" + key + "=" + elem
	}
	//fmt.Println(ret)
	return strings.Replace(ret, "=.", "=", -1)
}

func main() {
	rule := map[string]string{}
	tails := map[string]string{}

	tails["СКАЗУЕМОЕ"] = "rod=&0.num=&0.who=&0<type"
	tails["ДОПОЛНЕНИЕ"] = "pad=&0.type=&0"

	rule["ТЕКСТ"] = "АБЗАЦ+(ТЕКСТ|)"
	rule["АБЗАЦ"] = "абзац+(П|П|П|АБЗАЦ+АБЗАЦ|ДИАЛОГ)+\n"
	rule["П"] = "ПРЕДЛОЖЕНИЕ+ПРЕДЛОЖЕНИЕ+ПРЕДЛОЖЕНИЕ|П+П"
	rule["ПРЕДЛОЖЕНИЕ"] = plus(orR([]int{4, 1, 1, 1, 1, 1, 1}, "П1|П2|П3|П4|П5|П6|П7"), ".")
	rule["П1"] = "ЧАСТЬ"
	rule["П2"] = "ЧАСТЬ+,+(хотя|если|как только|не смотря на то, что|когда|как будто)+ЧАСТЬ"
	rule["П3"] = "ЧАСТЬ+как+ДОПОЛНЕНИЕ:pad=им"
	rule["П4"] = "ЧАСТЬ+,+как будто+СОЧЕТАНИЕ"
	rule["П5"] = "ЧАСТЬ+,+чтобы+СКАЗУЕМОЕ:time=инф"
	rule["П6"] = "до чего же+СОЧЕТАНИЕ"
	rule["П7"] = "это+каково+, что+(ЧАСТЬ|СОЧЕТАНИЕ)"
	//rule["П10"] = "ЧАСТЬ:rod=*&p.key=m.num=*&p.type=*&p+,+как будто+СКАЗУЕМОЕ:rod=&m.num=&m.who=&m<type"

	rule["ЧАСТЬ"] = plus("сущ:key=p", or("", "", "ДЕЕПОБОР:who=&p<type"), tail(&tails, "СКАЗУЕМОЕ", "p"), or("", "", plus("и", tail(&tails, "СКАЗУЕМОЕ", "p"))))
	rule["СКАЗУЕМОЕ"] = orSame("Д1|Д1|Д2", up())
	rule["Д1"] = plus("(||нар)", "гл:key=s."+up(), orR([]int{5, 4, 1}, "", "ДОП:"+all("s"), "о том, чтобы+СКАЗУЕМОЕ:time=инф"))
	rule["Д2"] = plus("(||нар)", "гл:key=s."+up()+".type=относ", "гл:key=ha.time=инф.proc=проц.who=^",
		orR([]int{5, 4, 1}, "", "ДОП:"+all("ha"), "о том, чтобы+СКАЗУЕМОЕ:time=инф"))
	rule["ДОПОЛНЕНИЕ"] = "сущ:key=podl." + up() + "+(|||,+прил:type=который.rod=&podl.num=&podl+СКАЗУЕМОЕ:rod=&podl.num=&podl.who=&podl<type+,)"
	rule["СОЧЕТАНИЕ"] = "сущ:key=podl.pad=им+прил:rod=&podl.pad=&podl.num=&podl.type=&podl"
	rule["ДЕЕПОБОР"] = ",+(||нар)+дееп:key=de.who=^+(|" + "ДОП:" + all("de") + ")+,"
	//rule["ПРИЧОБОР"] = ",+(||нар)+прич:key=podl+" + tail(&tails, "КОГО/ЧТО", "podl") + "+,"
	rule["ДОП"] = plus("предл:key=z."+up(), tail(&tails, "ДОПОЛНЕНИЕ", "z"))

	rule["каково"] = "здорово|плохо|хорошо|лучше некуда"

	rule["ДИАЛОГ"] = "ВОПРОСОТВЕТ"
	rule["ВОПРОСОТВЕТ"] = "КАК?:rod=&podl2.num=&podl2+-+гл:rod=&podl.num=&podl.type=речь+сущ:key=podl.od=од.type=чел+.\n" +
		"+ТАК!:rod=&podl2.num=&podl2+-+сущ:key=podl2.od=од.type=чел+гл:rod=&podl2.num=&podl2.type=речь+."
	rule["КАК?"] = "как+ты+гл:key=skaz.rod=^.num=^+" + "ДОП:" + all("skaz") + "+?"
	rule["ТАК!"] = "тире+(я+гл:key=skaz.rod=^.num=^+" + "ДОП:" + all("skaz") + "|нар|нар+и+нар|ДОПОЛНЕНИЕ:pad=тв)"
	rule["тире"] = "\t-"
	rule["абзац"] = "\t"

	rule["W"] = "сущ:key=d.type=слово+Q:who=&d<type"
	rule["Q"] = "гл:rod=м.who=^+гл:rod=м.who=^"

	set := Set{
		Mest: map[string]map[string]string{
			"я": map[string]string{
				"им":  "я",
				"род": "меня",
				"дат": "мне",
				"вин": "меня",
				"тв":  "мной",
				"о":   "мне",
			},
			"ты": map[string]string{
				"им":  "ты",
				"род": "тебя",
				"дат": "тебе",
				"вин": "тебя",
				"тв":  "тобой",
				"о":   "тебе",
			},
			"мы": map[string]string{
				"им":  "мы",
				"род": "нас",
				"дат": "нам",
				"вин": "нас",
				"тв":  "нами",
				"о":   "нас",
			},
			"вы": map[string]string{
				"им":  "вы",
				"род": "вас",
				"дат": "вам",
				"вин": "вас",
				"тв":  "вами",
				"о":   "вас",
			},
			"он": map[string]string{
				"им":  "он",
				"род": "его",
				"дат": "ему",
				"вин": "его",
				"тв":  "им",
				"о":   "нем",
			},
			"она": map[string]string{
				"им":  "она",
				"род": "ее",
				"дат": "ей",
				"вин": "ее",
				"тв":  "ею",
				"о":   "ней",
			},
			"оно": map[string]string{
				"им":  "оно",
				"род": "его",
				"дат": "ему",
				"вин": "его",
				"тв":  "им",
				"о":   "нем",
			},
			"они": map[string]string{
				"им":  "они",
				"род": "их",
				"дат": "им",
				"вин": "их",
				"тв":  "ими",
				"о":   "них",
			},
		},
		Such: []map[string]string{
			createMO("Аркадий", "м.од.имя.чел"),
			createMO("медведь", "м.од.имя.чел"),
			createMO("бомж", "м.од.чел"),
			createMO("монстр", "м.од.чел"),
			createMO("царь", "м.од.чел"),
			createJO("Катерина", "ж.од.ед.имя.чел"),
			createJO("Наташа", "ж.од.ед.имя.чел"),
			createJO("тетя", "ж.од.чел"),
			createJO("старуха", "ж.од.чел"),
			createJO("корова", "ж.од.жив"),
			createJO("рыба", "ж.од.жив.еда"),
			createMN("стул", "м.не.пред"),
			createMN("навоз", "м.не.ед.пред.мат.вмест"),
			createMN("тополь", "м.не.раст.дер"),
			createJN("тарелка", "ж.не.пред"),
			createJN("нога", "ж.не.пред.чт"),
			createJN("педаль", "ж.не.пред"),
			createS("чудовище", "ср.од.чел"),
			createMN("золото", "ср.не.ед.мат"),
			createS("время", "ср.не"),
			createJN("деревня", "ж.не.место"),
			createJN("сосредоточенность", "ж.не.ед.кач"),
			createJN("задача", "ж.не.проб"),
			createMN("вопрос", "м.не.проб"),
			createMN("час", "м.не.время"),
			createMN("год", "м.не.время"),
			createJN("вермишель", "ж.не.ед.еда"),
			createMO("Михаил", "м.од.ед.чел"),
			createS("слово", "ср.не.слово"),
			createJN("работа", "ж.не.дейст"),
			createMN("север", "м.не.ед.напр"),
		},
		Gl: []map[string]string{
			createG("сидеть", "посидеть", "type=сост|who=од|whom=время"+predl("сост")),
			createG("спать", "уснуть", "type=сост|who=од|whom=время"+predl("сост")),
			createG("бежать", "убежать", "type=движ|who=од.время|whom=время"+predl("движ")),
			createG("ронять", "уронить", "type=дейст|who=од|whom=пред"+predl("куда")),
			createG("есть", "съесть", "type=дейст|who=од|whom=еда"+predl("сост")),
			createG("делать", "сделать", "type=дейст|who=од|whom=дейст"+predl("сост")),
			createG("взрывать", "взорвать", "type=дейст|who=од|whom=пред.место"+predl("сост")),
			createG("бить", "убить", "type=дейст|who=од|whom=од"+predl("сост")),
			createG("мешать", "помешать", "type=дейст|who=од.пред.раст|whom="+predl("сост", "кому")),
			createG("догонять", "догнать", "type=дейст|who=од|whom=од"+predl("сост")),
			createG("думать", "подумать", "type=относ|who=чел|whom="+predl("мысль")),
			createG("говорить", "сказать", "type=речь|who=чел|whom=слово"+predl("мысль", "кому")),
			createG("видеть", "увидеть", "type=дейст|who=од|whom=*"+predl("сост")),
			createG("уметь", "суметь", "type=относ|who=од|whom="+predl()),
			createG("хотеть", "захотеть", "type=относ|who=од|whom=пред.жив"+predl()),
			createG("красть", "украсть", "type=действ|who=од|whom=пред.жив.мат"+predl("у")),
			createG("спрашивать", "спросить", "type=речь|who=чел|whom=чел"+predl("мысль", "у")),
			createG("дарить", "подарить", "type=дейст|who=чел|whom=пред.жив"+predl("кому")),
			createG("забывать", "забыть", "type=дейст|who=чел|whom=од.место"+predl()),
			createG("начинать", "начать", "type=относ|who=*|whom=дейст"+predl()),
			createG("пугать", "испугать", "type=дейст|who=*|whom=од"+predl()),
		},
		Predl: [][]string{
			[]string{"на", "о"},
			[]string{"в", "о"},
			[]string{"за", "тв"},
			[]string{"над", "тв"},
			[]string{"под", "тв"},
			[]string{"с", "тв"},
			[]string{"на~", "вин"},
			[]string{"в~", "вин"},
			[]string{"за~", "вин"},
			[]string{"над~", "вин"},
			[]string{"под~", "вин"},
			[]string{"с~", "род"},
			[]string{"о", "о"},
			[]string{"у", "род"},
			[]string{"к", "дат"},
			[]string{"кому", "дат"},
			[]string{"whom", "вин"},
		},
		Pril: []map[string]string{
			createP("который", "который"),
			createP("хороший", "од.не.чел.жив.предм.место.слово.дейст.мат"),
			createP("большой", "од.чел.жив.предм.место"),
			createP("тупой", "од.чел.жив.предм.место"),
			createP("радостный", "од.чел.жив"),
			createP("солнечный", "время.место"),
			createP("ужасный", "од.чел.жив.предм.место.слово.мат"),
			createP("светлый", "время.пред"),
			createP("вкусный", "еда.жив"),
			createP("старый", "жив.чел.место.пред"),
		},
		Nar: []map[string]string{
			createN("быстро", "движ.мысль.действ"),
			createN("медленно", "движ.мысль.действ"),
			createN("с любовью", "дейст"),
			createN("с трудом", "*"),
			createN("легко", "*"),
			createN("весело", "*"),
		},
		Srav: [][]string{
			[]string{"слишком", "+"},
			[]string{"очень", "++"},
			[]string{"немного", ""},
			[]string{"совсем", "+++"},
		},
	}

	pars := map[string]*string{
		"key":  addr("*"),
		"elem": addr("*"),
		"time": addr("прош"),
		"proc": addr("сов"),
		"rod":  addr("*"),
		"pad":  addr("им"),
		"num":  addr("*"),
		"od":   addr("*"),
		"type": addr("*"),
		"who":  addr("*"),
		"whom": addr("*"),
		"кому": addr("*"),
		"на":   addr("*"),
		"на~":  addr("*"),
		"в":    addr("*"),
		"в~":   addr("*"),
		"за":   addr("*"),
		"над":  addr("*"),
		"под":  addr("*"),
		"за~":  addr("*"),
		"над~": addr("*"),
		"под~": addr("*"),
		"у":    addr("*"),
		"о":    addr("*"),
		"к":    addr("*"),
		"с":    addr("*"),
		"с~":   addr("*"),
	}

	rule["START"] = "ТЕКСТ"
	tree := makeTree(&rule, "START", &pars, &pars, 0)
	tree.back = tree
	setPointers(tree)
	makeText(&rule, tree, &set)
	printCompact(check(printText(tree)), 20)
}
