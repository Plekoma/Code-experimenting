module local/main

go 1.27.1

replace CRS => ./CRS

require CRS v0.0.0-00010101000000-000000000000 // direct

require github.com/joho/godotenv v1.5.1 // indirect
