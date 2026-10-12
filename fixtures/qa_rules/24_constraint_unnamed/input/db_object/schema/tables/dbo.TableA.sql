CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) CONSTRAINT DF_TableA__status DEFAULT ('NEW') NULL,
  amount DECIMAL(10,2) NULL,
  PRIMARY KEY (id)
);
