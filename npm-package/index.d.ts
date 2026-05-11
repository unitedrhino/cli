/**
 * @unitedrhino/ur-api-skills
 * TypeScript 类型声明
 */

export declare const APPS: string[];

export declare function resolveSkillPath(name: string): string | null;

export declare function getSkillMarkdown(name: string): string | null;

export interface SkillMeta {
  name: string;
  path: string;
  description: string;
}

export declare function listSkills(): SkillMeta[];

export declare function getSwaggerIndex(app: string): string | null;
