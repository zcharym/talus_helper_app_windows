# Settings Components

A modern, sidebar-based settings interface for the Talus Helper application with organized sections and improved user experience.

## Architecture

The settings page is now organized into a sidebar navigation with individual sections for different categories:

```
Settings/
├── SettingsSidebar.tsx      # Sidebar navigation component
├── AppearanceSettings.tsx   # Theme and visual preferences
├── TodoSettings.tsx         # Todo list configuration
├── OCRSettings.tsx          # OpenAI API and OCR settings
├── GeneralSettings.tsx      # General application preferences
├── LanguageSettings.tsx     # Interface language settings
└── README.md               # This documentation
```

## Features

### 🎨 **Appearance Settings**
- **Theme Selection**: Light, Dark, and Auto themes with visual previews
- **Real-time Switching**: Instant theme changes without page reload
- **System Integration**: Auto theme follows system preference
- **Visual Feedback**: Clear indication of current theme selection

### ✅ **Todo Settings**
- **Default Category**: Set default category for new todos
- **Todo Limits**: Configure maximum number of todos
- **Auto-save Behavior**: Toggle automatic saving of changes
- **Notifications**: Enable/disable system notifications

### 👁️ **OCR Settings**
- **API Configuration**: OpenAI API key and base URL setup
- **Provider Support**: Information about supported providers
- **Connection Testing**: Test API connectivity
- **Usage Guide**: Clear explanation of how OCR works

### ⚙️ **General Settings**
- **Auto-save Toggle**: Control automatic saving behavior
- **Notification Preferences**: System notification settings
- **App Information**: Version, build, and platform details
- **Application Metadata**: Framework and system information

### 🌍 **Language Settings**
- **Multi-language Support**: 10+ languages with flags and names
- **Visual Selection**: Flag-based language picker
- **Current Language Display**: Clear indication of active language
- **Language Information**: Details about what changes with language

## Component Structure

### SettingsSidebar
```tsx
interface SettingsSection {
  id: string
  label: string
  path: string
  icon: React.ReactNode
  description: string
}
```

**Features:**
- Active section highlighting
- Icon-based navigation
- Responsive design
- Dark mode support

### Individual Settings Components
Each settings section follows a consistent pattern:

```tsx
function SettingsSection() {
  const [config, setConfig] = useState<AppConfig | null>(null)
  const [loading, setLoading] = useState(true)
  const [saving, setSaving] = useState(false)
  const [message, setMessage] = useState<Message | null>(null)

  // Load config on mount
  useEffect(() => {
    loadConfig()
  }, [])

  // Handle config changes
  const handleConfigChange = (key: keyof AppConfig, value: any) => {
    if (!config) return
    setConfig({ ...config, [key]: value })
  }

  // Save configuration
  const handleSaveConfig = async () => {
    // Save logic with error handling
  }
}
```

## Routing

The settings use nested routing for clean URLs:

```tsx
// Main App routing
<Route path="/settings/*" element={<Settings />} />

// Settings internal routing
<Routes>
  <Route path="/" element={<Navigate to="/settings/appearance" replace />} />
  <Route path="/appearance" element={<AppearanceSettings />} />
  <Route path="/todos" element={<TodoSettings />} />
  <Route path="/ocr" element={<OCRSettings />} />
  <Route path="/general" element={<GeneralSettings />} />
  <Route path="/language" element={<LanguageSettings />} />
</Routes>
```

## Styling

### CSS Classes Used
- **Layout**: `flex`, `h-full`, `overflow-auto`
- **Cards**: `card` (predefined component class)
- **Forms**: `form-label`, `form-description`, `input-field`
- **Buttons**: `btn-primary`, `btn-secondary`
- **Messages**: `message-success`, `message-error`
- **Toggles**: `toggle-bg`, `toggle-checked`

### Dark Mode Support
All components include comprehensive dark mode styling:
- Automatic theme detection
- Smooth transitions between themes
- Consistent color schemes
- Accessible contrast ratios

## User Experience

### Navigation
- **Sidebar Navigation**: Easy access to all settings sections
- **Active State**: Clear indication of current section
- **Descriptions**: Helpful descriptions for each section
- **Icons**: Visual icons for quick recognition

### Form Interactions
- **Real-time Updates**: Changes apply immediately where appropriate
- **Save Feedback**: Clear success/error messages
- **Loading States**: Visual feedback during operations
- **Validation**: Input validation and error handling

### Responsive Design
- **Mobile Friendly**: Responsive layout for all screen sizes
- **Touch Optimized**: Large touch targets for mobile devices
- **Accessible**: Proper ARIA labels and keyboard navigation

## Configuration Management

### Backend Integration
- **Wails Bindings**: Uses `GetConfig` and `SaveConfig` from backend
- **Type Safety**: Full TypeScript support with `AppConfig` interface
- **Error Handling**: Comprehensive error handling and user feedback
- **Persistence**: Settings automatically saved to backend storage

### State Management
- **Local State**: Each component manages its own loading/saving state
- **Config Synchronization**: Real-time updates across components
- **Theme Integration**: Seamless integration with theme context

## Usage Examples

### Adding a New Settings Section

1. **Create the component**:
```tsx
// NewSettingsSection.tsx
function NewSettingsSection() {
  // Follow the standard pattern
}
```

2. **Add to sidebar**:
```tsx
// SettingsSidebar.tsx
export const settingsSections: SettingsSection[] = [
  // ... existing sections
  {
    id: 'new-section',
    label: 'New Section',
    path: '/settings/new-section',
    icon: <NewIcon className="w-4 h-4" />,
    description: 'Description of new section'
  }
]
```

3. **Add routing**:
```tsx
// Settings.tsx
<Route path="/new-section" element={<NewSettingsSection />} />
```

### Customizing Existing Sections

Each section is self-contained and can be easily modified:
- **Add new fields**: Extend the form with additional inputs
- **Modify layout**: Change the visual arrangement
- **Add validation**: Implement custom validation logic
- **Enhance UX**: Add tooltips, help text, or interactive elements

## Benefits

### For Users
- **Better Organization**: Settings grouped logically by function
- **Easier Navigation**: Sidebar makes finding settings simple
- **Visual Feedback**: Clear indication of current settings and changes
- **Consistent Experience**: Uniform interface across all sections

### For Developers
- **Modular Architecture**: Easy to add new settings sections
- **Reusable Components**: Consistent patterns across sections
- **Type Safety**: Full TypeScript support
- **Maintainable Code**: Clear separation of concerns

### For Maintenance
- **Scalable Structure**: Easy to extend with new features
- **Consistent Styling**: Shared CSS classes and patterns
- **Error Handling**: Standardized error handling across sections
- **Testing**: Individual components can be tested in isolation
