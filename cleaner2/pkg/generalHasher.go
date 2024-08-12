package pkg

type GeneralHasher interface{
	HashPassword(password string)
	CompareHashAndPassword(password,hash string)
}