Habit Tracker Backend
This repository contains the entire backend of a web application I created to track habits. It serves as a proof of concept for a potential mobile app I’d like to build in the future, given the resources to do so.
The main inspiration for this project came from my father, who sometimes forgets to eat. If this application were ever put into production, it could be used to send reminders and help with habit consistency.
Note: This project is strictly backend logic. There is currently no frontend.
Database Schema
The application uses a standard SQL database with the following structure.
Users Table
Column Name	Description
id	Randomly generated UUID (using Google’s UUID package)
username	Any non-null text
email	Any non-null text (not required to be a valid email)
hashed_password	Hashed version of the user’s password (never returned by handlers)
created_at	Timestamp, defaults to current time
updated_at	Timestamp, defaults to current time and updates on user changes
Habits Table
Column Name	Description
id	Randomly generated UUID
habit_name	Name of the habit (non-null text)
frequency	How often the habit should be completed (currently unused, intended for reminders)
category	Used for sorting habits
created_at	Timestamp, defaults to current time
updated_at	Timestamp, defaults to current time and updates on habit changes
user_id	UUID reference to the owning user (users.id)
Completion Table
Column Name	Description
id	Randomly generated UUID (used for unique identification and searching)
habit_id	UUID reference to a habit
user_id	UUID reference to a user
completed_date	Date the habit was completed
completed_at	Timestamp, defaults to current time
Constraints:
The combination of habit_id and completed_date must be unique.
Completion data cannot be stored twice for the same habit on the same date.
Known Limitation:
Because of this constraint, habits that need to be completed multiple times per day must be represented as separate habit entries. For example, taking medication twice daily would require two separate habits.
Authentication & Middleware
Any handler wrapped with the authentication middleware requires an Authorization header.
Example Header
-H "Authorization: Bearer <user-token>"
Obtaining a User Token
You can obtain a token by calling the login handler with a valid username and password:
curl -X POST http://localhost:9999/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "<chosen-username>", "password": "<password>"}'
The token will be returned in the response under:
token:
Contributing
This project currently has no frontend and focuses entirely on backend functionality.
If you spot an issue, have suggestions, or want to add features:
Create a new branch
Open a pull request
Contributions and feedback are welcome.
