package constants

const MinUsernameLength = 3
const MinPasswordLength = 12
const SessionExpiry = 1440 // minutes until session expires

// ------- users table ---------
const CheckUniqueEmail = "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
const CheckUniqueUserName = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
const CreateUser = "INSERT INTO users (id, username, email, password_hash) VALUES (?, ?, ?, ?)"
const GetUserIDAndPasswordByEmail = "SELECT id, password_hash FROM users WHERE email = ?"

// ------- sessions table ---------
const AddSession = "INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?)" // expires_at syntax: YYYY-MM-DD HH:MM:SS
const GetExpiryTime = "SELECT expires_at FROM sessions WHERE id = ?"
