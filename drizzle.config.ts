import dotenv from 'dotenv';
import dotenvExpand from 'dotenv-expand';
import { defineConfig } from 'drizzle-kit';
import { Resource } from 'sst';

dotenvExpand.expand(dotenv.config())

export default defineConfig({
  out: './drizzle',
  schema: './src/lib/server/db/schema.ts',
  dialect: 'postgresql',
  casing: 'snake_case',
  dbCredentials: {
    host: Resource.AutofyDb.host,
    port: Resource.AutofyDb.port,
    user: Resource.AutofyDb.username,
    password: Resource.AutofyDb.password,
    database: Resource.AutofyDb.database,
    ssl: false
  }
});

