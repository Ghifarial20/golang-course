# MyGram Project

## Description

Program ini merupakan final project untuk Kelas Scalable Web Service with Golang - DTS Kominfo

## Cara Menjalankan

- clone github MyGram
- create database postgre dengan nama "mygram"
- buka terminal
- ketikkan "go run main.go"

## Package yang digunakan

- go get -u github.com/gin-gonic/gin
- go get -u gorm.io/driver/postgres
- go get -u gorm.io/gorm
- go get github.com/asaskevich/govalidator
- go get github.com/dgrijalva/jwt-go
- go get golang.org/x/crypto/bcrypt

## How To Test

### Via Postman

- Buka [link ini](./Mygram.postman_collection.json) dan simpan sebagai json
- Import file json pada Postman
- Lakukan tes pada setiap Endpoint

### Link To Test

#### Users

| Name Endpoint | Endpoint                             | Method | Details                              |
| ------------- | ------------------------------------ | ------ | ------------------------------------ |
| Register User | http://localhost:5000/users/register | POST   | [Open](./API-SPEC.md#register-users) |
| Login User    | http://localhost:5000/users/login    | POST   | [Open](./API-SPEC.md#login-users)    |
| Update User   | http://localhost:5000/users/{id}     | PUT    | [Open](./API-SPEC.md#update-users)   |
| Delete User   | http://localhost:5000/users/{id}     | DELETE | [Open](./API-SPEC.md#delete-users)   |

#### Photos

| Name Endpoint | Endpoint                          | Method | Details                             |
| ------------- | --------------------------------- | ------ | ----------------------------------- |
| Create photo  | http://localhost:5000/photos      | POST   | [Open](./API-SPEC.md#create-photos) |
| Get photos    | http://localhost:5000/photos      | GET    | [Open](./API-SPEC.md#get-photos)    |
| Update photo  | http://localhost:5000/photos/{id} | PUT    | [Open](./API-SPEC.md#update-photos) |
| Delete photo  | http://localhost:5000/photos/{id} | DELETE | [Open](./API-SPEC.md#delete-photos) |

#### Comments

| Name Endpoint  | Endpoint                            | Method | Details                               |
| -------------- | ----------------------------------- | ------ | ------------------------------------- |
| Create comment | http://localhost:5000/comments      | POST   | [Open](./API-SPEC.md#create-comments) |
| Get comments   | http://localhost:5000/comments      | GET    | [Open](./API-SPEC.md#get-comments)    |
| Update comment | http://localhost:5000/comments/{id} | PUT    | [Open](./API-SPEC.md#update-comments) |
| Delete comment | http://localhost:5000/comments/{id} | DELETE | [Open](./API-SPEC.md#delete-comments) |

#### Social Media

| Name Endpoint       | Endpoint                                | Method | Details                                    |
| ------------------- | --------------------------------------- | ------ | ------------------------------------------ |
| Create social media | http://localhost:5000/socialmedias      | POST   | [Open](./API-SPEC.md#create-social-medias) |
| Get social medias   | http://localhost:5000/socialmedias      | GET    | [Open](./API-SPEC.md#get-social-medias)    |
| Update social media | http://localhost:5000/socialmedias/{id} | PUT    | [Open](./API-SPEC.md#update-social-medias) |
| Delete social media | http://localhost:5000/socialmedias/{id} | DELETE | [Open](./API-SPEC.md#delete-social-medias) |


# About Me

- Nama: Ghifari Ahmad Lustiansyah
- Kode Peserta: 149368582101-594
