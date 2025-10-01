import React, { useState } from 'react';
import { Template, ProjectType } from '../types';

interface TemplatePreviewProps {
  templates: Template[];
  loading: boolean;
  projectType: ProjectType;
}

export const TemplatePreview: React.FC<TemplatePreviewProps> = ({
  templates,
  loading,
  projectType
}) => {
  const [selectedTemplate, setSelectedTemplate] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');

  const filteredTemplates = templates.filter(template =>
    template.path.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getFileIcon = (path: string) => {
    if (path.endsWith('.go')) return '🐹';
    if (path.endsWith('.md')) return '📝';
    if (path.endsWith('.yaml') || path.endsWith('.yml')) return '⚙️';
    if (path === 'Makefile') return '🔨';
    if (path === '.gitignore') return '🚫';
    if (path === 'go.mod') return '📦';
    return '📄';
  };

  const getLanguage = (path: string) => {
    if (path.endsWith('.go')) return 'go';
    if (path.endsWith('.md')) return 'markdown';
    if (path.endsWith('.yaml') || path.endsWith('.yml')) return 'yaml';
    if (path === 'Makefile') return 'makefile';
    return 'text';
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-12">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <span className="ml-2 text-gray-600">Generating templates...</span>
      </div>
    );
  }

  if (templates.length === 0) {
    return (
      <div className="text-gray-500 text-center py-8">
        <div className="text-4xl mb-4">📁</div>
        <p>Select a project type and click "Preview Templates" to see the generated files</p>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {/* Search and Stats */}
      <div className="flex justify-between items-center">
        <div className="flex items-center space-x-4">
          <input
            type="text"
            placeholder="Search files..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          <span className="text-sm text-gray-600">
            {filteredTemplates.length} files
          </span>
        </div>
        <div className="text-sm text-gray-600">
          Project Type: <span className="font-medium capitalize">{projectType}</span>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {/* File List */}
        <div className="space-y-2">
          <h4 className="font-medium text-gray-700 mb-2">Generated Files</h4>
          <div className="border border-gray-200 rounded-lg max-h-96 overflow-y-auto">
            {filteredTemplates.map((template, index) => (
              <button
                key={index}
                onClick={() => setSelectedTemplate(template.path)}
                className={`w-full text-left p-3 border-b border-gray-100 hover:bg-gray-50 transition-colors ${
                  selectedTemplate === template.path ? 'bg-blue-50 border-blue-200' : ''
                }`}
              >
                <div className="flex items-center space-x-2">
                  <span className="text-lg">{getFileIcon(template.path)}</span>
                  <span className="font-mono text-sm">{template.path}</span>
                </div>
                <div className="text-xs text-gray-500 mt-1">
                  {template.content.split('\n').length} lines
                </div>
              </button>
            ))}
          </div>
        </div>

        {/* File Content */}
        <div className="space-y-2">
          <h4 className="font-medium text-gray-700 mb-2">
            {selectedTemplate ? `Content: ${selectedTemplate}` : 'Select a file to preview'}
          </h4>
          <div className="border border-gray-200 rounded-lg">
            {selectedTemplate ? (
              <div className="max-h-96 overflow-auto">
                <div className="bg-gray-50 px-3 py-2 border-b border-gray-200 flex justify-between items-center">
                  <span className="font-mono text-sm text-gray-600">{selectedTemplate}</span>
                  <button
                    onClick={() => {
                      const template = templates.find(t => t.path === selectedTemplate);
                      if (template) {
                        navigator.clipboard.writeText(template.content);
                      }
                    }}
                    className="text-xs bg-gray-200 hover:bg-gray-300 px-2 py-1 rounded transition-colors"
                  >
                    Copy
                  </button>
                </div>
                <pre className="p-4 text-sm overflow-x-auto">
                  <code className={`language-${getLanguage(selectedTemplate)}`}>
                    {templates.find(t => t.path === selectedTemplate)?.content}
                  </code>
                </pre>
              </div>
            ) : (
              <div className="p-8 text-center text-gray-500">
                <div className="text-4xl mb-4">👆</div>
                <p>Click on a file to preview its content</p>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Template Statistics */}
      <div className="bg-gray-50 rounded-lg p-4">
        <h4 className="font-medium text-gray-700 mb-2">Template Statistics</h4>
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm">
          <div>
            <div className="text-gray-600">Total Files</div>
            <div className="font-semibold">{templates.length}</div>
          </div>
          <div>
            <div className="text-gray-600">Go Files</div>
            <div className="font-semibold">
              {templates.filter(t => t.path.endsWith('.go')).length}
            </div>
          </div>
          <div>
            <div className="text-gray-600">Config Files</div>
            <div className="font-semibold">
              {templates.filter(t => 
                t.path.endsWith('.yaml') || 
                t.path.endsWith('.yml') || 
                t.path === 'go.mod' ||
                t.path === 'Makefile'
              ).length}
            </div>
          </div>
          <div>
            <div className="text-gray-600">Total Lines</div>
            <div className="font-semibold">
              {templates.reduce((acc, t) => acc + t.content.split('\n').length, 0)}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};