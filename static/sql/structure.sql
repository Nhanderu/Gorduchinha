CREATE EXTENSION "uuid-ossp";

CREATE FUNCTION fn_trigger_updated_at() RETURNS TRIGGER AS $$
    BEGIN
        IF ROW(NEW.*) IS DISTINCT FROM ROW(OLD.*) THEN
            NEW.updated_at = NOW(); 
            RETURN NEW;
        ELSE
            RETURN OLD;
        END IF;
    END;
$$ LANGUAGE 'plpgsql';

-- tb_champ

CREATE TABLE tb_champ
    ( id         SERIAL
    , created_at TIMESTAMP NOT NULL DEFAULT NOW()
    , updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    , deleted_at TIMESTAMP NULL
    , CONSTRAINT pk_champ_id PRIMARY KEY (id)

    , slug VARCHAR(30)  NOT NULL
    , name VARCHAR(100) NOT NULL
    )
;

CREATE TRIGGER tr_champ_updated_at
    BEFORE UPDATE ON tb_champ
    FOR EACH ROW EXECUTE PROCEDURE fn_trigger_updated_at()
;

CREATE UNIQUE INDEX uq_champ_slug
    ON tb_champ (slug ASC)
    WHERE deleted_at IS NULL
; 

INSERT INTO tb_champ (slug, name)
    VALUES
          ( 'national-league-1-div',   'Campeaonato Brasileiro Série A' )
        , ( 'national-league-2-div',   'Campeaonato Brasileiro Série B' )
        , ( 'national-cup',            'Copa do Brasil' )
        , ( 'world-cup',               'Copa do Mundo de Clubes' )
        , ( 'intercontinental-cup',    'Copa Intercontinental' )
        , ( 'south-american-cup-a',    'Copa Libertadores da América' )
        , ( 'south-american-cup-b',    'Copa Sul-Americana' )
        , ( 'south-american-supercup', 'Recopa Sul-Americana' )
        , ( 'sp-state-cup',            'Campeonato Paulista' )
        , ( 'rj-state-cup',            'Campeonato Carioca' )
        , ( 'rs-state-cup',            'Campeonato Gaúcho' )
        , ( 'mg-state-cup',            'Campeonato Mineiro' )
;

-- tb_team

CREATE TABLE tb_team
    ( id         SERIAL
    , created_at TIMESTAMP NOT NULL DEFAULT NOW()
    , updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    , deleted_at TIMESTAMP NULL
    , CONSTRAINT pk_team_id PRIMARY KEY (id)

    , abbr      VARCHAR(10)  NOT NULL
    , name      VARCHAR(20)  NOT NULL
    , full_name VARCHAR(100) NOT NULL
    )
;

CREATE TRIGGER tr_team_updated_at
    BEFORE UPDATE ON tb_team
    FOR EACH ROW EXECUTE PROCEDURE fn_trigger_updated_at()
;

CREATE UNIQUE INDEX uq_team_abbr
    ON tb_team (abbr ASC)
    WHERE deleted_at IS NULL
;

INSERT INTO tb_team (abbr, name, full_name)
    VALUES
          ( 'sccp',  'Corinthians',   'Sport Club Corinthians Paulista' )
        , ( 'sep',   'Palmeiras',     'Sociedade Esportiva Palmeiras' )
        , ( 'spfc',  'São Paulo',     'São Paulo Futebol Clube' )
        , ( 'sfc',   'Santos',        'Santos Futebol Clube' )
        , ( 'crf',   'Flamengo',      'Clube de Regatas do Flamengo' )
        , ( 'crvg',  'Vasco',         'Clube de Regatas Vasco da Gama' )
        , ( 'ffc',   'Fluminense',    'Fluminense Football Club' )
        , ( 'bfr',   'Botafogo',      'Botafogo de Futebol e Regatas' )
        , ( 'cam',   'Atlético',      'Clube Atlético Mineiro' )
        , ( 'cec',   'Cruzeiro',      'Cruzeiro Esporte Clube' )
        , ( 'gfbpa', 'Grêmio',        'Grêmio Foot-Ball Porto Alegrense' )
        , ( 'iec',   'Internacional', 'Internacional Esporte Clube' )
;

-- tb_trophy

CREATE TABLE tb_trophy
    ( id         SERIAL
    , created_at TIMESTAMP NOT NULL DEFAULT NOW()
    , updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    , deleted_at TIMESTAMP NULL
    , CONSTRAINT pk_trophy_id PRIMARY KEY (id)

    , uuid CHAR(36) NOT NULL
    , year SMALLINT NOT NULL

    , team_id INT NOT NULL
    , CONSTRAINT fk_trophy_team_id FOREIGN KEY (team_id) REFERENCES tb_team (id)
    
    , champ_id INT NOT NULL
    , CONSTRAINT fk_trophy_champ_id FOREIGN KEY (champ_id) REFERENCES tb_champ (id)
    )
;

CREATE TRIGGER tr_trophy_updated_at
    BEFORE UPDATE ON tb_trophy
    FOR EACH ROW EXECUTE PROCEDURE fn_trigger_updated_at()
;

CREATE INDEX ix_trophy_team_id  ON tb_trophy (team_id ASC);
CREATE INDEX ix_trophy_champ_id ON tb_trophy (champ_id ASC);

CREATE UNIQUE INDEX uq_trophy_uuid
    ON tb_trophy (uuid ASC)
    WHERE deleted_at IS NULL
;
