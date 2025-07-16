-- CLIENTE
CREATE TABLE Cliente (
   id uuid PRIMARY KEY,
   nome VARCHAR(100) NOT NULL,
   numero_identificacao VARCHAR(20) UNIQUE NOT NULL,
   cpf VARCHAR(14) UNIQUE NOT NULL,
   sexo CHAR(1),
   idade INT,
   endereco VARCHAR(255),
   conta VARCHAR(30)
);

-- FUNCIONÁRIO
CREATE TABLE Funcionario (
    id uuid PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    numero_identificacao VARCHAR(20) UNIQUE NOT NULL,
    salario DECIMAL NOT NULL,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    sexo CHAR(1),
    idade INT,
    horario_trabalho VARCHAR(50)
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
    numero_transacao uuid,
    data_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    cliente_id INT REFERENCES Cliente(id),
    funcionario_id INT REFERENCES Funcionario(id),
    servico_id INT REFERENCES Servico(id),
    meio_pagamento VARCHAR(20),
    CHECK (meio_pagamento IS NULL OR meio_pagamento IN (
       'dinheiro',
       'boleto',
        'cartão',
        'cheque',
        'PIX',
        'TED'
    ))
);
