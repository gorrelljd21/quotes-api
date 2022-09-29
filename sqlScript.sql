drop table if exists quotes;

create table quotes (
    id varchar primary key unique,
    phrase varchar,
    author varchar
);

insert into quotes(id, phrase, author)
values ('0d949b68-6b04-4b35-82e5-63159b7608f8', 'Dont communicate by sharing memory, share memory by communicating.', 'Rob Pike'),
('f9a7ab3a-9fc5-40b3-8c2e-76239ca037ce', 'Concurrency is not parallelism.', 'Rob Pike'),
('1f9c3697-5232-45a8-82b7-ba9ac5f0799c', 'Channels orchestrate; mutexes serialize.', 'Rob Pike'),
('a240c7e9-1570-4c36-ae5f-699e4cb5e4d7', 'The bigger the interface, the weaker the abstraction.', 'Rob Pike'),
('a2523b46-42d4-42f6-aeb9-42da4b928c4a', 'Use consistent spelling of certain words.', 'Dmitri Shuralyov'),
('f5a05e7f-1e71-462f-8036-9b7c8bfbed65', 'Single spaces between spaces.', 'Dmitri Shuralyov'),
('7dbde6f1-c411-40ca-af84-cc7fec7c06ec', 'Avoid unused method receiver names.', 'Dmitri Shuralyov'),
('170f9d56-369e-4088-a23d-5c8bc3e4a973', 'Comments for humans always have a single space after the slashes.', 'Dmitri Shuralyov');
