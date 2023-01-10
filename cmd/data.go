package main

// <- Dies sind Kommentar Angaben
// Dies ist ein Sample
// Die einzelnen Structs sind selber hier anzulegen. Der Merge der Daten sollte ohne Probleme vollzogen werden können
// Grundsätzlich sollte der Code auch dokumentiert werden. Wie auch hier gilt, je mehr Inhalt
// in weniger Zeilen ist hilfreicher, als ein ganzer Dokublock.
// Hinweis: Die Einrückungen dienen lediglich der Übersicht!
type Customer struct {
	Name    string `json:"name"`
	Address string `json:"adress"`
	Tel     string `json:"tel"`
	Email   string `json:"mail"`
}

// Ab Hier können einzelne Structs hinzugefpügt werden
<<<<<<< HEAD
type Car struct {
	Type  string `json:"name"`
	Price string `json:"adress"`
}
=======
type City struct {
	Name       string `json:"name"`
	Einwohner  int    `json:"einwohner"`
	Bundesland string `json:"bundesland"`
}

type Project struct {
	Name     string   `json:"name"`
	Customer Customer `json:"customer"`
	City     City     `json:"city"`
	Budget   string   `json:"budget"`
}

type InfologistixMembers struct {
	Name    string  `json:"name"`
	Email   string  `json:"mail"`
	City    City    `json:"city"`
	Project Project `json:"project"`
}

//Neue Zeile
>>>>>>> 9f2b1ed69feb278228ea0d88cab4c5e20b5355b6
