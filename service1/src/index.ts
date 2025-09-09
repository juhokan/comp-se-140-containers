import { Hono } from "hono";
import { env } from "bun";
import { appendFile } from "fs/promises";

const app = new Hono();

const STORAGE_URL = env.STORAGE_URL || "http://storage:5000";
const SERVICE2_URL = env.GO_SERVER_URL || "http://go-server:8080";
const VSTORAGE_PATH = "/vstorage/log.txt";

async function getUptimeHours() {
  const uptimeSeconds = Bun.nanoseconds() / 1e9;
  return (uptimeSeconds / 3600).toFixed(2);
}

async function getFreeDiskMB() {
  // Bun does not provide disk info, so we mock this for the sake of the example
  return 9999;
}

function getTimestamp() {
  // Since bun does not support toISOString without milliseconds, we remove them manually using regex
  return new Date().toISOString().replace(/\.\d{3}Z$/, "Z");
}

async function createStatusRecord() {
  const timestamp = getTimestamp();
  const uptime = await getUptimeHours();
  const freeDisk = await getFreeDiskMB();
  return `${timestamp}: uptime ${uptime} hours, free disk in root: ${freeDisk} MBytes`;
}

app.get("/", (c) => c.text("Service1 is up and running!"));

app.get("/status", async (c) => {
  const record1 = await createStatusRecord();

  await fetch(`${STORAGE_URL}/log`, {
    method: "POST",
    body: record1,
    headers: { "Content-Type": "text/plain" },
  });

  await appendFile(VSTORAGE_PATH, record1 + "\n");
  const res = await fetch(`${SERVICE2_URL}/status`);
  const record2 = await res.text();

  return c.text(`${record1}\n${record2}`);
});

app.get("/log", async (c) => {
  const res = await fetch(`${STORAGE_URL}/log`);
  const log = await res.text();

  return c.text(log);
});

export default app;
