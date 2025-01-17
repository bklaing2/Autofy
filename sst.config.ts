/// <reference path="./.sst/platform/config.d.ts" />


const DB_CONFIG = {
  username: "postgres",
  password: "password",
  database: "local",
  host: "localhost",
  port: 5432,
}


export default $config({

  app(input) {
    return {
      name: "autofy",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
    };
  },

  async run() {
    const vpc = new sst.aws.Vpc("AutofyVpc", { bastion: true, nat: "ec2" });
    const rds = new sst.aws.Postgres("AutofyDb", {
      vpc,
      proxy: true,
      dev: DB_CONFIG
    });

    new sst.x.DevCommand("AutofyDbLocal", {
      dev: {
        command: `docker run --rm -p ${DB_CONFIG.port}:5432 -v autofy-db:/var/lib/postgresql/data -e POSTGRES_USER=${DB_CONFIG.username} -e POSTGRES_PASSWORD=${DB_CONFIG.password} -e POSTGRES_DB=${DB_CONFIG.database} postgres:17.2-alpine`,
      },
    });

    new sst.x.DevCommand("DrizzleStudio", {
      link: [rds],
      dev: {
        command: "bunx drizzle-kit studio",
      },
    });

    const cluster = new sst.aws.Cluster("AutofyCluster", { vpc });

    // TODO: make dev use the container image
    cluster.addService("AutofyFrontend", {
      link: [rds],
      loadBalancer: {
        ports: [{ listen: "80/http", forward: "3000/http" }],
      },
      dev: {
        command: "bun run dev",
      },
    });
  },
});
