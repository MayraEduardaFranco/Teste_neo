-- Higienização de espaços para a tabela clientes
UPDATE clientes
SET 
    cpf = TRIM(cpf),  -- Remove espaços do CPF
    
    private = NULLIF(TRIM(private::text), '')::INT,  -- Remove espaços da coluna private e converte para INT
    
    incompleto = NULLIF(TRIM(incompleto::text), '')::INT,  -- Remove espaços da coluna incompleto e converte para INT
    
    data_ultima_compra = NULLIF(TRIM(data_ultima_compra::text), '')::DATE,  -- Remove espaços da coluna data_ultima_compra e converte para DATE
    
    ticket_medio = NULLIF(TRIM(ticket_medio::text), '')::NUMERIC(10, 2),  -- Remove espaços da coluna ticket_medio e converte para NUMERIC
    
    ticket_ultima_compra = NULLIF(TRIM(ticket_ultima_compra::text), '')::NUMERIC(10, 2),  -- Remove espaços da coluna ticket_ultima_compra e converte para NUMERIC
    
    loja_frequente = TRIM(loja_frequente),  -- Remove espaços da Loja Frequente
    
    loja_ultima_compra = TRIM(loja_ultima_compra);  -- Remove espaços da Loja da Última Compra
