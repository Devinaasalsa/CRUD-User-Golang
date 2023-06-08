package config

import "time"

var (
	ExpiresTime = time.Now().Add(time.Hour * 24).Unix() //Set expire time jwt to 1 day (24 hour)
	SignatureKey = "eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY4NDc0NzMwNSwiaWF0IjoxNjg0NzQ3MzA1fQ.ryRIUbZr6zzsoVTSyzu-j29g_2UzX0ywn_KrFVUZwKI" //ES256 ALGORITHM
)