-- Creación de la base de datos
DROP DATABASE IF EXISTS investment;
CREATE DATABASE investment;

-- usar
USE investment;

CREATE TABLE assigner(
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    investment INT NOT NULL,
    credit_type_300 INT DEFAULT 0,
    credit_type_500 INT DEFAULT 0,
    credit_type_700 INT DEFAULT 0,
    processed BOOLEAN
);

SELECT BIN_TO_UUID(a.id) id, a.investment, a.credit_type_300, a.credit_type_500, a.credit_type_700, a.processed
FROM assigner AS a
LIMIT 100;

DELETE FROM assigner as a WHERE a.id in (UUID_TO_BIN('068b381c-e436-11ee-8a3b-75f156da42bc'));