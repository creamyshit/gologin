package helper

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"time"

	"github.com/creamyshit/gologin/model"

	"github.com/gofiber/fiber/v2"
)

type authResponse struct {
	Result       interface{} `json:"result"`
	Message      string      `json:"message"`
	Success      bool        `json:"success"`
	SessionToken string      `json:"sessiontoken"`
}

func AuthResponse(c fiber.Ctx, success bool, data interface{}, message string, code int, token string) error {
	return c.Status(code).JSON(&authResponse{
		Success:      success,
		Result:       data,
		Message:      message,
		SessionToken: token,
	})
}

func subtractTime(time1, time2 time.Time) float64 {
	diff := time2.Sub(time1).Seconds()
	return diff
}

// Define salt size
const saltSize = 16

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a base64 encoded string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// Check if two passwords match
func DoPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}

func HideCredential(a *model.User) *model.User {
	a.Password = ""
	a.Salt = nil
	return a
}
