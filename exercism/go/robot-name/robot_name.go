// Package robotname provides functions to name robots with unique and random names.
package robotname

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

// Robot struct keeps the robot's name
type Robot struct {
	name string
}

// Name will return the robots name, or generate a new one if no name.
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = randName()
	}
	return r.name
}

// Reset sets the robot's name to "" and causes us to generate a new name.
func (r *Robot) Reset() { r.name = "" }

// keep track of names so we don't use the same name twice
var usedNames = map[string]bool{}

// max number of unique names
const maxNames = 26 * 26 * 10 * 10 * 10

// generate a new random name, and try again if we already have it.
func randName() string {
	if len(usedNames) == maxNames {
		log.Fatalf("we have used all the combinations of robot names")
	}
	name := genNewName()
	for usedNames[name] {
		name = genNewName()
	}
	usedNames[name] = true
	return name
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// have to guarantee that the name has two capital letters and three digits
// I am assuming the three digits can include 0's anywhere.
func genNewName() string {
	return string(r.Intn(26)+65) + string(r.Intn(26)+65) +
		strconv.Itoa(r.Intn(10)) + strconv.Itoa(r.Intn(10)) + strconv.Itoa(r.Intn(10))
}
