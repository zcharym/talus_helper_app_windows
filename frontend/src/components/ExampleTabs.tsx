import { TabItem } from './Tabs'
import { CheckSquare, Folder, BarChart3, Users, Settings, LayoutDashboard, ListTodo, Calendar } from 'lucide-react'

// Example of how to easily add more tabs
export const exampleTabsWithMoreSections: TabItem[] = [
  {
    id: 'todos',
    label: 'Todo List',
    path: '/',
    icon: <CheckSquare className="w-4 h-4" />,
    badge: '5' // Example: show number of pending todos
  },
  {
    id: 'projects',
    label: 'Projects',
    path: '/projects',
    icon: <Folder className="w-4 h-4" />
  },
  {
    id: 'analytics',
    label: 'Analytics',
    path: '/analytics',
    icon: <BarChart3 className="w-4 h-4" />
  },
  {
    id: 'team',
    label: 'Team',
    path: '/team',
    icon: <Users className="w-4 h-4" />,
    badge: '3' // Example: show number of team members
  },
  {
    id: 'settings',
    label: 'Settings',
    path: '/settings',
    icon: <Settings className="w-4 h-4" />
  }
]

// Example of tabs with different styling
export const compactTabs: TabItem[] = [
  {
    id: 'dashboard',
    label: 'Dashboard',
    path: '/dashboard',
    icon: <LayoutDashboard className="w-4 h-4" />
  },
  {
    id: 'tasks',
    label: 'Tasks',
    path: '/tasks',
    icon: <ListTodo className="w-4 h-4" />
  },
  {
    id: 'calendar',
    label: 'Calendar',
    path: '/calendar',
    icon: <Calendar className="w-4 h-4" />
  }
]
