CREATE TABLE IF NOT EXISTS User (
    ID        INTEGER,
	Nome      VARCHAR(30),
	Sobrenome VARCHAR(40),
	Email     VARCHAR(255),
	CPFCNPJ   VARCHAR(30),
	Senha     VARCHAR(255),
	Saldo     NUMERIC,
	Tipo      VARCHAR(30),
    PRIMARY KEY (ID)
)

CREATE TABLE IF NOT EXISTS Transaction (
    ID        INTEGER,
	Value     NUMERIC,
	IDOrigin  INTEGER,
	IDDestiny INTEGER,
	DateTime  TIME
)