# Test BE PrivyID

## Deskripsi 
Membuat CRUD Api untuk cake store

## Stack 
- Golang(go1.19.4)
- GorillaMUX
- Database: MySQL

## Start Project
- go run ./app/main.go

## Fungsional
Silahkan Import File Json yang ada pada folder Postman ke Postman

## Endpoint
- Add-Cake : /cakes (method POST)
- Detail-Cake : /cakes/:id (method GET)
- List-Of-Cake : /cakes (method GET)
- Update-Cake : /cakes/:id (method PATCH)
- Delete-Cake : /cakes/:id (method DELETE)

## Point Test Yang Belum Dapat di Selesaikan
- Provide unit test on your project
- Running in docker contain