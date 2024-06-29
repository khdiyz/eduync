-- +goose Up
CREATE TABLE "lessons" (
    "id" BIGSERIAL PRIMARY KEY,
    "theme" VARCHAR(255) NOT NULL,
    "group_id" BIGINT NOT NULL,
    "lesson_date" DATE NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
);

-- +goose Down
DROP TABLE "lessons";