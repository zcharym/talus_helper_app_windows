import { Link, useLocation } from 'react-router-dom'
import { Palette, CheckSquare, Eye, Globe, Settings as SettingsIcon } from 'lucide-react'

export interface SettingsSection {
  id: string
  label: string
  path: string
  icon: React.ReactNode
  description: string
}

export const settingsSections: SettingsSection[] = [
  {
    id: 'appearance',
    label: 'Appearance',
    path: '/settings/appearance',
    icon: <Palette className="w-4 h-4" />,
    description: 'Theme, colors, and visual preferences'
  },
  {
    id: 'todos',
    label: 'Todo Settings',
    path: '/settings/todos',
    icon: <CheckSquare className="w-4 h-4" />,
    description: 'Todo list configuration and behavior'
  },
  {
    id: 'ocr',
    label: 'OCR Settings',
    path: '/settings/ocr',
    icon: <Eye className="w-4 h-4" />,
    description: 'OpenAI API and image recognition settings'
  },
  {
    id: 'general',
    label: 'General',
    path: '/settings/general',
    icon: <SettingsIcon className="w-4 h-4" />,
    description: 'General application preferences'
  },
  {
    id: 'language',
    label: 'Language',
    path: '/settings/language',
    icon: <Globe className="w-4 h-4" />,
    description: 'Interface language and localization'
  }
]

interface SettingsSidebarProps {
  className?: string
}

function SettingsSidebar({ className = '' }: SettingsSidebarProps) {
  const location = useLocation()

  return (
    <div className={`w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 ${className}`}>
      <div className="p-6">
        <h2 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-6">
          Settings
        </h2>
        
        <nav className="space-y-1">
          {settingsSections.map((section) => {
            const isActive = location.pathname === section.path
            return (
              <Link
                key={section.id}
                to={section.path}
                className={`group flex items-start gap-3 p-3 rounded-lg transition-colors duration-200 ${
                  isActive
                    ? 'bg-primary-50 dark:bg-primary-900/20 border border-primary-200 dark:border-primary-800'
                    : 'hover:bg-gray-50 dark:hover:bg-gray-700'
                }`}
              >
                <div className={`flex-shrink-0 p-1.5 rounded-md ${
                  isActive
                    ? 'bg-primary-100 dark:bg-primary-800/30 text-primary-600 dark:text-primary-400'
                    : 'bg-gray-100 dark:bg-gray-700 text-gray-600 dark:text-gray-400 group-hover:bg-gray-200 dark:group-hover:bg-gray-600'
                }`}>
                  {section.icon}
                </div>
                
                <div className="flex-1 min-w-0">
                  <div className={`text-sm font-medium ${
                    isActive
                      ? 'text-primary-900 dark:text-primary-100'
                      : 'text-gray-900 dark:text-gray-100 group-hover:text-gray-700 dark:group-hover:text-gray-300'
                  }`}>
                    {section.label}
                  </div>
                  <div className={`text-xs mt-0.5 ${
                    isActive
                      ? 'text-primary-700 dark:text-primary-300'
                      : 'text-gray-500 dark:text-gray-400'
                  }`}>
                    {section.description}
                  </div>
                </div>
              </Link>
            )
          })}
        </nav>
      </div>
    </div>
  )
}

export default SettingsSidebar
