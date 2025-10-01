import React, { useState, useEffect } from 'react';
import { ProjectType, Framework, TemplateData, Template } from '../types';
import { TemplatePreview } from './TemplatePreview';
import { ProjectConfiguration } from './ProjectConfiguration';
import { ProjectTypeSidebar } from './ProjectTypeSidebar';
import { generateProject, previewTemplates } from '../service';

export const TemplateManager: React.FC = () => {
  const [selectedProjectType, setSelectedProjectType] = useState<ProjectType>('microservice');
  const [selectedFramework, setSelectedFramework] = useState<Framework>('gin');
  const [templates, setTemplates] = useState<Template[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  
  const [templateData, setTemplateData] = useState<TemplateData>({
    projectName: '',
    moduleName: '',
    description: '',
    framework: 'gin',
    goVersion: '1.22.0',
    projectType: 'microservice'
  });

  // Update template data when selections change
  useEffect(() => {
    setTemplateData(prev => ({
      ...prev,
      projectType: selectedProjectType,
      framework: selectedFramework
    }));
  }, [selectedProjectType, selectedFramework]);

  const handlePreviewTemplates = async () => {
    setLoading(true);
    setError(null);
    
    try {
      const response = await previewTemplates(templateData);
      setTemplates(response.templates);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to generate templates');
    } finally {
      setLoading(false);
    }
  };

  const handleDownloadProject = async () => {
    setLoading(true);
    setError(null);
    
    try {
      const blob = await generateProject(templateData);
      
      // Create download link
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.style.display = 'none';
      a.href = url;
      a.download = `${templateData.projectName}.zip`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to download project');
    } finally {
      setLoading(false);
    }
  };

  const handleConfigurationChange = (newData: Partial<TemplateData>) => {
    setTemplateData(prev => ({ ...prev, ...newData }));
  };

  return (
    <div className="min-h-screen bg-gray-100">
      <div className="container mx-auto px-4 py-8">
        <header className="mb-8">
          <h1 className="text-4xl font-bold text-gray-800 mb-2">Go Template Manager</h1>
          <p className="text-gray-600">Manage and preview Go project templates</p>
        </header>

        <div className="flex gap-6">
          {/* Sidebar */}
          <ProjectTypeSidebar
            selectedProjectType={selectedProjectType}
            selectedFramework={selectedFramework}
            onProjectTypeChange={setSelectedProjectType}
            onFrameworkChange={setSelectedFramework}
            onPreview={handlePreviewTemplates}
            loading={loading}
          />

          {/* Main Content */}
          <div className="flex-1">
            <div className="bg-white rounded-lg shadow-md">
              {/* Configuration */}
              <ProjectConfiguration
                templateData={templateData}
                onChange={handleConfigurationChange}
              />

              {/* Template Preview */}
              <div className="p-6">
                <div className="flex justify-between items-center mb-4">
                  <h3 className="text-lg font-semibold">Template Preview</h3>
                  <div className="flex gap-2">
                    <button
                      onClick={handleDownloadProject}
                      disabled={loading}
                      className="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors disabled:opacity-50"
                    >
                      {loading ? 'Generating...' : 'Download Project'}
                    </button>
                    <button
                      onClick={handlePreviewTemplates}
                      disabled={loading}
                      className="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors disabled:opacity-50"
                    >
                      Refresh
                    </button>
                  </div>
                </div>

                {error && (
                  <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-4">
                    {error}
                  </div>
                )}

                <TemplatePreview
                  templates={templates}
                  loading={loading}
                  projectType={selectedProjectType}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

