# fsync
Remote file sync utility

```sql
CREATE TABLE DataSourceTypes(
    id integer primary key,
    Type varchar(500) unique
);

CREATE TABLE ConnectionInfo(
    id integer primary key,
    [Name] varchar(500) unique,
    UserName varchar(500),
    Password nvarchar(500),
    Passkey varchar(2000)
);

CREATE TABLE DestInfo(
    id integer primary key,
    [Name] varchar(500) unique,
    DestType varchar(500),
    DestConnStr varchar(5000),
    [Path] varchar(5000),
    ConnId integer,
    FOREIGN KEY(DestType) REFERENCES DataSourceTypes(Type),
    FOREIGN KEY(ConnId) REFERENCES ConnectionInfo(id)
);

CREATE TABLE SourceInfo(
    id integer primary key,
    [Name] varchar(500) unique,
    SourceType varchar(500),
    SourceConnStr varchar(5000),
    [Path] varchar(5000),
    ConnId integer,
    FOREIGN KEY(SourceType) REFERENCES DataSourceTypes(Type),
    FOREIGN KEY(ConnId) REFERENCES ConnectionInfo(id)
);

CREATE TABLE DataCopyTasks(
    id integer primary key,
    TaskName varchar(500) unique,
    Schedule varchar(200),
    SourceId integer,
    DestId integer,
    Details nvarchar(max),
    FOREIGN KEY(SourceId) REFERENCES SourceInfo(id),
    FOREIGN KEY(DestId) REFERENCES DestInfo(id)
)

CREATE TABLE TransformOps(
    id integer primary key,
    [Name] varchar(500) unique,
    SourceDelimiter varchar(10),
    DestDelimiter varchar(10)
)
```