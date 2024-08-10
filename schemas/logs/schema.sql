CREATE TABLE logs (
  "id" UUID NOT NULL DEFAULT gen_random_uuid(),
  endpoint TEXT NOT NULL,
  ip TEXT NULL,
  status_code integer NULL,
  request TEXT NOT NULL,
  response TEXT,
  create_at TIMESTAMP WITH TIME ZONE NULL
);
