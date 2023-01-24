# Backend Service API Property Marketplace (Air BnB Clone App)

![air-bnb-clone-app-logo](https://user-images.githubusercontent.com/112921258/214287286-83b65329-4b2d-4000-a9ba-02917dcc5078.png)

## What is Property Marketplace?

A year after pandemic is over, want go on vacation? Reserve homestay for your vacation using [Property Marketplace App](https://feproject3kel3-7j4c7hclm-airbbnb.vercel.app/user/myproperties)! Many place to stay here, just choose, check availbility and book! So simple! That awesome app was powered with this API.
This project is about the API Program for Property Marketplace using powerful languange, Golang. Build in Docker Container and also using Mysql at Docker Container.  This API will be consumed by Air BnB Alta App that our Frontend team build. Checkout more about this API App below.

## Database Design

Database design using Entity Relationship Diagram.

![ERD API](https://user-images.githubusercontent.com/112921258/214288354-fb647434-9b94-4240-a0d2-1b55d887533b.jpg)

## Open API Documentation

Explore our API at this [Documentation](https://app.swaggerhub.com/apis-docs/ACHMADQIZWINI4_1/GP3_Kelompok3/1.0.0) powered by [Swagger.io](https://swagger.io/).

## Postman Collection API

Test the API using with Postman Collection ath this [Property-Marketplace-AirBnBCloneApp.postman_collection v6.zip](https://github.com/Property-Marketplace-Air-BnB-Clone-App/Property-Marketplace-Air-BnB-Clone-App-Backend-Service/files/10489410/Property-Marketplace-AirBnBCloneApp.postman_collection.v6.zip)

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
