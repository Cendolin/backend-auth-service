#!/bin/bash

export API_LISTEN_PORT=9898
export API_LISTEN_HOST=0.0.0.0
export JWT_KEY=NiGGACendol

export DATABASE_HOST=127.0.0.1
export DATABASE_PORT=5432
export DATABASE_USER=hanif
export DATABASE_PASSWORD=haniF123//
export DATABASE_DBNAME=cendolin

export RABBITMQ_URI=amqp://guest:guest@localhost:5672

go run .