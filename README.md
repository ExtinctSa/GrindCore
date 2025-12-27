\# Habit Tracker Backend

This repository contains the entire backend of a web application I created to track habits. It serves as a proof of concept for a potential mobile app I’d like to build in the future, given the resources to do so.

Note: This project is strictly backend logic. There is currently no frontend.

\---

\## Motivation

The main inspiration for this project came from my father, who sometimes forgets to eat. If this application were ever put into production, it could be used to send reminders and help with habit consistency. This project explores how habit data, completion tracking, and user authentication could work in a real-world system.

\---

\## Quick Start

1\. Clone the repository:

bash

git clone

cd habit-tracker-backend

Install dependencies (example):

go mod tidy

Configure environment variables (database connection, port, etc.).

Start the server:

go run main.go

The backend will run locally (default: http://localhost:9999).

Usage

Authentication & Middleware

Any handler wrapped with the authentication middleware requires an Authorization header.

Example Header:

Authorization: Bearer

Obtaining a User Token

You can obtain a token by calling the login handler with a valid username and password:

curl -X POST http://localhost:9999/api/login \\

\-H "Content-Type: application/json" \\

\-d '{"username": "", "password": ""}'

The token will be returned in the response:

{

"token": ""

}

API Endpoints

\-------------

All endpoints are prefixed with /api.Endpoints marked \*\*(Auth Required)\*\* must include a valid Authorization: Bearer  header.

\### User Endpoints

MethodEndpointDescriptionPOST/usersCreate a new user accountPOST/loginAuthenticate a user and return a JWTPUT/usersUpdate user information \*\*(Auth Required)\*\*GET/meRetrieve the currently authenticated user \*\*(Auth Required)\*\*

\### Token Validation & Refresh

MethodEndpointDescriptionPOST/refreshRefresh an expired access token using a refresh tokenPOST/revokeRevoke an existing refresh token

\### Habit Endpoints

MethodEndpointDescriptionPOST/habitsCreate a new habit \*\*(Auth Required)\*\*GET/habitsList all habits for the authenticated user \*\*(Auth Required)\*\*GET/habits?id=Retrieve a habit by ID \*\*(Auth Required)\*\*PATCH/habitsUpdate an existing habit \*\*(Auth Required)\*\*

\### Habit Completion Endpoints

MethodEndpointDescriptionPOST/habits/{id}/completionMark a habit as completed for the current date \*\*(Auth Required)\*\*GET/habits/{id}/completionCheck if a habit has been completed for the current date \*\*(Auth Required)\*\*

\### Notes on Completion Logic

\*   Each habit can only be completed \*\*once per day\*\*

\*   This is enforced by a database constraint on (habit\\\_id, completed\\\_date)

\*   Habits requiring multiple daily completions must be modeled as separate habits

Database Schema

The application uses a standard SQL database with the following structure.

Users Table

Column NameDescription

idRandomly generated UUID (using Google’s UUID package)

usernameAny non-null text

emailAny non-null text (not required to be a valid email)

hashed\\\_passwordHashed version of the user’s password (never returned by handlers)

created\\\_atTimestamp, defaults to current time

updated\\\_atTimestamp, defaults to current time and updates on user changes

Habits Table

Column NameDescription

idRandomly generated UUID

habit\\\_nameName of the habit (non-null text)

frequencyHow often the habit should be completed (currently unused)

categoryUsed for sorting habits

created\\\_atTimestamp, defaults to current time

updated\\\_atTimestamp, defaults to current time and updates on habit changes

user\\\_idUUID reference to the owning user (users.id)

Completion Table

Column NameDescription

idRandomly generated UUID

habit\\\_idUUID reference to a habit

user\\\_idUUID reference to a user

completed\\\_dateDate the habit was completed

completed\\\_atTimestamp, defaults to current time

Constraints:

The combination of habit\\\_id and completed\\\_date must be unique.

Completion data cannot be stored twice for the same habit on the same date.

Known Limitation:

Because of this constraint, habits that must be completed multiple times per day must be represented as separate habit entries (e.g., medication taken twice daily).

Contributing

This project currently has no frontend and focuses entirely on backend functionality.

If you spot an issue, have suggestions, or want to add features:

Create a new branch

Make your changes

Open a pull request

Contributions and feedback are welcome.
