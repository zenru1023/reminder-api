CREATE TABLE IF NOT EXISTS reminders (
    id BIGINT UNSINGED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    remind_at DATETIME NOT NULL,
    status ENUM('pending', 'done', 'cancelled') NOT NULL DEFAULT 'pending',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_remind_at (remind_at),
    INDEX idx_status (status)
)