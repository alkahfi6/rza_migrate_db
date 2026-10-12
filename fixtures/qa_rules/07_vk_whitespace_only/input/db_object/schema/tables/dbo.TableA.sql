CREATE TABLE dbo.TableA (
    id INT NOT NULL,
    status VARCHAR(10) NULL,
    amount DECIMAL(10,2) NULL,
    CONSTRAINT PK_TableA PRIMARY KEY (id),
    CONSTRAINT DF_TableA__status DEFAULT ('NEW') FOR status
);