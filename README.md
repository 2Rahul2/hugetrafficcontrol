## Using pooling and sharding to make writes opertaion of data base more quicker
the server allows more than 15k+ users to book concert tickets , using sharding and pooling it enables to allow multiple user to send request to the server and modify the contents in the data base

`/book-ticket-shards` this url uses the sharding method. This method takes approx 1min for booking 15K+users
`book-ticket-uni` and this is without sharding which heavily reduces the overall time to book as each user has to wait for the before user to book the ticket.This method approx takes 2+mins for booking 15K+ users
