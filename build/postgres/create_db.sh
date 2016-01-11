#!/bin/bash
echo "******CREATING DOCKER DATABASE******"
psql --username postgres  <<- EOSQL
  CREATE DATABASE development;
  GRANT ALL PRIVILEGES ON DATABASE development to postgres;
EOSQL
psql --username postgres development <<- EOSQL
  CREATE TABLE users (id INT PRIMARY KEY NOT NULL, github_id INT, github_user_name TEXT, email TEXT);
  CREATE TABLE votes (id INT PRIMARY KEY NOT NULL, user_id INT, repo_id INT, weight INT);
EOSQL
echo "******DOCKER DATABASE CREATED******"
