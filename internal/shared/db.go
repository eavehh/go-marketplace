package shared

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func Wait_for_db(db *sql.DB, max_attemps int, interval time.Duration) error {
	for i := 1; i < max_attemps; i++ {
		if err := db.Ping(); err != nil {
			log.Printf("Wait_for_db: try %d/%d: %v", i, max_attemps, err)
			time.Sleep(interval)
			continue
		}
		log.Printf("Wait_for_db: DB available; try %d/%d", i, max_attemps)
		return nil
	}
	return fmt.Errorf("Wait_for_db: DB unavailable after %d times try", max_attemps)

}
