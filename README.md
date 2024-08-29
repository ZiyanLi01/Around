# Welcome to Panda Post!
This is a dynamic social networking platform I've passionately developed to enhance user interactions in a seamless and scalable manner. This project is designed to handle the challenges of high user activity, provide an enriched user experience with AI-generated content, and offer fast, accurate search capabilities. Whether you're posting, browsing, or connecting with others, Around is built to support your needs efficiently and securely.

## Project Demo


## Table of Contents
- [Introduction](#introduction)
- [Backend Design](#backend-design)
- [Frontend Design](#frontend-design)
- [Search and AI Integration](#search-and-ai-integration)
- [Authentication and Security](#authentication-and-security)
- [Reference](#reference)
  
## Introduction
Panda Post is structured with robust backend and frontend components, designed to ensure smooth user interactions and scalability. This document provides a detailed overview of these components, highlighting how they work together to create a responsive and user-friendly social networking platform.

## Backend Design
The backend is developed using Go and hosted on Google App Engine, focusing on scalability and efficiency.

Key Features:

Post Handling: Efficiently manage and scale user posts, ensuring a smooth experience even as user activity grows.

ElasticSearch Integration: Allows for fast and accurate searching of posts, enabling users to find and list posts effortlessly.

The backend code is available in the "go" folder.

## Frontend Design
The frontend is developed with React.js, ensuring a user-friendly interface that delivers a seamless interaction experience.

Key Features:

Post Creation and Management: Users can create and manage their posts with ease, with the interface designed for smooth navigation.

React Router v4: Ensures dynamic and responsive routing throughout the application.

Frontend files are located in the "around-web" folder.

## Search and AI Integration
To enhance user experience, Around integrates advanced technologies:

AI-Generated Content:
OpenAI DALL-E 3: Integrated to assist users in creating and updating posts with AI-generated images, making posts more engaging.

Enhanced Search:
ElasticSearch: Enables quick and accurate search functionality, allowing users to find recent posts and list their personal posts with ease.

## Authentication and Security
Panda Post prioritizes security with a comprehensive authentication system:

Key Features:

Token-Based Authentication: Utilizes server-side JWT for secure registration, login, and logout processes.

React Router v4: Ensures secure and efficient navigation for users throughout the application.

## Reference
Panda Post is a personal project inspired by modern social networking needs. It combines my passion for creating scalable, user-friendly applications with advanced technologies to enhance the user experience. Feel free to explore the repository and see how Around connects people through efficient, secure, and interactive features.
