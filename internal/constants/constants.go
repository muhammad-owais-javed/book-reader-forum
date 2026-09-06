package constants

const MinUsernameLength = 3
const MinPasswordLength = 12
const SessionExpiry = 24

// ------- users table ---------
const CheckUniqueEmail = "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
const CheckUniqueUserName = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
const CreateUser = "INSERT INTO users (id, username, email, password_hash) VALUES (?, ?, ?, ?)"
const GetUserIDAndPasswordByEmail = "SELECT id, password_hash FROM users WHERE email = ?"

// ------- sessions table ---------
const AddSession = "INSERT INTO sessions (uuid, user_id, expires_at) VALUES (?, ?, ?)"
