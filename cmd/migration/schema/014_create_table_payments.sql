-- +goose Up
CREATE TABLE "payments" (
    "id" BIGSERIAL PRIMARY KEY,
    "student_id" BIGINT NOT NULL,
    "group_id" BIGINT NOT NULL,
    "paid_amount" INTEGER NOT NULL,
    "payment_date" DATE NOT NULL,
    "returned_amount" INTEGER NOT NULL DEFAULT 0,
    "returned_date" DATE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY ("student_id") REFERENCES "students"("id"),
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
);

-- +goose Down
DROP TABLE "payments";