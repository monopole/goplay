module github.com/monople/goplay/fruituser

go 1.21.2

replace (
	github.com/monopole/goplay/apple => ../apple
	github.com/monopole/goplay/banana => ../banana
)

require (
	github.com/monopole/goplay/apple v0.0.0-00010101000000-000000000000 // indirect
	github.com/monopole/goplay/banana v0.0.0-00010101000000-000000000000 // indirect
)
