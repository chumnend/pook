# Pook: A storybook creation application

## About

Pook is a web platform that allows users to create fun "picture" books to share with others.

### Built With

- Go v1.23

## Getting Started

To get started with the application, follow these steps:

1. **Clone the repository:**

   ```sh
   git clone git@github.com:chumnend/pook.git
   cd pook
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Setup .env file.**

   ```sh
   cp .env.example .env
   vi .env
   ```

4. Install `golang-migrate`. This is for Linux, if using Mac or Windows please look up installation guide.

   ```sh
   curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
   sudo mv migrate /usr/local/bin/
   ```

5. **Run database migrations**

   ```sh
   make migrate
   ```

6. **Run the application:**

   ```sh
   go run cmd/main.go
   ```

7. **Access the application:**

   Open your web browser and navigate to `http://localhost:8080` to see the application in action.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

Nicholas Chumney - [nicholas.chumney@outlook.com](nicholas.chumney@outlook.com)
