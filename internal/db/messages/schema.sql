CREATE TABLE messages (
  message_id SERIAL PRIMARY KEY,
  topic text NOT NULL,
  message text NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
