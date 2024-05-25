# StudentFest API

Welcome to StudentFest API - Your Ultimate Event Management Solution! 🎉

StudentFest API is a robust REST API built from scratch using GoLang and the powerful GIN framework. It's designed to streamline event management for college fests, offering seamless integration, authentication, and SQL database support, making event organization a breeze for students and organizers alike.

## Key Features

- **Authentication**: Secure your events with JWT token-based authentication, ensuring only authorized users can access and manage event data.
- **SQL Database Integration**: Harness the power of SQL databases to store and manage event information efficiently, enabling quick retrieval and seamless data manipulation.
- **Event Management**: Easily create, update, delete, and retrieve event details, including dates, venues, participants, and more, all through intuitive RESTful endpoints.
- **Scalability**: Built with scalability in mind, StudentFest API can effortlessly handle growing event data and user demands, ensuring smooth performance even during peak times.

## Getting Started

To get started with StudentFest API, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/soshal/StudentFest-API.git
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up your database configuration in `config.yaml`.

4. Run the application:

```bash
go run main.go
```

5. Access the API endpoints as documented in the [API Documentation](#api-documentation).

## API Documentation

The API endpoints and their usage are documented in the [API Documentation](API_DOCUMENTATION.md) file. Please refer to this document for detailed information on how to interact with the StudentFest API.

## Contributing

Contributions are welcome! If you'd like to contribute to StudentFest API, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/your-feature`).
6. Create a new Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize this README.md file further to include any additional information or sections relevant to your project!
