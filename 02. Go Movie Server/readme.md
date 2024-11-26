### A simple Api server in Gorilla Mux

### Routes
> The following routes are available in the api
- `GET` /movies 
  > Returns all movies in the database
- `GET` /movies/{id}
  > Returns a single movie with the given id
- `POST` /movies
  > Adds a new movie to the database
- `PUT` /movies/{id}
  > Updates a movie with the given id
- `DELETE` /movies/{id}
  > Deletes a movie with the given id