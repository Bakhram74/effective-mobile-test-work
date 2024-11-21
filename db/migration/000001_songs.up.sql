
CREATE TABLE "songs" (
  "id" uuid DEFAULT gen_random_uuid(),
  "group" varchar NOT NULL,
  "name" varchar NOT NULL,
  "releaseDate" varchar NOT NULL,
  "text" varchar NOT NULL,
  "link" varchar NOT NULL, 
  PRIMARY KEY (id)
);