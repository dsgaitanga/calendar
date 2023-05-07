CREATE TABLE event (id bigint AUTO_INCREMENT NOT NULL, description varchar(255), link varchar(255), start_time datetime NOT NULL, end_time datetime NOT NULL, status BOOLEAN NOT NULL, PRIMARY KEY (id));
CREATE TABLE user (id bigint AUTO_INCREMENT NOT NULL, email varchar(50) NOT NULL, PRIMARY KEY (id));
CREATE TABLE tutor (id bigint AUTO_INCREMENT NOT NULL, email varchar(50) NOT NULL, PRIMARY KEY (id));
CREATE INDEX user_id ON user (id);
CREATE INDEX tutor_id ON tutor (id);
CREATE INDEX event_id ON event (id);
ALTER TABLE event ADD CONSTRAINT asiste_a FOREIGN KEY (id_user) REFERENCES user (id);
ALTER TABLE event ADD CONSTRAINT tutoria FOREIGN KEY (id_tutoria) REFERENCES tutor (id);