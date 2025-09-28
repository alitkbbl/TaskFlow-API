# TaskFlow API 🚀

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![Gin Framework](https://img.shields.io/badge/Gin-Framework-green)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue)
![JWT](https://img.shields.io/badge/JWT-Auth-orange)
![License](https://img.shields.io/badge/license-MIT-blue)

**A RESTful Task Management API with Advanced Features and High Security**

[Features](#features) • [Installation](#installation) • [API Documentation](#api-documentation) • [Project Structure](#project-structure) • [Contributing](#contributing)

</div>

## 📖 Overview

TaskFlow API is a personal and team task management service built with **Golang** and **Gin Framework**. This project implements modern architecture and follows software development best practices.

## ✨ Features

### 🔐 Authentication System
- User registration and login
- JWT Token-based Authentication
- Password hashing with bcrypt
- Security middleware

### 📝 Task Management
- Create, edit, delete, and view tasks
- Task status updates (TODO, IN_PROGRESS, DONE)
- Task categorization
- Start and due dates

### 🔍 Search & Filtering
- Advanced search in titles and descriptions
- Filter by status, priority, and category
- Pagination and sorting

### 🛡️ Security
- Rate Limiting
- CORS Configuration
- Input Validation
- SQL Injection Prevention

### 🚀 Advanced Features
- Due dates and reminders
- Task prioritization (LOW, MEDIUM, HIGH, URGENT)
- Task comments and discussions
- File attachments
- Task history and audit logs
- Bulk operations

## 🚀 Installation

### Prerequisites
- Go 1.21 or higher
- PostgreSQL 12 or higher
- Git

### Installation Steps

1. **Clone the Repository**
```bash
git clone https://github.com/your-username/taskflow-api.git
cd taskflow-api
