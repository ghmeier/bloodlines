DROP TABLE IF EXISTS b_trigger;
CREATE TABLE b_trigger(
	id VARCHAR(36) NOT NULL,
	tkey VARCHAR(1024) NOT NULL,
	vals VARCHAR(4096) NOT NULL,
	contentId  VARCHAR(36) NOT NULL, FOREIGN KEY fk_triggercontent(contentId) REFERENCES content(id),
	PRIMARY KEY (id)
);