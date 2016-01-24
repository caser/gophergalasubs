#!/bin/bash
echo "******CREATING DOCKER DATABASE******"
psql --username postgres  <<- EOSQL
  CREATE DATABASE development;
  GRANT ALL PRIVILEGES ON DATABASE development to postgres;
EOSQL
echo "******DOCKER DATABASE CREATED******"
