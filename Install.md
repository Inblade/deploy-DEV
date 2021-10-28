1. DB - Postgres

2. Dumping and restoring on the same server:

$ pg_dump -h localhost -Fc test > /home/postgres/dump.sql
$ pg_restore -h localhost test < /home/postgres/dump.sql

3. Dumping and restoring on multiple hosts:

$ # Dump a remote database to your local machine
$ pg_dump -h remotedb.mydomain.com -f /home/postgres/dump.sql test
$ # dump a local database and write to a remote machine
$ pg_dump -h remotedb.mydomain.com test | ssh postgres@remotedb.mydomain.com 'cat > dump.sql'
$ # dump a remote database and write to the same remote machine
$ pg_dump -h remotedb.mydomain.com test | ssh postgres@remotedb.mydomain.com 'cat > dump.sql'
$ # or a different remote machine
$ pg_dump -h remotedb1.mydomain.com test | ssh postgres@remotedb2.mydomain.com 'cat > dump.sql'
4. 
