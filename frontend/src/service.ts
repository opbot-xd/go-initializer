// src/service.ts
// Standard HTTP methods for API calls

import { CreateProjectRequest, ProjectMetadata, Template, TemplateData } from './types';


const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8181';

export async function get<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
  });
  if (!response.ok) throw new Error(await response.text());
  return response.json();
}

export async function post<T>(endpoint: string, data: any, options: RequestInit = {}): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) throw new Error(await response.text());
  return response.json();
}

export async function put<T>(endpoint: string, data: any, options: RequestInit = {}): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) throw new Error(await response.text());
  return response.json();
}

export async function del<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
  });
  if (!response.ok) throw new Error(await response.text());
  return response.json();
}

export async function generateProject(data: TemplateData): Promise<Blob> {
  const request: CreateProjectRequest = {
    projectType: data.projectType,
    goVersion: data.goVersion,
    framework: data.framework,
    moduleName: data.moduleName,
    name: data.projectName,
    description: data.description,
  };

  const response = await fetch(`${API_BASE_URL}/api/generate`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(request),
  });
  
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(`Failed to generate project: ${errorText}`);
  }
  
  return await response.blob();
}

export async function getMetaData(): Promise<ProjectMetadata> {
  const response = await get<ProjectMetadata>('/api/meta');
  return response;
}

// Template preview API
export async function previewTemplates(data: TemplateData): Promise<{ templates: Template[], count: number }> {
  const request = {
    projectType: data.projectType,
    framework: data.framework,
    projectName: data.projectName,
    moduleName: data.moduleName,
    description: data.description,
    goVersion: data.goVersion,
  };
  
  const response = await post<{ templates: Template[], count: number }>('/api/preview', request);
  return response;
}

// Template management APIs
export async function getAvailableTemplates(): Promise<string[]> {
  const response = await get<string[]>('/api/templates');
  return response;
}

export async function getTemplateContent(templatePath: string): Promise<string> {
  const response = await get<{ content: string }>(`/api/templates/${encodeURIComponent(templatePath)}`);
  return response.content;
}

export async function updateTemplate(templatePath: string, content: string): Promise<void> {
  await put(`/api/templates/${encodeURIComponent(templatePath)}`, { content });
}

export async function createTemplate(templatePath: string, content: string): Promise<void> {
  await post('/api/templates', { path: templatePath, content });
}

export async function deleteTemplate(templatePath: string): Promise<void> {
  await del(`/api/templates/${encodeURIComponent(templatePath)}`);
}