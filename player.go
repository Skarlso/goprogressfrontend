package main

// Item a representation of an Item and it's properties.
type Item struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Dmg    int    `json:"dmg"`
	Weight int    `json:"weight"`
	Armor  int    `json:"armor"`
	Value  int    `json:"value"`
	Chance int    `json:"chance"`
}

// Inventory holds an endless number of Items
type Inventory struct {
	Items    []Item `json:"items"`
	Capacity int    `json:"capacity"`
}

// Stats contains a groupped information about stats of a character
type Stats struct {
	Strenght     int `json:"strength"`
	Agility      int `json:"agility"`
	Intelligence int `json:"intelligence"`
	Perception   int `json:"perception"`
	Luck         int `json:"luck"`
	Constitution int `json:"consititution"`
}

// Body Represents a body of a Player which defines what he wears,
// Player will always automatically wear the best gear.
type Body struct {
	LRing   Item `json:"lring"`
	RRing   Item `json:"rring"`
	Armor   Item `json:"armor"`
	Head    Item `json:"head"`
	Weapond Item `json:"weapond"`
	Shield  Item `json:"shield"`
}

// Cast the cast of a player, like mage, rouge, warrior...
type Cast struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Race the race of the player, like elf, gnome, human, dwarf...
type Race struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Character is a player character.
type Character struct {
	ID          string    `json:"id"`
	Inventory   Inventory `json:"inventory"`
	Body        Body      `json:"body"`
	Name        string    `json:"name"`
	Stats       Stats     `json:"stats"`
	Hp          int       `json:"hp"`
	MaxHp       int       `json:"maxhp"`
	CurrentXp   int       `json:"currentxp"`
	NextLevelXp int       `json:"nextlevelxp"`
	Gold        int       `json:"gold"`
	Level       int       `json:"level"`
	Race        int       `json:"race"`
	Cast        int       `json:"cast"`
}
