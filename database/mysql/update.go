package mysql

import "log"

func UpdateUser(value [3]string, email, password string) {
	db := Connect()
	q := "UPDATE Signup SET fullnames=?, accountnames=?, is_active=? WHERE emails=? AND passwords=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	_, err = update.Exec(value[0], value[1], value[2], email, password)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateProfile(value [4]string, email string) {
	db := Connect()
	q := "UPDATE Profiles SET titles=?, phones=?, locations=?, working_statuses=? WHERE emails=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	_, err = update.Exec(value[0], value[1], value[2], value[3], email)
	if err != nil {
		log.Fatal(err)
	}
}
