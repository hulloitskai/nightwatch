package reddit

import (
	"sync"
)

// A Graph is a concurrent map of relationships between Reddit users.
type Graph struct {
	rels  map[string]map[string]*Relation // relationships between users
	users map[string]*User

	mux *sync.RWMutex // mutual exclusion sync lock
}

func NewGraph() *Graph {
	return &Graph{
		rels:  make(map[string]map[string]*Relation),
		users: make(map[string]*User),
		mux:   new(sync.RWMutex),
	}
}

func (g *Graph) RemoveUser(uid string) {
	g.mux.Lock()
	delete(g.users, uid)
	delete(g.rels, uid)
	g.mux.Unlock()
}

func (g *Graph) AddActivity(from, to *User, a *Activity) {
	g.mux.Lock()

	// Initialize maps if users are new.
	if _, ok := g.users[to.ID]; !ok {
		g.users[to.ID] = to
	}
	if _, ok := g.users[from.ID]; !ok {
		g.users[from.ID] = from
		g.rels[from.ID] = make(map[string]*Relation)
	}

	// Update relationship between users.
	if rel, ok := g.rels[from.ID][to.ID]; ok {
		rel.ActCount++
		rel.LatestAct = a
	} else {
		g.rels[from.ID][to.ID] = &Relation{
			ActCount:  1,
			LatestAct: a,
		}
	}

	g.mux.Unlock()
}
