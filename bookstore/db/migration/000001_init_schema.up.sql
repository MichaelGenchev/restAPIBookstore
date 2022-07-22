CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "books" int[]
);

CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "author" bigint,
  "category" bigint,
  "price" float
);

CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "biography" varchar,
  "books" int[]
);

CREATE INDEX ON "categories" ("id");

CREATE INDEX ON "books" ("id");

CREATE INDEX ON "authors" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("id") REFERENCES "categories" ("books");

ALTER TABLE "books" ADD FOREIGN KEY ("author") REFERENCES "authors" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");
