-- name: Createuser :one
insert into users(id ,name ,created_at)
values($1 ,$2 ,$3)
RETURNING *;

-- name: GetUser :one
select * from users where id=$1;