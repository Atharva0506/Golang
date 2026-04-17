CREATE TABLE IF NOT EXISTS notifications (
    id                   UUID         PRIMARY KEY,
    user_id              UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message_type         VARCHAR(100) NOT NULL,
    notification_status  VARCHAR(20)  NOT NULL DEFAULT 'PENDING'
                                      CHECK (notification_status IN ('PENDING', 'SENT', 'FAILED')),
    timestamp            TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_notifications_user_id        ON notifications (user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_status         ON notifications (notification_status);
CREATE INDEX IF NOT EXISTS idx_notifications_timestamp_desc ON notifications (timestamp DESC);
