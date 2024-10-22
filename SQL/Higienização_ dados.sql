-- Higienização de espaços para a tabela clientes
UPDATE clientes
SET 
    cpf = TRIM(cpf), 
    
    private = NULLIF(TRIM(private::text), '')::INT,  
    
    incompleto = NULLIF(TRIM(incompleto::text), '')::INT,  T
    
    data_ultima_compra = NULLIF(TRIM(data_ultima_compra::text), '')::DATE,     
    ticket_medio = NULLIF(TRIM(ticket_medio::text), '')::NUMERIC(10, 2),     
    ticket_ultima_compra = NULLIF(TRIM(ticket_ultima_compra::text), '')::NUMERIC(10, 2),     
    loja_frequente = TRIM(loja_frequente),     
    loja_ultima_compra = TRIM(loja_ultima_compra); 