package constants

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
)

const MinUsernameLength = 3
const MinPasswordLength = 12
const SessionExpiry = 1440  // minutes until session expires
const ResetTokenExpiry = 15 // minutes until reset token expires

// ------- users table ---------
const CheckUniqueEmail = "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
const CheckUniqueUserName = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
const CreateUser = "INSERT INTO users (id, username, email, password_hash) VALUES (?, ?, ?, ?)"
const GetUserIDAndPasswordByEmail = "SELECT id, password_hash FROM users WHERE email = ?"
const GetUserIDWithEmail = "SELECT id FROM users WHERE email = ?"

// ------- sessions table ---------
const AddSession = "INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?)" // expires_at syntax: YYYY-MM-DD HH:MM:SS
const GetExpiryTime = "SELECT expires_at FROM sessions WHERE id = ?"
const DeleteSession = "DELETE FROM sessions WHERE id = ?"

// ------- reset_tokens table ---------
const AddResetToken = "INSERT INTO reset_tokens (reset_token, user_id, expires_at) VALUES (?, ?, ?)" // expires_at syntax: YYYY-MM-DD HH:MM:SS
const GetResetTokenExpiryTime = "SELECT expires_at FROM reset_tokens WHERE reset_token = ?"
const GetResetTokenDetails = `SELECT user_id, expires_at FROM reset_tokens WHERE reset_token = ?`
const DeleteResetToken = "DELETE FROM reset_tokens WHERE reset_token = ?"
