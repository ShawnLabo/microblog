CREATE TABLE Messages (
  MessageId     STRING(36)  NOT NULL,
  CreatedAt     TIMESTAMP   NOT NULL,
  Name          STRING(MAX) NOT NULL,
  Body          STRING(MAX) NOT NULL,
  WrittenRegion STRING(MAX) NOT NULL,
) PRIMARY KEY (MessageId, CreatedAt DESC);
