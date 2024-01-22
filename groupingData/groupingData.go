package groupingData

import "fmt"

func GroupingData() {
	//arraies()
	//slice()
	//arraiesTwo()
	//arraiesThree()
	//maps()
	//mapsTwo()
	structs()
}

func arraies() {
	//literal array
	arr := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "BittersweetChocolate (GF)", "Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)", "Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy ChocolatePeanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy SunbutterCinnamon (GF, V)",
		"Orange Cream (GF) ", "Peanut Butter Cookie Dough", "RaspberrySorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
		"Strawberry LemonadeSorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "WolverineTracks (GF)"}

	fmt.Printf("array length: %v %T", len(arr), arr)
}

func slice() {
	s := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "BittersweetChocolate (GF)", "Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)", "Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)", "Non-Dairy ChocolatePeanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy SunbutterCinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "RaspberrySorbet (GF, V)", "Salty Caramel (GF)",
		"Slate Slate Different", "Strawberry LemonadeSorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "WolverineTracks (GF)"}

	for i, v := range s {
		fmt.Printf("index %v to the value %v\n", i, v)
	}

	s = append(s, "added value")
}

func arraiesTwo() {
	a := [5]int{5, 10, 15, 20, 25}
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range a {
		fmt.Printf("the value %v, in the index %v, with type %T\n", i, v, v)
	}
	for i, v := range b {
		fmt.Printf("the value %v, in the index %v, with type %T in slice\n", i, v, v)
	}
	b1 := b[0:5]
	fmt.Println("the b1 is: ", b1)

	b2 := b[5:]
	fmt.Println("the b2 is: ", b2)

	b3 := b[2:7]
	fmt.Println("the b3 is: ", b3)

	b4 := b[1:6]
	fmt.Println("the b4 is: ", b4)
}

func arraiesThree() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println("the x slice is: ", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("the x slice is: ", x)

	fmt.Println("----------")

	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := x[0:3]
	fmt.Println("the x1 slice is: ", x1)
	x2 := x[6:]
	fmt.Println("the x2 slice is: ", x2)
	z := []int{}
	z = append(z, x1...)
	z = append(z, x2...)
	fmt.Println("the z slice is: ", z)

	fmt.Println("----------")

	a := make([]string, 0)
	a = append(a, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", "Delaware",
		" Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", "Kentucky", " Louisiana", " Maine",
		" Maryland", " Massachusetts", " Michigan", " Minnesota", "Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey",
		" New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon",
		" Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", "Utah", " Vermont",
		" Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")

	for i, v := range a {
		fmt.Printf("this is the s slice, with cap %v, len %v, and the value in the index %v, is %v\n", cap(a), len(a), i, v)
	}

	fmt.Println("----------")

	b := make([][]string, 0)

	b = append(b, []string{"James", "Bond", "Shaken, not stirred"})
	b = append(b, []string{"Miss", "Moneypenny", "I'm 008."})

	for _, v := range b {
		fmt.Printf("this is the value %v\n", v)
	}

}

func maps() {
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	for k, v := range m {
		fmt.Printf("this is the key %v\n", k)
		for i, vlr := range v {
			fmt.Println("\t", i, vlr)
		}
	}

	fmt.Println("-----------------")

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Printf("this is the key %v\n", k)
		for i, vlr := range v {
			fmt.Println("\t", i, vlr)
		}
	}

	fmt.Println("-----------------")

	delete(m, `fleming_ian`)

	for k, v := range m {
		fmt.Printf("this is the key %v\n", k)
		for i, vlr := range v {
			fmt.Println("\t", i, vlr)
		}
	}
	fmt.Println("-----------------")
}

func mapsTwo() {
	words := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some",
		"advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever",
		"you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all",
		"the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve",
		"had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually",
		"communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant",
		"a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to",
		"reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious",
		"natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few",
		"veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach",
		"itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so",
		"it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a",
		"politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown",
		"men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have",
		"feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by",
		"some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on",
		"the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least",
		"the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and",
		"marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of",
		"infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i",
		"forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly",
		"repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out",
		"unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i",
		"come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be",
		"founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain",
		"point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from",
		"the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in",
		"uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no",
		"more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.",
		"only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was",
		"exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which",
		"i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of",
		"successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,",
		"some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were",
		"related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes",
		"ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with",
		"that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the",
		"creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic",
		"readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and",
		"which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out",
		"all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust",
		"floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my",
		"interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my",
		"family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western",
		"city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,",
		"and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of",
		"buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s", "brother,", "who", "came", "here", "in", "fifty-one,", "sent", "a", "substitute", "to", "the", "civil",
		"war,", "and", "started", "the", "wholesale", "hardware", "business", "that", "my", "father",
		"carries", "on", "today.", "i", "never", "saw", "this", "great-uncle,", "but", "i’m", "supposed", "to",
		"look", "like", "him—with", "special", "reference", "to", "the", "rather", "hard-boiled", "painting",
		"that", "hangs", "in", "father’s", "office.", "i", "graduated", "from", "new", "haven", "in", "1915,",
		"just", "a", "quarter", "of", "a", "century", "after", "my", "father,", "and", "a", "little", "later", "i",
		"participated", "in", "that", "delayed", "teutonic", "migration", "known", "as", "the", "great",
		"war.", "i", "enjoyed", "the", "counter-raid", "so", "thoroughly", "that", "i", "came", "back",
		"restless.", "instead", "of", "being", "the", "warm", "centre", "of", "the", "world,", "the", "middle",
		"west", "now", "seemed", "like", "the", "ragged", "edge", "of", "the", "universe—so", "i",
		"decided", "to", "go", "east", "and", "learn", "the", "bond", "business.", "everybody", "i", "knew",
		"was", "in", "the", "bond", "business,", "so", "i", "supposed", "it", "could", "support", "one",
		"more", "single", "man.", "all", "my", "aunts", "and", "uncles", "talked", "it", "over", "as", "if",
		"they", "were", "choosing", "a", "prep", "school", "for", "me,", "and", "finally", "said,",
		"why—ye-es,", "with", "very", "grave,", "hesitant", "faces.", "father", "agreed", "to", "finance",
		"me", "for", "a", "year,", "and", "after", "various", "delays", "i", "came", "east,", "permanently,",
		"i", "thought,", "in", "the", "spring", "of", "twenty-two.", "the", "practical", "thing", "was", "to",
		"find", "rooms", "in", "the", "city,", "but", "it", "was", "a", "warm", "season,", "and", "i", "had",
		"just", "left", "a", "country", "of", "wide", "lawns", "and", "friendly", "trees,", "so", "when", "a",
		"young", "man", "at", "the", "office", "suggested", "that", "we", "take", "a", "house", "together",
		"in", "a", "commuting", "town,", "it", "sounded", "like", "a", "great", "idea.", "he", "found", "the",
		"house,", "a", "weather-beaten", "cardboard", "bungalow", "at", "eighty", "a", "month,", "but",
		"at", "the", "last", "minute", "the", "firm", "ordered", "him", "to", "washington,", "and", "i", "went",
		"out", "to", "the", "country", "alone.", "i", "had", "a", "dog—at", "least", "i", "had", "him", "for", "a",
		"few", "days", "until", "he", "ran", "away—and", "an", "old", "dodge", "and", "a", "finnish",
		"woman,", "who", "made", "my", "bed", "and", "cooked", "breakfast", "and", "muttered",
		"finnish", "wisdom", "to", "herself", "over", "the", "electric", "stove.", "it", "was", "lonely", "for",
		"a", "day", "or", "so", "until", "one", "morning", "some", "man,", "more", "recently", "arrived",
		"than", "i,", "stopped", "me", "on", "the", "road.", "how", "do", "you", "get", "to", "west", "egg",
		"village?", "he", "asked", "helplessly.", "i", "told", "him.", "and", "as", "i", "walked", "on", "i",
		"was", "lonely", "no", "longer.", "i", "was", "a", "guide,", "a", "pathfinder,", "an", "original",
		"settler.", "he", "had", "casually", "conferred", "on", "me", "the", "freedom", "of", "the",
		"neighbourhood.", "and", "so", "with", "the", "sunshine", "and", "the", "great", "bursts", "of",
		"leaves", "growing", "on", "the", "trees,", "just", "as", "things", "grow", "in", "fast", "movies,", "i",
		"had", "that", "familiar", "conviction", "that", "life", "was", "beginning", "over", "again", "with",
		"the", "summer.", "there", "was", "so", "much", "to", "read,", "for", "one", "thing,", "and", "so",
		"much", "fine", "health", "to", "be", "pulled", "down", "out", "of", "the", "young", "breath-giving",
		"air.", "i", "bought", "a", "dozen", "volumes", "on", "banking", "and", "credit", "and",
		"investment", "securities,", "and", "they", "stood", "on", "my", "shelf", "in", "red"}

	m := map[string]int{}
	for _, v := range words {

		if _, ok := m[v]; !ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		fmt.Printf("this is the m map: word %v, count: %v\n", k, v)
	}
}

type person struct {
	name     string
	surname  string
	flavours []string
}
type engine struct {
	electric bool
}
type vehicule struct {
	engine
	model string
	color string
}

func structs() {
	p1 := person{
		name:     "Natalia",
		surname:  "Peláez",
		flavours: []string{"Chocolate", "orellana"},
	}
	p2 := person{
		name:     "Jordan",
		surname:  "Sáenz",
		flavours: []string{"Vainilla", "Almendra"},
	}
	fmt.Println("This is the p1 person: ", p1)
	fmt.Println("This is the p2 person: ", p2)

	fmt.Println("----------------")

	m := map[string]person{
		p1.surname: p1,
		p2.surname: p2,
	}

	fmt.Println("This is the m map: ", m)
	fmt.Println("----------------")

	v1 := vehicule{
		engine: engine{
			electric: false,
		},
		model: "1995",
		color: "blanco",
	}
	v2 := vehicule{
		engine: engine{
			electric: true,
		},
		model: "1985",
		color: "verde",
	}
	fmt.Println("This is the v1 vehicule: ", v1)
	fmt.Println("This is the v2 vehicule: ", v2)
	fmt.Println("----------------")

	v3 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Camilo",
		friends: map[string]int{
			"carol": 2,
		},
		favDrinks: []string{"Soda", "aguadepanela"},
	}

	fmt.Println("This is the v3 vehicule: ", v3)
}
