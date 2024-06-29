-- +goose Up
CREATE TABLE "attendances" (
    "id" BIGSERIAL PRIMARY KEY,
    "lesson_id" BIGINT NOT NULL,
    "student_id" BIGINT NOT NULL,
    "has_attended" BOOLEAN NOT NULL,
    "is_reasonable" BOOLEAN,
    "reason" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("lesson_id") REFERENCES "lessons"("id"),
    FOREIGN KEY ("student_id") REFERENCES "students"("id")
);

-- +goose Down
DROP TABLE "attendances";