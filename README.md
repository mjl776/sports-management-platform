
Overview

This project is a Golang project that creates a Sports Management Platform for managing Professional Sports Teams in terms of budgeting, player contracts, and employees. The idea is create a robust sports management backend application that allows sports professionals of different role types to access the platform and perform tasks on the platform based on their permissioning level.

Testing:

To test post routes

Test

Create Team Test Curl Request

curl localhost:8080/create-team --include --header "Content-Type: application/json" -d @test/teams_test_create.json --request "POST"

Create Teams Employee Curl Request

curl localhost:8080/create-team-employee --include --header "Content-Type: application/json" -d @test/new_employee_test_create.json --request "POST"

curl localhost:8080/create-user --include --header "Content-Type: application/json" -d @test/new_user_test_create.json --request "POST"

curl localhost:8080/create-player --include --header "Content-Type: application/json" -d @test/new_player_test_create.json --request "POST"

curl localhost:8080/create-player-contract --include --header "Content-Type: application/json" -d @test/new_player_contract_create.json --request "POST"

curl localhost:8080/login --include --header "Content-Type: application/json" -d @test/login_test.json --request "POST"