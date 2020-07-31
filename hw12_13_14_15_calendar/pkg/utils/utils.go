package utils

import "fmt"

func GetDSN(host, port, user, password, name string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
}
