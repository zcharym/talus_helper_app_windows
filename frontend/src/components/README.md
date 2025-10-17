# Tabs Component

A reusable navigation tabs component for the Talus Helper application.

## Features

- ✅ **Responsive Design**: Desktop and mobile-friendly navigation
- ✅ **Icon Support**: Lucide icons integration
- ✅ **Badge Support**: Optional badges for notifications or counts
- ✅ **Active State**: Automatic active state based on current route
- ✅ **Smooth Transitions**: CSS transitions for better UX
- ✅ **Accessibility**: Proper ARIA labels and keyboard navigation

## Usage

### Basic Usage

```tsx
import Tabs, { TabItem } from './components/Tabs'
import { CheckSquare, Settings } from 'lucide-react'

const navigationTabs: TabItem[] = [
  {
    id: 'todos',
    label: 'Todo List',
    path: '/',
    icon: <CheckSquare className="w-4 h-4" />
  },
  {
    id: 'settings',
    label: 'Settings',
    path: '/settings',
    icon: <Settings className="w-4 h-4" />
  }
]

function App() {
  return (
    <Router>
      <Tabs tabs={navigationTabs} />
      {/* Your routes */}
    </Router>
  )
}
```

### With Badges

```tsx
import { CheckSquare, Bell } from 'lucide-react'

const tabsWithBadges: TabItem[] = [
  {
    id: 'todos',
    label: 'Todo List',
    path: '/',
    icon: <CheckSquare className="w-4 h-4" />,
    badge: '5' // Shows number of pending todos
  },
  {
    id: 'notifications',
    label: 'Notifications',
    path: '/notifications',
    icon: <Bell className="w-4 h-4" />,
    badge: '3' // Shows number of unread notifications
  }
]
```

### Adding More Tabs

To add new tabs, simply add them to the `navigationTabs` array:

```tsx
import { CheckSquare, Folder, BarChart3, Settings } from 'lucide-react'

const navigationTabs: TabItem[] = [
  {
    id: 'todos',
    label: 'Todo List',
    path: '/',
    icon: <CheckSquare className="w-4 h-4" />
  },
  {
    id: 'projects', // New tab
    label: 'Projects',
    path: '/projects',
    icon: <Folder className="w-4 h-4" />
  },
  {
    id: 'analytics', // New tab
    label: 'Analytics',
    path: '/analytics',
    icon: <BarChart3 className="w-4 h-4" />
  },
  {
    id: 'settings',
    label: 'Settings',
    path: '/settings',
    icon: <Settings className="w-4 h-4" />
  }
]
```

## Props

### TabsProps

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `tabs` | `TabItem[]` | - | Array of tab items to display |
| `className` | `string` | `''` | Additional CSS classes |
| `showMobileMenu` | `boolean` | `true` | Whether to show mobile menu button |

### TabItem

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | `string` | ✅ | Unique identifier for the tab |
| `label` | `string` | ✅ | Display text for the tab |
| `path` | `string` | ✅ | Route path for navigation |
| `icon` | `ReactNode` | ✅ | Icon component to display |
| `badge` | `string \| number` | ❌ | Optional badge content |

## Mobile Responsiveness

The component automatically handles mobile responsiveness:

- **Desktop**: Horizontal tab navigation
- **Mobile**: Collapsible hamburger menu
- **Tablet**: Responsive breakpoints for optimal display

## Styling

The component uses Tailwind CSS classes and follows the design system:

- **Active State**: Primary color border and text
- **Hover State**: Gray border and text
- **Badges**: Primary color background with white text
- **Mobile Menu**: Slide-down animation with proper spacing

## Examples

See `ExampleTabs.tsx` for more complex examples including:
- Multiple tabs with badges
- Different icon combinations
- Various styling options
