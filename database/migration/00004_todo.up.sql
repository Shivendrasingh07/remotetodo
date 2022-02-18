ALTER TABLE todos
    DROP COLUMN task_completed;

ALTER TABLE todos
    ADD COLUMN task_completed BOOLEAN DEFAULT FALSE;

ALTER TABLE todos
    ADD COLUMN archived_at timestamp DEFAULT null;
