package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct{
}

// agar semua kontrak bisa ke import
func NewService() *jwtService{
	return &jwtService{}
}

var SECRET_KEY = []byte("BWASTARTUP_SUCCESS")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// buat payloadnya
	claim := jwt.MapClaims{}
	// isi dari payload
	claim["user_id"] = userID

	// buat token jwt berserta algoritma yg digunakan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// buat verify signature/ secret key
	signedToken, err := token.SignedString(SECRET_KEY)
	
	if err != nil {
		return signedToken, err
	}
	
	return signedToken, nil
}