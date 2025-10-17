import { ReactNode, useState } from 'react'
import { Link, useLocation } from 'react-router-dom'


export const navigationTabs: TabItem[] = [
  {
    id: 'todos',
    label: 'Todo List',
    path: '/',
    icon: <span className="icon-[tabler--check] w-4 h-4" />
  },
  {
    id: 'settings',
    label: 'Settings',
    path: '/settings',
    icon: <span className="icon-[tabler--settings] w-4 h-4" />
  }
]

export interface TabItem {
  id: string
  label: string
  path: string
  icon: ReactNode
  badge?: string | number
}

interface TabsProps {
  tabs: TabItem[]
  className?: string
  showMobileMenu?: boolean
}

function Tabs({ tabs, className = '', showMobileMenu = true }: TabsProps) {
  const location = useLocation()
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)

  return (
    <nav className={`bg-white shadow-sm border-b border-gray-200 ${className}`}>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex">
            <div className="flex-shrink-0 flex items-center">
              <h1 className="text-xl font-bold text-gray-900">Talus</h1>
            </div>
            
            {/* Desktop Navigation */}
            <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
              {tabs.map((tab) => {
                const isActive = location.pathname === tab.path
                return (
                  <Link
                    key={tab.id}
                    to={tab.path}
                    className={`inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium transition-colors duration-200 ${
                      isActive
                        ? 'border-primary-500 text-gray-900'
                        : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700'
                    }`}
                  >
                    <span className="mr-2">{tab.icon}</span>
                    {tab.label}
                    {tab.badge && (
                      <span className="ml-2 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-primary-100 text-primary-800">
                        {tab.badge}
                      </span>
                    )}
                  </Link>
                )
              })}
            </div>
          </div>

          {/* Mobile menu button */}
          {showMobileMenu && (
            <div className="sm:hidden flex items-center">
              <button
                onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
                className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500"
                aria-expanded="false"
              >
                <span className="sr-only">Open main menu</span>
                <span className="icon-[tabler--menu-2] w-6 h-6" />
              </button>
            </div>
          )}
        </div>

        {/* Mobile Navigation Menu */}
        {showMobileMenu && mobileMenuOpen && (
          <div className="sm:hidden">
            <div className="pt-2 pb-3 space-y-1">
              {tabs.map((tab) => {
                const isActive = location.pathname === tab.path
                return (
                  <Link
                    key={tab.id}
                    to={tab.path}
                    onClick={() => setMobileMenuOpen(false)}
                    className={`block pl-3 pr-4 py-2 border-l-4 text-base font-medium transition-colors duration-200 ${
                      isActive
                        ? 'bg-primary-50 border-primary-500 text-primary-700'
                        : 'border-transparent text-gray-500 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-700'
                    }`}
                  >
                    <div className="flex items-center">
                      <span className="mr-3">{tab.icon}</span>
                      {tab.label}
                      {tab.badge && (
                        <span className="ml-auto inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-primary-100 text-primary-800">
                          {tab.badge}
                        </span>
                      )}
                    </div>
                  </Link>
                )
              })}
            </div>
          </div>
        )}
      </div>
    </nav>
  )
}

export default Tabs
