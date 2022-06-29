CREATE TABLE Usuarios (
    ID        SERIAL PRIMARY KEY UNIQUE,
	Nome      VARCHAR,
	Sobrenome VARCHAR,
	Email     VARCHAR UNIQUE,
	CPFCNPJ   VARCHAR UNIQUE,
	Senha     VARCHAR,
	Saldo     FLOAT,
	Tipo      VARCHAR
);

CREATE TABLE Transaction (
    ID        SERIAL PRIMARY KEY UNIQUE,
	Valor     FLOAT,
	IDOrigem  INT,
	IDDestino INT,
	DateTime  TIMESTAMP
);

INSERT INTO Usuarios (Nome, Sobrenome, Email, CPFCNPJ, Senha, Saldo, Tipo) VALUES ('ANDERSON', 'SILVA', 'anderson@gmail.com', '123.123.123-90', '123456', 200, 'comum');
INSERT INTO Usuarios (Nome, Sobrenome, Email, CPFCNPJ, Senha, Saldo, Tipo) VALUES ('MARIA', 'SOUSA', 'maria@gmail.com',  '444.123.123-90', '444444', 800, 'comum');
INSERT INTO Usuarios (Nome, Sobrenome, Email, CPFCNPJ, Senha, Saldo, Tipo) VALUES ('Q2Bank', '', 'q2bank@q2bank.com', '47.263.110/0001-00',  '555555', 250.00, 'lojista');