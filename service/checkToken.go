package service

//type AccesDetails struct {
//	AccessUuid string
//	UserId     uint64
//}
//
//func ExtractToken(r *http.Request) string {
//	bearToken := r.Header.Get("Authorization")
//	strArr := strings.Split(bearToken, " ")
//	if len(strArr) == 2 {
//		return strArr[1]
//	}
//
//	return ""
//}
//
//func VerifyToken(r *http.Request) (*jwt.Token, error) {
//	tokenString := ExtractToken(r)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return environment.ViperEnvVariable("JWT_SECRET"), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return token, nil
//}
//
//func TokenValid(r *http.Request) error {
//	token, err := VerifyToken(r)
//	if err != nil {
//		return err
//	}
//	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
//		return err
//	}
//	return nil
//}

//func ExtractTokenMetadata(r *http.Request) *AccesD {
//
//}
