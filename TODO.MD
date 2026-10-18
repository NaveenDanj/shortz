
- [X] Connection pool settings. Your ConnectPostgres sets no SetMaxOpenConns, SetMaxIdleConns or SetConnMaxLifetime on the underlying sql.DB. Without a cap, a traffic spike can open hundreds of connections and exhaust Postgres. This is the most likely first failure, so fix it now. Something like 20–25 open connections per instance is a reasonable start.

- [ ] Redis cache-aside on profile:{id} with a short TTL (a few minutes), invalidated on profile update. Most /me calls then never touch Postgres. Cache the "not found" result briefly too, so new users or bad ids don't hammer the DB.

- [ ]  Read replicas for profile reads, if one primary still struggles after the cache.

- [ ] Maybe add pgpooler (pbbouncer) for handle connection pooling