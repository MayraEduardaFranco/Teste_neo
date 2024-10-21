CREATE OR REPLACE FUNCTION validar_cpf(cpf TEXT) RETURNS BOOLEAN AS $$
DECLARE
    soma INTEGER;
    resto INTEGER;
    digito1 INTEGER;
    digito2 INTEGER;
BEGIN
    -- Remover caracteres não numéricos
    cpf := regexp_replace(cpf, '[^0-9]', '', 'g');

    -- Verificar se o CPF tem 11 dígitos
    IF LENGTH(cpf) <> 11 THEN
        RETURN FALSE;
    END IF;

    -- Verificar se todos os dígitos são iguais
    IF cpf = '00000000000' OR cpf = '11111111111' OR cpf = '22222222222' OR
       cpf = '33333333333' OR cpf = '44444444444' OR cpf = '55555555555' OR
       cpf = '66666666666' OR cpf = '77777777777' OR cpf = '88888888888' OR
       cpf = '99999999999' THEN
        RETURN FALSE;
    END IF;

    -- Cálculo do primeiro dígito verificador
    soma := 0;
    FOR i IN 1..9 LOOP
        soma := soma + (CAST(SUBSTRING(cpf, i, 1) AS INTEGER) * (11 - i));
    END LOOP;

    resto := soma % 11;
    digito1 := CASE
                   WHEN resto < 2 THEN 0
                   ELSE 11 - resto
               END;

    IF digito1 <> CAST(SUBSTRING(cpf, 10, 1) AS INTEGER) THEN
        RETURN FALSE;
    END IF;

    -- Cálculo do segundo dígito verificador
    soma := 0;
    FOR i IN 1..10 LOOP
        soma := soma + (CAST(SUBSTRING(cpf, i, 1) AS INTEGER) * (12 - i));
    END LOOP;

    resto := soma % 11;
    digito2 := CASE
                   WHEN resto < 2 THEN 0
                   ELSE 11 - resto
               END;

    IF digito2 <> CAST(SUBSTRING(cpf, 11, 1) AS INTEGER) THEN
        RETURN FALSE;
    END IF;

    RETURN TRUE;
END;
$$ LANGUAGE plpgsql;
