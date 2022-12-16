# API Air BnB Alta

![AirBNB Alta](https://trialtestbucket.s3.amazonaws.com/profile/rAzuOVHDUfpVTw5UBPj1-sultan%20thumbnail.jpeg)

## What is API Air Bnb Alta?

A year after pandemic is over, want go on vacation? Reserve homestay for your vacation using [Air BnB Alta App](https://feproject3kel3-7j4c7hclm-airbbnb.vercel.app/user/myproperties)! Many place to stay here, just choose, check availbility and book! So simple! That awesome app was powered with this API.
This project is about the API Program for Air BnB Alta using powerful languange, Golang. Build in Docker Container and also using Mysql at Docker Container. This project as part of the requirements to pass Phase 2 of the [Alterra Academy’s](https://academy.alterra.id/) Immersive Program. This API will be consumed by Air BnB Alta App that our Frontend team build. Checkout more about this API App below.

## Open API Documentation

Explore our API at this [Documentation](https://app.swaggerhub.com/apis-docs/ACHMADQIZWINI4_1/GP3_Kelompok3/1.0.0) powered by [Swagger.io](https://swagger.io/).

## Database Design

Database design using Entity Relationship Diagram.
![ERD](https://raw.githubusercontent.com/Fase2-Project3-Group3/api-airbnb-alta/main/erd_api.jpg)

## Tech Stack

#### Code written based on

- Go Programming Language
- Echo Framework
- Gorm Package
- Mysql Database

#### Project Structure

Clean Architechture

#### Development

- Git Trunk Based Developepment
- Git Worklow CI/CD
- JWT Authentication
- Unit Testingn

#### Deployment

- Cloud Server Google Cloud Platform
- Docker Container
- File Repository Amazon S3 Service
- Cloudflre HTTPS/SSL

We will be performing the CRUD operation for Book with few fields.

## Feature

- Auth (Login)
- Users CRUD : User management. Theres user that using app for booking as Guest and other user as Host who rent the property.
- Properties CRUD : Properties management. Public can check all property and detail of the property. As Host, user that have property, can add property with image gallery. Host can check list of all their properties. Property have image gallery and comments. Endpoint thor thoose made separated for flexibility access.
- Property Images CRUD. Host can add images to fullfill their gallery with this feature.
- Bookings CRUD : Guest can choose property they like and check availbility of check in checkout date. If available, guest can Reverse or Book the property. Guest also can check list of booking history.
- Comments CRUD : Ater Guest done booking a property, Guest can give rating and comment.

## How to install and Run the Project

1.  Git clone this project
2.  Create mysql database
3.  Setup Environtment (check detail below)
4.  Open the main folder with your vs code
5.  Once it opened, open terminal in your vs code
6.  Type "go run main.go"
7.  Enjoy the app.

## Environtment Setup

Create a file name ".env" and put at root folder of this project. Write this environtment detail below :

```
export DB_USERNAME=""
export DB_PASSWORD=""
export DB_HOST=""
export DB_PORT=""
export DB_NAME=""
export SERVER_PORT=""
export JWT_SECRET=""
export AWS_REGION=""
export AWS_BUCKET_NAME=""
export ACCESS_KEY_IAM=""
export SECRET_KEY_IAM=""
```

## Credits

Created by [Faishal](https://github.com/mfaishal882) and [Ahmad Qizwini](https://github.com/Achmadqizwini/)
Thanks to our mentor [Fakhri](https://github.com/iffakhry)
