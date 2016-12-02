package DAL

import (
	"time"
)

type Driver struct {
	Id int
	Uuid string
	Name string
	CreatedAt time.Time
}

/*type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}*/
/*type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}*/

/*// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = DB.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = DB.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = DB.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = DB.Exec(statement)
	return
}*/

// Create a new user, save user info into the database
func (user *Driver) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "insert into users (uuid, name, created_at) values ($1, $2, $3) returning id, uuid, created_at"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(createUUID(), user.Name, time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// Delete user from database
func (user *Driver) Delete() (err error) {
	statement := "delete from users where id = $1"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// Update user information in the database
func (user *Driver) Update() (err error) {
	statement := "update users set name = $2 where id = $1"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name)
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = DB.Exec(statement)
	return
}

// Get all users in the database and returns it
func Users() (users []Driver, err error) {
	rows, err := DB.Query("SELECT id, uuid, name FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := Driver{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}
/*
// Get a single user given the email
func UserByEmail(email string) (user Driver, err error) {
	user = Driver{}
	err = DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.CreatedAt)
	return
}*/

// Get a single user given the UUID
func UserByUUID(uuid string) (user Driver, err error) {
	user = Driver{}
	err = DB.QueryRow("SELECT id, uuid, name FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name)
	return
}

/*
// This function gets a slice of *Driver from Database and returns it as JSON
func (db *DB) GetAllDrivers() ([]*Driver, error) {
	db, err := GetPgConnection()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	drivers := make([]*Driver, 0)

	for rows.Next() {
		driver := new(Driver) // Allocate memory for Driver and returns pointer to it
		err = rows.Scan(&driver.Id, &driver.Username)

		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil
}

// This function gets a pointer to driver struct which has the driver id information and returns it as JSON
func (db *DB) GetDriverById(id int) (*Driver, error) {
	db, err := GetPgConnection()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	driver := new(Driver)

	for row.Next() {
		err = row.Scan(&driver.Id, &driver.Username)

		if err != nil {
			return nil, err
		}
	}

	return driver, nil
}*/
