CREATE TABLE IF NOT EXISTS apps
(
    id   int NOT NULL AUTO_INCREMENT,
    name varchar(80) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS rols
(
    id    int NOT NULL AUTO_INCREMENT,
    name  varchar(80) NOT NULL,
    state smallint(1) NOT NULL DEFAULT '1',
    del   tinyint(1)  NOT NULL DEFAULT '0',
    PRIMARY KEY (id),
    UNIQUE KEY rols_name_uindex (name),
    KEY rols_del_index (del),
    KEY rols_state_index (state)
);

CREATE TABLE IF NOT EXISTS rols_apps
(
    id     int(11) NOT NULL AUTO_INCREMENT,
    rol_id int(11) NOT NULL,
    app_id int(11) NOT NULL,
    PRIMARY KEY (id),
    KEY rols_apps_apps_id_fk (app_id),
    KEY rols_apps_rols_id_fk (rol_id),
    CONSTRAINT rols_apps_apps_id_fk FOREIGN KEY (app_id) REFERENCES apps (id),
    CONSTRAINT rols_apps_rols_id_fk FOREIGN KEY (rol_id) REFERENCES rols (id)
);

CREATE TABLE IF NOT EXISTS users
(
    id         int(11)     NOT NULL AUTO_INCREMENT,
    email      varchar(80) NOT NULL,
    password   varchar(80) NOT NULL,
    name       varchar(250)         DEFAULT NULL,
    state      smallint(1) NOT NULL DEFAULT '1',
    del        tinyint(1)  NOT NULL DEFAULT '0',
    created_at datetime             DEFAULT NULL,
    updated_at datetime             DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY users_email_uindex (email),
    KEY users_del_index (del),
    KEY users_email_password_index (email, password),
    KEY users_state_index (state)
);

CREATE TABLE IF NOT EXISTS users_rols
(
    id      int(11) NOT NULL AUTO_INCREMENT,
    user_id int(11) NOT NULL,
    rol_id  int(11) NOT NULL,
    PRIMARY KEY (id),
    KEY users_rols_rols_id_fk (rol_id),
    KEY users_rols_users_id_fk (user_id),
    CONSTRAINT users_rols_rols_id_fk FOREIGN KEY (rol_id) REFERENCES rols (id),
    CONSTRAINT users_rols_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
);

create view  v_users_apps  as
select ap.*,ur.user_id
from apps ap
inner join rols_apps ra on ap.id = ra.app_id
inner join users_rols ur on ra.rol_id = ur.rol_id;


-- DDL

INSERT INTO ucompensar.users (email, password, name, state, del, created_at, updated_at) VALUES ('admin@ucompensar.edu.co', '12345Admin#', 'Administrador del sistema', 1, 0, null, null);
INSERT INTO ucompensar.apps (name) VALUES ('Usuarios');
INSERT INTO ucompensar.apps (name) VALUES ('Roles');
INSERT INTO ucompensar.apps (name) VALUES ('Facturación');
INSERT INTO ucompensar.rols (name, state, del) VALUES ('Administrador', 1, 0);
INSERT INTO ucompensar.rols_apps (rol_id, app_id) VALUES (1, 1);
INSERT INTO ucompensar.rols_apps (rol_id, app_id) VALUES (1, 2);
INSERT INTO ucompensar.rols_apps (rol_id, app_id) VALUES (1, 3);
INSERT INTO ucompensar.users_rols (user_id, rol_id) VALUES (1, 1);