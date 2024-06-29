-- +goose Up
CREATE TABLE "groups" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "course_id" BIGINT NOT NULL,
    "teacher_id" BIGINT NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("course_id") REFERENCES "courses"("id"),
    FOREIGN KEY ("teacher_id") REFERENCES "users"("id")
);

-- +goose Down
DROP TABLE "groups";