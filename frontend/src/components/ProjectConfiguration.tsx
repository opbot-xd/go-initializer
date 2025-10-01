import React from 'react';
import { TemplateData } from '../types';

interface ProjectConfigurationProps {
  templateData: TemplateData;
  onChange: (data: Partial<TemplateData>) => void;
}

export const ProjectConfiguration: React.FC<ProjectConfigurationProps> = ({
  templateData,
  onChange
}) => {
  const handleInputChange = (field: keyof TemplateData) => (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    onChange({ [field]: e.target.value });
  };

  return (
    <div className="p-6 border-b border-gray-200">
      <h2 className="text-xl font-semibold mb-4">Template Configuration</h2>
      
      <div className="grid grid-cols-2 gap-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Project Name
          </label>
          <input
            type="text"
            value={templateData.projectName}
            onChange={handleInputChange('projectName')}
            className="w-full p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter project name"
          />
          <p className="text-xs text-gray-500 mt-1">
            Used for directory names and binary names
          </p>
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Module Name
          </label>
          <input
            type="text"
            value={templateData.moduleName}
            onChange={handleInputChange('moduleName')}
            className="w-full p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter module name"
          />
          <p className="text-xs text-gray-500 mt-1">
            Go module path (e.g., github.com/user/project)
          </p>
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Go Version
          </label>
          <select
            value={templateData.goVersion}
            onChange={handleInputChange('goVersion')}
            className="w-full p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="1.22.0">1.22.0</option>
            <option value="1.21.0">1.21.0</option>
            <option value="1.20.0">1.20.0</option>
            <option value="1.19.0">1.19.0</option>
          </select>
          <p className="text-xs text-gray-500 mt-1">
            Minimum Go version required
          </p>
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Description
          </label>
          <input
            type="text"
            value={templateData.description}
            onChange={handleInputChange('description')}
            className="w-full p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter project description"
          />
          <p className="text-xs text-gray-500 mt-1">
            Brief description of your project
          </p>
        </div>
      </div>

      <div className="mt-4 p-4 bg-gray-50 rounded-lg">
        <h3 className="text-sm font-medium text-gray-700 mb-2">Generated Project Structure</h3>
        <div className="text-xs text-gray-600 font-mono">
          <div>{templateData.projectName}/</div>
          <div>├── cmd/{templateData.projectName}/</div>
          <div>│   └── main.go</div>
          <div>├── internal/</div>
          {(templateData.projectType === 'microservice' || templateData.projectType === 'api-server') && (
            <>
              <div>│   ├── server/</div>
              <div>│   ├── handler/</div>
              <div>│   └── config/</div>
            </>
          )}
          {templateData.projectType === 'cli-app' && (
            <div>│   └── cli/</div>
          )}
          {templateData.projectType === 'simple-project' && (
            <>
              <div>│   └── app/</div>
              <div>├── pkg/{templateData.projectName}/</div>
            </>
          )}
          <div>├── go.mod</div>
          <div>├── .gitignore</div>
          <div>├── Makefile</div>
          <div>└── README.md</div>
        </div>
      </div>
    </div>
  );
};