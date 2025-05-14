# Stock Recommendation App
## Summary
This is a Web application developed to recommend the best stock investment from the Data collected from the database. The frontend was developed using Vite+Vue, using Pinia Stores and Tailwind for Styles.
The API is made using Golang and it counts with some Test functions using Testify library, the Database is powered by CockroachDB a cloud platform.



The Application is a simple interface with a Table that displays the data obtained from the API and it also shows the most recommended Companies to invest below. The API calculates a Score for every company depending on the information it has, so only companies with a scoreof 0.75 or better will show as recommended

## Instalation

In order to run the project after cloning the Repository run the command "go install" in the terminal to install all the dependencies needed for the API, and then you can run it with the command "go run main.go"

NOTE: Every API KEY, Authorization Bearer and Database Password is not included in the code.

For the Front-end (it is inside the vite-vue folder) all you need to do is install all the dependencies with (npm install) and then you can run the project with (npm run dev)
