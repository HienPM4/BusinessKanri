import { spawn } from "node:child_process";

const argv = process.argv.slice(2);
const normalizedArgs = [];

for (let i = 0; i < argv.length; i += 1) {
  const arg = argv[i];

  if (arg === "--host") {
    normalizedArgs.push("--hostname");

    const nextArg = argv[i + 1];
    if (nextArg && !nextArg.startsWith("-")) {
      normalizedArgs.push(nextArg);
      i += 1;
    }

    continue;
  }

  normalizedArgs.push(arg);
}

const child = spawn(
  process.platform === "win32" ? "npx.cmd" : "npx",
  ["next", "dev", ...normalizedArgs],
  { stdio: "inherit" }
);

child.on("exit", (code, signal) => {
  if (signal) {
    process.kill(process.pid, signal);
    return;
  }

  process.exit(code ?? 1);
});
