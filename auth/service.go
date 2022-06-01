package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error){
	// parse encodeToken
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error){
		// cek apakah algoritma yang digunakan sama tidak dengan yg dikirim
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}