package goMod2

func GetDB() string {
	return "GetDB - работает из private"
}

func SetDB(dbname string) (string, int) {
	err := 333
	return dbname, err
}
