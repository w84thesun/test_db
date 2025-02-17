package pagila

import (
	"time"
)

//go:generate reform

// Category represents a row in category table.
//reform:category
type Category struct {
	CategoryID int32     `reform:"category_id,pk"`
	Name       string    `reform:"name"`
	LastUpdate time.Time `reform:"last_update"`
}
