#!/bin/bash
echo "******CREATING DOCKER DATABASE******"
psql --username postgres  <<- EOSQL
  CREATE DATABASE development;
  GRANT ALL PRIVILEGES ON DATABASE development to postgres;
EOSQL
psql --username postgres development <<- EOSQL
  CREATE TABLE users (id SERIAL PRIMARY KEY NOT NULL, github_id INT, login TEXT, email TEXT, avatar_url TEXT, vote1 INT, vote2 INT, vote3 INT, vote4 INT, vote5 INT);
EOSQL
echo "******DOCKER DATABASE CREATED******"
