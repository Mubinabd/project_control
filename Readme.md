# Project Control

**Project Control** is a CRM-like platform designed to simplify the management of personal and team projects. The platform acts as a centralized hub, allowing team leaders and developers to monitor project progress, manage documentation, and communicate effectively.

## Key Features

1. **API Documentation Management**: 
   - Manage and share API documentation for each project through a dedicated Swagger URL. This feature ensures that team members have quick access to the API documentation relevant to the project they are working on.
   
2. **Contact Information Management**: 
   - Store and manage phone numbers and Telegram usernames, facilitating easy communication within the team. This feature helps improve collaboration and communication between team members by keeping all contact information in one place.
   
3. **Document Management**: 
   - Centralized management of documents related to each project. Users can add, update, and delete documents, ensuring that all project-related information is stored and easily accessible in one location.
   
4. **Team Member Management**: 
   - Team leaders can manage developers, store their contact information, and assign tasks related to projects. This allows for more efficient task delegation and communication within the team.

## Getting Started

### Prerequisites

- **Golang** (version 1.19 or higher)
- **Docker** (for containerization)
- **PostgreSQL** (for database management)
- **Swagger** (for API documentation)
  
### Installation

1. Clone the repository:

   ```bash
   git clone https://gitlab.com/your-repo/project-control.git

2. Navigate to the project directory:
    ```bash
    cd project-control

3. Start the project with Docker Compose:

    ```bash
    docker-compose up --build