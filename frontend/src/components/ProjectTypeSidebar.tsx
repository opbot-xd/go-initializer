import React from 'react';
import { ProjectType, Framework } from '../types';

interface ProjectTypeSidebarProps {
  selectedProjectType: ProjectType;
  selectedFramework: Framework;
  onProjectTypeChange: (type: ProjectType) => void;
  onFrameworkChange: (framework: Framework) => void;
  onPreview: () => void;
  loading: boolean;
}

const projectTypes = [
  {
    type: 'microservice' as ProjectType,
    title: 'Microservice',
    description: 'HTTP server with REST API',
    color: 'blue',
    frameworks: ['gin', 'echo', 'fiber', 'chi'] as Framework[]
  },
  {
    type: 'cli-app' as ProjectType,
    title: 'CLI Application',
    description: 'Command-line interface',
    color: 'green',
    frameworks: ['cobra', 'urfave', 'kingpin'] as Framework[]
  },
  {
    type: 'api-server' as ProjectType,
    title: 'API Server',
    description: 'REST API with middleware',
    color: 'purple',
    frameworks: ['gin', 'echo', 'fiber', 'chi'] as Framework[]
  },
  {
    type: 'simple-project' as ProjectType,
    title: 'Simple Project',
    description: 'Basic Go project structure',
    color: 'yellow',
    frameworks: ['golly'] as Framework[]
  }
];

export const ProjectTypeSidebar: React.FC<ProjectTypeSidebarProps> = ({
  selectedProjectType,
  selectedFramework,
  onProjectTypeChange,
  onFrameworkChange,
  onPreview,
  loading
}) => {
  const selectedProject = projectTypes.find(p => p.type === selectedProjectType);
  const availableFrameworks = selectedProject?.frameworks || [];

  // Update framework if current selection is not available for the selected project type
  React.useEffect(() => {
    if (!availableFrameworks.includes(selectedFramework)) {
      onFrameworkChange(availableFrameworks[0]);
    }
  }, [selectedProjectType, availableFrameworks, selectedFramework, onFrameworkChange]);

  return (
    <div className="w-1/4 bg-white rounded-lg shadow-md p-6 min-h-[calc(100vh-12rem)]">
      <h2 className="text-xl font-semibold mb-4">Project Types</h2>
      
      <div className="space-y-2 mb-6">
        {projectTypes.map((project) => (
          <button
            key={project.type}
            onClick={() => onProjectTypeChange(project.type)}
            className={`w-full text-left p-3 rounded-lg transition-colors ${
              selectedProjectType === project.type
                ? `bg-${project.color}-100 border-2 border-${project.color}-300`
                : `bg-${project.color}-50 hover:bg-${project.color}-100 border-2 border-transparent`
            }`}
          >
            <div className="font-medium">{project.title}</div>
            <div className="text-sm text-gray-600">{project.description}</div>
          </button>
        ))}
      </div>

      <div className="mb-6">
        <h3 className="text-lg font-medium mb-3">Framework</h3>
        <select
          value={selectedFramework}
          onChange={(e) => onFrameworkChange(e.target.value as Framework)}
          className="w-full p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          {availableFrameworks.map((framework) => (
            <option key={framework} value={framework}>
              {framework.charAt(0).toUpperCase() + framework.slice(1)}
            </option>
          ))}
        </select>
      </div>

      <div className="mb-6">
        <h3 className="text-lg font-medium mb-3">Features</h3>
        <div className="text-sm text-gray-600 space-y-1">
          {selectedProjectType === 'microservice' && (
            <>
              <div>• HTTP server setup</div>
              <div>• Health check endpoint</div>
              <div>• Graceful shutdown</div>
              <div>• Configuration management</div>
            </>
          )}
          {selectedProjectType === 'cli-app' && (
            <>
              <div>• Command structure</div>
              <div>• Flag parsing</div>
              <div>• Subcommands</div>
              <div>• Help documentation</div>
            </>
          )}
          {selectedProjectType === 'api-server' && (
            <>
              <div>• REST API endpoints</div>
              <div>• CORS middleware</div>
              <div>• Authentication</div>
              <div>• OpenAPI documentation</div>
            </>
          )}
          {selectedProjectType === 'simple-project' && (
            <>
              <div>• Basic structure</div>
              <div>• Package organization</div>
              <div>• Testing setup</div>
              <div>• Build configuration</div>
            </>
          )}
        </div>
      </div>

      <button
        onClick={onPreview}
        disabled={loading}
        className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {loading ? 'Loading...' : 'Preview Templates'}
      </button>
    </div>
  );
};