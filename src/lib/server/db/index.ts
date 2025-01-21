import { Resource } from 'sst';
import { Pool } from 'pg';
import { drizzle } from 'drizzle-orm/node-postgres';
import * as schema from './schema';

const pool = new Pool({
  host: Resource.Db.host,
  port: Resource.Db.port,
  user: Resource.Db.username,
  password: Resource.Db.password,
  database: Resource.Db.database,
})

const db = drizzle(pool, { schema, casing: 'snake_case' });

export default db
export type Database = typeof db
