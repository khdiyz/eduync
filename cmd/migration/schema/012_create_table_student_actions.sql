-- +goose Up
CREATE TABLE "student_actions" (
    "id" BIGSERIAL PRIMARY KEY,
    "student_id" BIGINT NOT NULL,
    "action_name" VARCHAR(64) NOT NULL,
    "action_text" TEXT,
    "action_date" DATE NOT NULL,
    "action_time" TIME NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("student_id") REFERENCES "students"("id")
);

-- +goose Down
DROP TABLE "student_actions";