# Database Information

To maintain a record of processed events, I decided to create a table in postgres. I run it locally using Docker:
```shell
$ docker run --name liveProject-postgres -e POSTGRES_PASSWORD=postgres -d postgres -p 5432:5432
```

Then, I created a database called `liveproject` and a schema within that database called `events` that contains the following table definition:
```sql
-- DROP TABLE events.processed_events;

CREATE TABLE events.processed_events (
	id uuid NOT NULL,
	processed_timestamp timestamp NOT NULL,
	event_name varchar(256) NOT NULL
);
```