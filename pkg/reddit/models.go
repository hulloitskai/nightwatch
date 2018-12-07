package reddit

// Activity describes some sort of activity that took place between two users
// on Reddit.
type Activity struct {
	Title, Desc string
	Link        string
}

// A Relation is a relationship between two users on Reddit.
type Relation struct {
	ActCount  int
	LatestAct *Activity
}

// A User describes a Reddit user.
type User struct {
	ID       string
	Username string
	Link     string // a link to the User's profile
}
