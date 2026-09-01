# ADR 0002 — Code Structure of Go

Date: 2026-08-18

One thing I have noticed about Go from the main.go entry point, is the way errors are handled in conjuction with how the server is started.
The errors are handled in a "non-conventional way in relation to other languages i.e. python or java exceptions.
