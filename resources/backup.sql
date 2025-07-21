DROP TABLE IF EXISTS Atendimento;
DROP TABLE IF EXISTS Cliente;
DROP TABLE IF EXISTS Funcionario;
DROP TABLE IF EXISTS Servico;

-- CLIENTE
CREATE TABLE Cliente (
   id UUID PRIMARY KEY,
   nome VARCHAR(100) NOT NULL,
   cpf VARCHAR(14) UNIQUE NOT NULL,
   sexo CHAR(1) NOT NULL,
   idade INT NOT NULL,
   endereco VARCHAR(255) NOT NULL,
   conta VARCHAR(30) NOT NULL
);

-- FUNCIONÁRIO
CREATE TABLE Funcionario (
    id UUID PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    numero_identificacao VARCHAR(20) UNIQUE NOT NULL,
    salario FLOAT NOT NULL,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    sexo CHAR(1) NOT NULL,
    idade INT NOT NULL,
    horario_trabalho VARCHAR(50) NOT NULL
);

-- SERVIÇO
CREATE TABLE Servico (
     id SERIAL PRIMARY KEY,
     nome_servico VARCHAR(50) NOT NULL CHECK (nome_servico IN (
        'pagamento de boletos',
        'apostas em loterias',
        'saques',
        'depósitos'
    ))
);

-- ATENDIMENTO
CREATE TABLE Atendimento (
    id SERIAL PRIMARY KEY,
    numero_transacao UUID NOT NULL,
    data_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    cliente_id UUID REFERENCES Cliente(id),
    funcionario_id UUID REFERENCES Funcionario(id),
    servico_id SERIAL REFERENCES Servico(id),
    meio_pagamento VARCHAR(20) NOT NULL,
    CHECK (meio_pagamento IS NULL OR meio_pagamento IN (
        'dinheiro',
        'boleto',
        'cartão',
        'cheque',
        'PIX',
        'TED'
    ))
);