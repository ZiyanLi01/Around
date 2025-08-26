# PandaPost - Social Networking Platform

PandaPost is a dynamic social networking platform designed to enhance user interactions through seamless content sharing, AI-powered features, and robust search capabilities. Built with modern technologies, Around provides a scalable and user-friendly experience for creating, sharing, and discovering content.

## Project Demo
https://github.com/user-attachments/assets/0555652b-80d4-4ad7-8dfd-b751ff97fa45

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Deployment](#deployment)
- [Contributing](#contributing)

## Overview

PandaPost is a full-stack social networking application that combines modern web technologies with AI capabilities. The platform allows users to create posts with images and videos, generate AI-powered content, and discover posts through advanced search functionality. Built with scalability in mind, Around handles high user activity while maintaining performance and security.

## Features

### Core Functionality
- **User Authentication**: Secure JWT-based authentication system with signup and signin capabilities
- **Content Creation**: Create and share posts with images and videos
- **Content Discovery**: Browse and search through posts with advanced filtering options
- **User Profiles**: Personalized user experience with individual post management

### Advanced Features
- **AI Image Generation**: Integration with OpenAI DALL-E 3 for AI-generated images
- **Elasticsearch Integration**: Fast and accurate search functionality for posts and users
- **Google Cloud Storage**: Scalable file storage for images and videos
- **Responsive Design**: Modern UI built with Ant Design components

### Search Capabilities
- **Global Search**: Search across all posts in the platform
- **User-specific Search**: Find posts by specific users
- **Keyword Search**: Search posts by content keywords
- **Real-time Results**: Instant search results with Elasticsearch

## Architecture

The application follows a modern client-server architecture:

- **Frontend**: React.js single-page application with Ant Design UI components
- **Backend**: Go-based REST API with Google App Engine deployment
- **Database**: Elasticsearch for search functionality and data storage
- **Storage**: Google Cloud Storage for media files
- **Authentication**: JWT-based token authentication

## Technology Stack

### Frontend
- **React 18.2.0**: Modern JavaScript library for building user interfaces
- **Ant Design 4.24.15**: Enterprise-level UI design language and React UI library
- **React Router 6.26.1**: Declarative routing for React applications
- **Axios 1.6.2**: Promise-based HTTP client for API communication
- **React Grid Gallery**: Component for displaying image galleries

### Backend
- **Go 1.21**: High-performance programming language
- **Gorilla Mux**: Powerful HTTP router and URL matcher
- **JWT Authentication**: Secure token-based authentication
- **Google App Engine**: Scalable cloud platform for deployment
- **Elasticsearch**: Distributed search and analytics engine
- **Google Cloud Storage**: Object storage for media files

### External Services
- **OpenAI DALL-E 3**: AI image generation service
- **Google Cloud Platform**: Cloud infrastructure and services

## Project Structure

```
Around/
├── around-web/                 # Frontend React application
│   ├── src/
│   │   ├── components/        # React components
│   │   │   ├── App.js         # Main application component
│   │   │   ├── Home.js        # Home page with post display
│   │   │   ├── Login.js       # User authentication
│   │   │   ├── Register.js    # User registration
│   │   │   ├── ImageGenerator.js # AI image generation
│   │   │   ├── PhotoGallery.js   # Image gallery display
│   │   │   ├── SearchBar.js   # Search functionality
│   │   │   └── ...
│   │   ├── constants.js       # Application constants
│   │   └── ...
│   ├── package.json           # Frontend dependencies
│   └── ...
├── go/                        # Backend Go application
│   └── src/around/
│       ├── main.go            # Application entry point
│       ├── handler/           # HTTP request handlers
│       │   ├── router.go      # Route definitions
│       │   ├── user.go        # User authentication handlers
│       │   ├── post.go        # Post management handlers
│       │   └── image.go       # Image generation handlers
│       ├── backend/           # Backend services
│       │   ├── elasticsearch.go # Elasticsearch integration
│       │   └── gcs.go         # Google Cloud Storage integration
│       ├── model/             # Data models
│       │   └── model.go       # Post and User structures
│       ├── app.yaml           # Google App Engine configuration
│       └── go.mod             # Go module dependencies
└── README.md                  # Project documentation
```

## Getting Started

### Prerequisites
- Node.js (v16 or higher)
- Go (v1.21 or higher)
- Google Cloud Platform account
- Elasticsearch instance
- OpenAI API key

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd around-web
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm start
   ```

4. The application will be available at `http://localhost:3000`

### Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd go/src/around
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables for:
   - Google Cloud credentials
   - Elasticsearch connection
   - OpenAI API key
   - JWT signing key

4. Run the application locally:
   ```bash
   go run main.go
   ```

5. The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication
- `POST /signup` - User registration
- `POST /signin` - User login

### Content Management
- `POST /upload` - Upload new post (requires authentication)
- `GET /search` - Search posts (requires authentication)
- `DELETE /post/{id}` - Delete post (requires authentication)

### AI Features
- `POST /api/generate-image` - Generate AI image (requires authentication)
- `GET /download-image` - Download generated image (requires authentication)

## Deployment

### Frontend Deployment
The frontend is configured for static hosting and can be deployed to:
- Google Cloud Storage
- Netlify
- Vercel
- Any static hosting service

### Backend Deployment
The backend is configured for Google App Engine deployment:

1. Ensure you have the Google Cloud SDK installed
2. Configure your project in `app.yaml`
3. Deploy using:
   ```bash
   gcloud app deploy
   ```

### Environment Configuration
Set up the following environment variables in your deployment:
- `ELASTICSEARCH_URL`: Your Elasticsearch instance URL
- `GOOGLE_CLOUD_PROJECT`: Your Google Cloud project ID
- `OPENAI_API_KEY`: Your OpenAI API key
- `JWT_SIGNING_KEY`: Secret key for JWT token signing

## Contributing

This is a personal project showcasing modern web development practices. The codebase demonstrates:

- Full-stack development with React and Go
- Cloud-native architecture with Google Cloud Platform
- AI integration with OpenAI services
- Scalable search functionality with Elasticsearch
- Modern authentication and security practices
- Responsive and accessible user interface design

## License

This project is developed as a personal portfolio piece to demonstrate full-stack development capabilities and modern web application architecture.
