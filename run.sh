#!/bin/zsh

go build -o bookings cmd/web/*.go
./bookings -dbname=bookings -dbuser=nepo -dbpass=Waffentrager-720 -dbhost=localhost -dbport=5432 -dbssl=disable