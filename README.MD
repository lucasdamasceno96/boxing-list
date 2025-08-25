Aqui está a tradução completa para inglês:

---

# Boxing API 🥊

A complete RESTful API for managing boxers, rankings, and fight analysis, built with Go (Golang), Gin, and SQLite. This project demonstrates best practices in backend development, including clean architecture, testing, CI/CD, and integration with cloud services and AI.

## 🌟 About the Project

**Boxing API** was created as a portfolio project to showcase advanced skills in the Go ecosystem. The goal is to provide a robust and efficient backend for a boxing application—whether it’s an informational website, a mobile app, or a fantasy game.

The choice of **Go + Gin** was driven by its high performance and low resource consumption, while **SQLite** was chosen as a lightweight, low-cost database solution for deployment, ideal for small to medium-sized projects, eliminating the need for a separate database server.

## ✨ Key Features

* **Full Boxer CRUD:** Create, list, update, and delete boxers.
* **Flexible Search:** Find fighters by UUID or full name.
* **Image Management:** Upload fighter images with size standardization and cloud storage on **Azure Blob Storage**.
* **Fight History:** Store a detailed record of all a boxer’s fights.
* **Dynamic Ranking System:** A dedicated endpoint that calculates and returns fighter rankings based on a customizable point system.
* **Weight Class Rules:** Business logic that prevents scheduling fights between boxers with more than a 3-weight-class difference.
* **AI Fight Analysis:** An innovative endpoint that uses a Large Language Model (LLM) to generate predictive analysis of the outcome of a fight between two boxers.
* **API Documentation:** Endpoints documented with Swagger for easy API consumption.

## 🛠️ Tech Stack

| Category             | Technology/Tool                          |
| -------------------- | ---------------------------------------- |
| **Language**         | Go (Golang)                              |
| **Web Framework**    | Gin Gonic                                |
| **Database**         | SQLite                                   |
| **ORM**              | GORM                                     |
| **Containerization** | Docker & Docker Compose                  |
| **Storage**          | Azure Blob Storage                       |
| **AI**               | Integration with OpenAI API (or similar) |
| **Testing**          | Go `testing` & `httptest`                |
| **Environment Vars** | Viper & `.env` files                     |
| **Documentation**    | Swaggo (Swagger for Gin)                 |

## 🏗️ Architecture

The project follows a **Clean Architecture**-inspired structure, separating responsibilities into distinct layers to ensure cleaner, more testable, and maintainable code.

```
boxing-api/
├── cmd/api/             # Application entry point
├── config/              # Configuration management
├── internal/
│   ├── api/             # Presentation layer (Handlers, Routes, Middlewares)
│   ├── domain/          # Business logic and entities (Models, Services)
│   ├── infrastructure/  # Data layer (Repositories, External Clients)
├── test/                # Unit and integration tests
├── .env                 # Environment variables file
└── Dockerfile           # Application container definition
```

## 🔌 API Endpoints

The API is versioned to ensure future compatibility. The base prefix for all endpoints is `/v1`.

| HTTP Method | Endpoint                         | Description                                       |
| :---------- | :------------------------------- | :------------------------------------------------ |
| `POST`      | `/boxers`                        | Creates a new boxer.                              |
| `GET`       | `/boxers`                        | Lists all boxers.                                 |
| `GET`       | `/boxers/{id}`                   | Retrieves a boxer by UUID.                        |
| `GET`       | `/boxers/search?name={fullName}` | Searches for a boxer by name.                     |
| `PUT`       | `/boxers/{id}`                   | Updates a boxer’s data.                           |
| `DELETE`    | `/boxers/{id}`                   | Deletes a boxer.                                  |
| `GET`       | `/ranking`                       | Returns the list of boxers sorted by points.      |
| `POST`      | `/fight-analysis`                | Analyzes a hypothetical fight between two boxers. |
| `GET`       | `/swagger/*any`                  | Displays API documentation via Swagger UI.        |



## 🚀 Highlighted Logic

#### Ranking Point System

The ranking is dynamically calculated based on the following rules:

| Event                          | Points     | Notes                                                |
| ------------------------------ | ---------- | ---------------------------------------------------- |
| **Per Fight**                  | 1 Point    | Rewards the fighter’s activity.                      |
| **Per Draw**                   | 2 Points   | A neutral result, but better than a loss.            |
| **Per Win**                    | 5 Points   | The foundation of success.                           |
| **Bonus: Win by Knockout**     | +2 Points  | Rewards a dominant victory (total of 7 points).      |
| **Bonus: Title Fight Win**     | +10 Points | A high-level achievement (total of 15 or 17 points). |
| **Bonus: Draw in Title Fight** | +5 Points  | The champ still champ.                               |

---



## 🤖 AI Fight Analysis

The LLM integration provides a qualitative fight analysis. Upon receiving two boxer IDs, the API builds a dynamic prompt with their stats and attributes and sends it to the AI, which returns a prediction of the likely winner.

**Example of Dynamic Prompt:**

```
Analyze a hypothetical boxing match between the following fighters and provide an opinion on who would likely win and why.

Fighter A:
- Name: Floyd Mayweather (Money)
- Record: 50 wins (27 KOs), 0 losses, 0 draws
- Weight Class: Welterweight

Fighter B:
- Name: Manny Pacquiao (PacMan)
- Record: 62 wins (39 KOs), 8 losses, 2 draws
- Weight Class: Welterweight

Consider their styles, physical attributes, and knockout power. Provide a prediction and likely outcome.
```

## ⚙️ How to Run the Project Locally

### Prerequisites

* Go 1.21+
* Docker & Docker Compose
* Azure account for Blob Storage
* API key from an LLM provider (e.g., OpenAI)

### Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/boxing-api.git
   cd boxing-api
   ```

2. **Set up environment variables:**
   Rename the `.env.example` file to `.env` and fill in your credentials:

   ```env
   # API Settings
   API_PORT=8080

   # Azure Blob Storage Credentials
   AZURE_STORAGE_ACCOUNT_NAME=your-account
   AZURE_STORAGE_ACCOUNT_KEY=your-key
   AZURE_STORAGE_CONTAINER_NAME=your-container

   # LLM API Key
   OPENAI_API_KEY=your-openai-key
   ```

3. **Run with Docker Compose:**
   This is the recommended method as it manages the application container and SQLite data volume.

   ```bash
   docker-compose up --build
   ```

4. **Access the API:**

   * API available at: `http://localhost:8080`
   * Swagger documentation: `http://localhost:8080/swagger/index.html`

## 🧪 Testing

To run the unit and integration test suite, use the following command:

```bash
go test ./...
```

## ☁️ Deployment

This application is optimized for low-cost deployment on platforms like **Railway** or **Render**. The strategy includes:

1. Connecting the Git repository to the platform.
2. Configuring the build to use the project’s `Dockerfile`.
3. Creating a **Volume** (or Persistent Disk) and mounting it to the directory where the `boxing.db` file is stored, ensuring data persistence.
4. Adding environment variables in the platform’s interface.

## ⏭️ Next Steps and Improvements

* [ ] Implement authentication and authorization with JWT.
* [ ] Add endpoints for managing fight events.
* [ ] Create a scraper (e.g., with `colly`) to populate the database with BoxRec information.
* [ ] Expand the data model to include more stats (e.g., punch accuracy).
* [ ] Set up a CI/CD pipeline with GitHub Actions to automate tests and deployments.

