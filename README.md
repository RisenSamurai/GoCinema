# GoCinema

GoCinema is a modern web application designed to provide users with an immersive cinema experience. Built with a robust backend using Golang and the Gin framework, and a dynamic frontend leveraging SvelteKit and TailwindCSS, GoCinema offers seamless browsing and management of movie content.

## Features

- **User Authentication**: Secure user registration and login functionalities.
- **Movie Browsing**: Explore a curated list of movies with detailed information.
- **Search Functionality**: Quickly find movies using the integrated search feature.
- **Responsive Design**: Optimized for various devices, ensuring a consistent user experience.

## Tech Stack

- **Backend**: [Golang](https://golang.org/) with [Gin](https://gin-gonic.com/) framework
- **Frontend**: [SvelteKit](https://kit.svelte.dev/) with [TailwindCSS](https://tailwindcss.com/)
- **Database**: [NoSQL Database] (MongoDB)
- **API Documentation**: [Swagger](https://swagger.io/) for API documentation

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) installed
- [Node.js](https://nodejs.org/) installed
- [MongoDB](https://www.mongodb.com/) or your chosen NoSQL database set up

### Backend Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/RisenSamurai/GoCinema.git
   cd GoCinema/backend

2.	Install dependencies:

   ``` go mod tidy
3.	Set up environment variables (create a .env file):

  PORT=8080
  DB_URI=your_database_uri
  JWT_SECRET=your_jwt_secret

### Frontend Setup

1.	Navigate to the frontend directory:
  cd ../frontend
2.	Install dependencies:
  npm install

### Usage
  1.	Access the frontend at http://localhost:3000.
	2.	Register a new account or log in with existing credentials.
	3.	Browse the movie catalog and enjoy the features of GoCinema.

### License
This project is licensed under the MIT License. See the LICENSE file for details.

