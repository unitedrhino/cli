/**
 * @unitedrhino/ur-api-skills
 * 联犀 SaaS 平台 ur-api 技能包 — JS 工具入口
 *
 * 提供 skill 文档的路径解析和内容读取能力，方便外部 AI 工具集成。
 */

const fs = require("fs");
const path = require("path");

const PACKAGE_ROOT = __dirname;

/**
 * 支持的 app 列表
 * @type {string[]}
 */
const APPS = [
  "ur-platform-manage",
  "ur-iot",
  "ur-org-manage",
  "ur-org-energy",
  "ur-console",
];

/**
 * 获取 skill 文档的绝对路径
 * @param {string} name - skill 名称，如 "ur-iot" 或 "ur-api"（顶层索引）
 * @returns {string|null}
 */
function resolveSkillPath(name) {
  if (name === "ur-api") {
    const p = path.join(PACKAGE_ROOT, "SKILL.md");
    return fs.existsSync(p) ? p : null;
  }
  if (APPS.includes(name)) {
    const p = path.join(PACKAGE_ROOT, name, "SKILL.md");
    return fs.existsSync(p) ? p : null;
  }
  return null;
}

/**
 * 读取 skill Markdown 内容
 * @param {string} name - skill 名称
 * @returns {string|null}
 */
function getSkillMarkdown(name) {
  const p = resolveSkillPath(name);
  if (!p) return null;
  return fs.readFileSync(p, "utf-8");
}

/**
 * 获取所有可用 skill 的元数据列表
 * @returns {{name: string, path: string, description: string}[]}
 */
function listSkills() {
  const result = [];
  const topPath = path.join(PACKAGE_ROOT, "SKILL.md");
  if (fs.existsSync(topPath)) {
    const content = fs.readFileSync(topPath, "utf-8");
    const descMatch = content.match(/description:\s*"([^"]+)"/);
    result.push({
      name: "ur-api",
      path: topPath,
      description: descMatch ? descMatch[1] : "ur-api skill index",
    });
  }
  for (const app of APPS) {
    const p = path.join(PACKAGE_ROOT, app, "SKILL.md");
    if (fs.existsSync(p)) {
      const content = fs.readFileSync(p, "utf-8");
      const descMatch = content.match(/description:\s*"([^"]+)"/);
      result.push({
        name: app,
        path: p,
        description: descMatch ? descMatch[1] : app,
      });
    }
  }
  return result;
}

/**
 * 获取指定 app 的 swagger-index.md 内容
 * @param {string} app - app 名称，如 "ur-iot"
 * @returns {string|null}
 */
function getSwaggerIndex(app) {
  if (!APPS.includes(app)) return null;
  const p = path.join(PACKAGE_ROOT, app, "swagger-index.md");
  return fs.existsSync(p) ? fs.readFileSync(p, "utf-8") : null;
}

module.exports = {
  APPS,
  resolveSkillPath,
  getSkillMarkdown,
  listSkills,
  getSwaggerIndex,
};
