create table leads (
    id         INT NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),

    webhook_id VARCHAR(256) NOT NULL,
	name       VARCHAR(256) NOT NULL,
	email      VARCHAR(256) NOT NULL,
	phone      VARCHAR(256) NOT NULL,
    others     JSON,
    PRIMARY KEY ( id ),
    INDEX webhook_id ( webhook_id )
);
