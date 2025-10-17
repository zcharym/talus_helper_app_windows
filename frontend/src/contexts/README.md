# Theme Context

A React context for managing dark mode and theme switching in the Talus Helper application.

## Features

- ✅ **Theme Management**: Light, Dark, and Auto (system) themes
- ✅ **Persistent Settings**: Theme preference saved to backend config
- ✅ **System Integration**: Auto theme follows system preference
- ✅ **Real-time Switching**: Instant theme changes without page reload
- ✅ **TypeScript Support**: Full type safety and IntelliSense

## Usage

### Basic Usage

```tsx
import { useTheme } from '../contexts/ThemeContext'

function MyComponent() {
  const { theme, isDark, setTheme, toggleTheme } = useTheme()

  return (
    <div>
      <p>Current theme: {theme}</p>
      <p>Is dark mode: {isDark ? 'Yes' : 'No'}</p>
      <button onClick={() => setTheme('dark')}>Set Dark</button>
      <button onClick={toggleTheme}>Toggle Theme</button>
    </div>
  )
}
```

### Theme Options

- **`light`**: Always use light theme
- **`dark`**: Always use dark theme  
- **`auto`**: Follow system preference (default)

### Context API

#### `useTheme()` Hook

Returns an object with the following properties:

| Property | Type | Description |
|----------|------|-------------|
| `theme` | `'light' \| 'dark' \| 'auto'` | Current theme setting |
| `isDark` | `boolean` | Whether dark mode is currently active |
| `setTheme` | `(theme: Theme) => void` | Set the theme preference |
| `toggleTheme` | `() => void` | Toggle between light and dark |

### Theme Provider Setup

The `ThemeProvider` must wrap your app to provide theme context:

```tsx
import { ThemeProvider } from './contexts/ThemeContext'

function App() {
  return (
    <ThemeProvider>
      {/* Your app components */}
    </ThemeProvider>
  )
}
```

## Implementation Details

### Theme Persistence

- Theme preference is automatically saved to the backend config
- Uses the existing `GetConfig` and `SaveConfig` Wails bindings
- Theme is loaded on app startup and applied immediately

### System Theme Detection

- Uses `window.matchMedia('(prefers-color-scheme: dark)')` for system detection
- Automatically updates when system theme changes (in auto mode)
- Listens for system theme changes and updates accordingly

### CSS Integration

- Uses Tailwind CSS `dark:` classes for styling
- Applies `dark` class to `document.documentElement` when dark mode is active
- Smooth transitions between themes with CSS transitions

### Theme Toggle Component

A convenient theme toggle component is available:

```tsx
import ThemeToggle from './components/ThemeToggle'

// Basic toggle (icon only)
<ThemeToggle />

// Toggle with label
<ThemeToggle showLabel={true} />
```

## Styling Guidelines

### CSS Classes

Use the predefined CSS classes for consistent theming:

```css
/* Form elements */
.form-label          /* Labels */
.form-description    /* Help text */

/* Navigation */
.nav-bg              /* Navigation background */
.nav-link            /* Navigation links */
.nav-link-active     /* Active navigation links */
.nav-title           /* Navigation title */

/* Messages */
.message-success     /* Success messages */
.message-error       /* Error messages */

/* Toggles */
.toggle-bg           /* Toggle background */
.toggle-checked      /* Checked toggle state */

/* Badges */
.badge-primary       /* Primary badges */
```

### Custom Styling

For custom components, use Tailwind's dark mode classes:

```tsx
<div className="bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100">
  Content
</div>
```

## Examples

### Theme-Aware Component

```tsx
import { useTheme } from '../contexts/ThemeContext'

function ThemeAwareCard() {
  const { isDark } = useTheme()
  
  return (
    <div className={`card ${isDark ? 'border-gray-700' : 'border-gray-200'}`}>
      <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100">
        Theme Aware Card
      </h3>
      <p className="text-gray-600 dark:text-gray-400">
        This card adapts to the current theme
      </p>
    </div>
  )
}
```

### Settings Integration

```tsx
import { useTheme } from '../contexts/ThemeContext'

function Settings() {
  const { theme, setTheme } = useTheme()
  
  const handleThemeChange = (newTheme) => {
    setTheme(newTheme)
    // Theme change is automatically saved to config
  }
  
  return (
    <select value={theme} onChange={(e) => handleThemeChange(e.target.value)}>
      <option value="light">Light</option>
      <option value="dark">Dark</option>
      <option value="auto">Auto</option>
    </select>
  )
}
```
