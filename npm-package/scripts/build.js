#!/usr/bin/env node
/**
 * build.js — 将 cli/ur/dist/skill/ 的文档复制到 npm-package/ 目录
 *
 * 用法:
 *   node scripts/build.js
 */

const fs = require("fs");
const path = require("path");

const DIST_SKILL_DIR = path.resolve(__dirname, "../../dist/skill");
const NPM_PACKAGE_DIR = path.resolve(__dirname, "..");

if (!fs.existsSync(DIST_SKILL_DIR)) {
  console.error(`[build] error: dist skill dir not found: ${DIST_SKILL_DIR}`);
  console.error("[build] please run: bash scripts/package-skill.sh");
  process.exit(1);
}

// 清理旧文件（保留 package.json, index.js, index.d.ts, scripts/）
const keep = new Set(["package.json", "index.js", "index.d.ts", "scripts"]);
for (const entry of fs.readdirSync(NPM_PACKAGE_DIR)) {
  if (keep.has(entry)) continue;
  const fullPath = path.join(NPM_PACKAGE_DIR, entry);
  fs.rmSync(fullPath, { recursive: true, force: true });
}

// 复制 skill 文档
function copyDir(src, dst) {
  fs.mkdirSync(dst, { recursive: true });
  for (const entry of fs.readdirSync(src, { withFileTypes: true })) {
    const srcPath = path.join(src, entry.name);
    const dstPath = path.join(dst, entry.name);
    if (entry.isDirectory()) {
      copyDir(srcPath, dstPath);
    } else {
      fs.copyFileSync(srcPath, dstPath);
    }
  }
}

copyDir(DIST_SKILL_DIR, NPM_PACKAGE_DIR);

console.log(`[build] copied skill docs from ${DIST_SKILL_DIR} -> ${NPM_PACKAGE_DIR}`);
console.log(`[build] files: ${fs.readdirSync(NPM_PACKAGE_DIR).join(", ")}`);
