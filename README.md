This is the entire backend of a web application I created to track a habit. I created this as I think it's a nice proof of concept for a possible mobile app I would want to create
given the resources to do so. My main inspiration behind this was my father forgets to eat sometimes and I figured should this ever go into production it would be helpful for him.

For the database side it's a standard SQL server I have set up like so:
User Table:
ID - a randomly generated UUID using Google's UUID package

Username - allowed to be any text, the only constraint is that it cannot be null

Email - same as username, as it is not in production it's not required to be an actual email

Hashed_Password - as implied by the name it's the hashed version of whatever text was input as a password. This is never returned by any handler but is stored in the 

database and checked against for authentication

Created_at - a timestamp defaulting to the current time

Updated_at - a timestamp defaulting to the current time, updates when the user data is updated

Habits Table:
ID - a randomly generated UUID using Google's UUID package

HabitName - name of the habit, any text so long as it's not null

Frequency - how often the habit/task is to be performed. As of right now it's functionally useless but it would be needed to send reminders to users accordingly

Category - Used for sorting habits

Created_at - a timestamp defaulting to the current time

Updated_at - a timestamp defaulting to the current time, updates when the habit data is updated

User_id - a UUID reference to a user in order to show ownership of habits in the database. References the ID section of the user table

Completion Table:
ID - a randomly generated UUID using Google's UUID package, this exists here just to make sure completion data is under a unique ID for storage and searching purposes

Habit_id - a UUID reference to a habit, used to mark the correct habit as completed

User_id - a UUID reference to a user in order to show ownership of habit completion in the database

Completed_date - Date the habit was completed

Completed_at - Timestamp with the default of the current time

Habit_id and Completed_date must be unique from all other entries. Completion data cannot be stored twice in the database.
The unfortunate side effect of this is that for tasks that are to be completed multiple times a day must be individual entries in the habit table.
Essentially, you'd have to create multiple habits for different times a day if say you took medication more than once a day.

Handler w/ middleware usage:
To use any handler wrapped in the Authentication Middleware you must have an Authorization header.
For example:
-H "Authorization: Bearer <user-token>"
The user token can be aquired by running the login handler and using a valid username and password:
curl -X POST http://localhost:9999/api/login -H "Content-Type: application/json" -d '{"username": "<chosen-username>", "password": "<password>"}'
it will be listed on the stdout return under "token: "

I unfortunately don't have a front end for this app and this is strictly the backend logic. If there's a problem you spot with the code or something you wish to add
please create a branch and a pull request.
