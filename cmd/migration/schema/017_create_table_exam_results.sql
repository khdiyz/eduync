-- +goose Up
CREATE TABLE "exam_results" (
    "id" BIGSERIAL PRIMARY KEY,
    "exam_id" BIGINT NOT NULL,
    "student_id" BIGINT NOT NULL,
    "result_ball" FLOAT NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("exam_id") REFERENCES "exams"("id"),
    FOREIGN KEY ("student_id") REFERENCES "students"("id")
);

-- +goose Down
DROP TABLE "exam_results";