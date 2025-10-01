// Project configuration types
export type ProjectType = 'microservice' | 'cli-app' | 'api-server' | 'simple-project';

export type Framework = 
  | 'gin' | 'echo' | 'fiber' | 'chi' | 'gokit'  // Web frameworks
  | 'cobra' | 'urfave' | 'kingpin'              // CLI frameworks
  | 'golly';                                     // Simple/default

export interface TemplateData {
  projectName: string;
  moduleName: string;
  description: string;
  framework: Framework;
  goVersion: string;
  projectType: ProjectType;
}

export interface Template {
  path: string;
  content: string;
}

// API request/response types
export interface CreateProjectRequest {
  projectType: ProjectType;
  goVersion: string;
  framework: Framework;
  moduleName: string;
  name: string;
  description: string;
}

export interface ProjectMetadata {
  supportedProjectTypes: Record<string, string>;
  supportedGoVersions: Record<string, boolean>;
  supportedFrameworks: Record<ProjectType, Record<Framework, boolean>>;
}

// Template management types
export interface TemplateFile {
  path: string;
  content: string;
  language: string;
  size: number;
}

export interface TemplatePreviewData {
  templates: TemplateFile[];
  projectStructure: string[];
  dependencies: string[];
  features: string[];
}