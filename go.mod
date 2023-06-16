module github.com/GreenMan-Network/Go-TokenID

go 1.20

replace github.com/GreenMan-Network/Go-KMS => ../Go-KMS

replace github.com/GreenMan-Network/Go-Utils => ../Go-Utils

require (
	github.com/GreenMan-Network/Go-KMS v0.0.0-00010101000000-000000000000
	github.com/GreenMan-Network/Go-Utils v0.0.0-20230611180231-8328adb80148
	github.com/go-jose/go-jose/v3 v3.0.0
)

require golang.org/x/crypto v0.9.0 // indirect
