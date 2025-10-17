import { TabItem } from './Tabs'

// Example of how to easily add more tabs
export const exampleTabsWithMoreSections: TabItem[] = [
  {
    id: 'todos',
    label: 'Todo List',
    path: '/',
    icon: <span className="icon-[tabler--check] w-4 h-4" />,
    badge: '5' // Example: show number of pending todos
  },
  {
    id: 'projects',
    label: 'Projects',
    path: '/projects',
    icon: <span className="icon-[tabler--folder] w-4 h-4" />
  },
  {
    id: 'analytics',
    label: 'Analytics',
    path: '/analytics',
    icon: <span className="icon-[tabler--chart-bar] w-4 h-4" />
  },
  {
    id: 'team',
    label: 'Team',
    path: '/team',
    icon: <span className="icon-[tabler--users] w-4 h-4" />,
    badge: '3' // Example: show number of team members
  },
  {
    id: 'settings',
    label: 'Settings',
    path: '/settings',
    icon: <span className="icon-[tabler--settings] w-4 h-4" />
  }
]

// Example of tabs with different styling
export const compactTabs: TabItem[] = [
  {
    id: 'dashboard',
    label: 'Dashboard',
    path: '/dashboard',
    icon: <span className="icon-[tabler--dashboard] w-4 h-4" />
  },
  {
    id: 'tasks',
    label: 'Tasks',
    path: '/tasks',
    icon: <span className="icon-[tabler--list-check] w-4 h-4" />
  },
  {
    id: 'calendar',
    label: 'Calendar',
    path: '/calendar',
    icon: <span className="icon-[tabler--calendar] w-4 h-4" />
  }
]
