-- +goose Up
ALTER TABLE booked_tickets 
ALTER COLUMN concert_id SET NOT NULL,
ALTER COLUMN user_id SET NOT NULL;

-- +goose Down
ALTER TABLE booked_tickets 
ALTER COLUMN concert_id SET NULL,
ALTER COLUMN user_id SET  NULL;