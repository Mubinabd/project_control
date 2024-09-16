# Project Control

**Project Control** is a CRM-like platform designed to simplify the management of personal and team projects. Acting as a centralized hub, the platform enables team leaders and developers to efficiently monitor project progress, manage documentation, and communicate seamlessly.

## Key Features

1. **API Documentation Management**  
   Easily manage and share API documentation for each project via a dedicated Swagger URL. This allows team members to quickly access all relevant API docs.

2. **Contact Information Management**  
   Store and organize phone numbers and Telegram usernames, making it easier to communicate within the team. Keep all contact details centralized for efficient collaboration.

3. **Document Management**  
   Centralized storage for project-related documents. Add, update, and delete files as needed, ensuring all important information is stored in one accessible location.

4. **Team Member Management**  
   Team leaders can manage developers, maintain contact information, and assign project-related tasks. This simplifies task delegation and enhances team collaboration.

## Getting Started

To get started with **Project Control**, follow these steps:

### Prerequisites

Make sure you have the following tools installed:

- **Golang** (version 1.19 or higher)
- **Docker** (for containerization)
- **PostgreSQL** (for database management)
- **Swagger** (for API documentation)

### Installation

1. Clone the repository:

    ```bash
    git clone https://gitlab.com/your-repo/project-control.git
    ```

2. Navigate to the project directory:

    ```bash
    cd project-control
    ```

3. Start the project using Docker Compose:

    ```bash
    docker-compose up --build
    ```

4. Once the project is running, you can access the API documentation at:

    ```bash
    http://localhost:8080/swagger/index.html
    ```

## Project Structure

Here's a breakdown of the main components of the project:

├── cmd │ └── main.go # Main entry point for the application ├── internal │ ├── api # API routes and handlers │ ├── docs # Swagger documentation files │ ├── models # Database models │ └── services # Core business logic ├── pkg │ └── utils # Utility functions and helpers ├── migrations # Database migration files ├── Dockerfile # Docker container configuration ├── docker-compose.yml # Docker Compose configuration ├── .env.example # Example environment variables └── README.md # Project documentation


## Usage

- **API Documentation**: Access and manage project-specific API documentation through Swagger URLs.
- **Contacts Management**: Store and retrieve team members' contact information, including phone numbers and Telegram usernames.
- **Document Management**: Upload, view, and manage project-related documents in one place.
- **Team Management**: Manage team members, assign tasks, and oversee project progress.

## Contributing

We welcome contributions! To get involved:

1. Fork the repository.
2. Create a new feature branch: `git checkout -b feature/new-feature`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/new-feature`.
5. Open a Pull Request for review.

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

---

**Project Control** – making project management easier, more efficient, and collaborative.
