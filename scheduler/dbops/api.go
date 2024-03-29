package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES (?)")
	if err != nil {
		log.Printf("AddVideoDeletionRecord: %v", err)
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeleteionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
