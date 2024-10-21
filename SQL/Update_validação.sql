ALTER TABLE clientes ADD COLUMN cnpj_valido_ultimacompra BOOLEAN;
ALTER TABLE clientes ADD COLUMN cpf_valido BOOLEAN;
ALTER TABLE clientes ADD COLUMN cnpj_valido_lojafrequente BOOLEAN;
UPDATE clientes
SET cpf_valido = validar_cpf(cpf);
UPDATE clientes
SET cnpj_valido_ultimacompra = validar_cnpj(loja_ultima_compra);
UPDATE clientes
set cnpj_valido_lojafrequente = validar_cnpj (loja_frequente)