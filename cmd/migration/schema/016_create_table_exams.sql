-- +goose Up
CREATE TABLE "exams" (
    "id" BIGSERIAL PRIMARY KEY,
    "group_id" BIGINT NOT NULL,
    "exam_date" DATE NOT NULL,
    "exam_start_time" TIME NOT NULL,
    "exam_end_time" TIME NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
);

-- +goose Down
DROP TABLE "exams";