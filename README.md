# PF2E Character API V0.5
A small api which stores pathfinder character details. Uses an insecure authorization token system (DO NOT EXPOSE PUBLICLY).
## Routes
v1/token: 
 - POST - Create users and makes a auth token which must be included in the 'Authorization' field of the header. Will eventually be replaced with JWTs.
 - GET - Gets the auth token for user

v1/character: GET, POST, PUT, and DELETE can be used as standard for CRUD apis

More details about parameters and responses can be found in 'api/api.go'
