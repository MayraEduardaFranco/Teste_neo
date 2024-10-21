CREATE OR REPLACE FUNCTION validar_cnpj(cnpj TEXT) RETURNS BOOLEAN AS $$
DECLARE
    soma INTEGER;
    resto INTEGER;
    digito1 INTEGER;
    digito2 INTEGER;
BEGIN
    -- Remover caracteres não numéricos
    cnpj := regexp_replace(cnpj, '[^0-9]', '', 'g');

    -- Verificar se o CNPJ tem 14 dígitos
    IF LENGTH(cnpj) <> 14 THEN
        RETURN FALSE;
    END IF;

    -- Verificar se todos os dígitos são iguais
    IF cnpj = '00000000000000' OR cnpj = '11111111111111' OR cnpj = '22222222222222' OR
       cnpj = '33333333333333' OR cnpj = '44444444444444' OR cnpj = '55555555555555' OR
       cnpj = '66666666666666' OR cnpj = '77777777777777' OR cnpj = '88888888888888' OR
       cnpj = '99999999999999' THEN
        RETURN FALSE;
    END IF;

    -- Cálculo do primeiro dígito verificador
    soma := 0;
    FOR i IN 1..12 LOOP
        soma := soma + (CAST(SUBSTRING(cnpj, i, 1) AS INTEGER) * CASE
            WHEN i <= 4 THEN 6 - i  -- Para os 5 primeiros dígitos
            ELSE 14 - i             -- Para os 8 últimos dígitos
        END);
    END LOOP;

    resto := soma % 11;
    digito1 := CASE
                   WHEN resto < 2 THEN 0
                   ELSE 11 - resto
               END;

    IF digito1 <> CAST(SUBSTRING(cnpj, 13, 1) AS INTEGER) THEN
        RETURN FALSE;
    END IF;

    -- Cálculo do segundo dígito verificador
    soma := 0;
    FOR i IN 1..13 LOOP
        soma := soma + (CAST(SUBSTRING(cnpj, i, 1) AS INTEGER) * CASE
            WHEN i <= 5 THEN 7 - i  -- Para os 5 primeiros dígitos
            ELSE 15 - i             -- Para os 8 últimos dígitos
        END);
    END LOOP;

    resto := soma % 11;
    digito2 := CASE
                   WHEN resto < 2 THEN 0
                   ELSE 11 - resto
               END;

    IF digito2 <> CAST(SUBSTRING(cnpj, 14, 1) AS INTEGER) THEN
        RETURN FALSE;
    END IF;

    RETURN TRUE;  -- Retorna TRUE se o CNPJ for válido
END;
$$ LANGUAGE plpgsql;
