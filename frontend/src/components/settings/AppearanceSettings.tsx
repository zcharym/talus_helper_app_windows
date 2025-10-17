import { useState, useEffect } from 'react'
import { GetConfig, SaveConfig } from '@wailsjs/go/main/App'
import { AppConfig } from '../../types'
import { Sun, Moon, Monitor } from 'lucide-react'
import { useTheme } from '../../contexts/ThemeContext'

function AppearanceSettings() {
  const { setTheme } = useTheme()
  const [config, setConfig] = useState<AppConfig | null>(null)
  const [loading, setLoading] = useState(true)
  const [saving, setSaving] = useState(false)
  const [message, setMessage] = useState<{ type: 'success' | 'error', text: string } | null>(null)

  useEffect(() => {
    loadConfig()
  }, [])

  const loadConfig = async () => {
    try {
      setLoading(true)
      const configData = await GetConfig()
      setConfig(configData)
    } catch (error) {
      console.error('Failed to load config:', error)
      setMessage({ type: 'error', text: 'Failed to load configuration' })
    } finally {
      setLoading(false)
    }
  }

  const handleConfigChange = (key: keyof AppConfig, value: any) => {
    if (!config) return
    
    setConfig({ ...config, [key]: value })
    
    // Handle theme change immediately
    if (key === 'Theme') {
      setTheme(value)
    }
  }

  const handleSaveConfig = async () => {
    if (!config) return

    try {
      setSaving(true)
      await SaveConfig(config)
      setMessage({ type: 'success', text: 'Appearance settings saved successfully!' })
      setTimeout(() => setMessage(null), 3000)
    } catch (error) {
      console.error('Failed to save config:', error)
      setMessage({ type: 'error', text: 'Failed to save appearance settings' })
    } finally {
      setSaving(false)
    }
  }

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
      </div>
    )
  }

  if (!config) {
    return (
      <div className="text-center py-12">
        <p className="text-gray-500 dark:text-gray-400">Failed to load configuration</p>
      </div>
    )
  }

  return (
    <div className="max-w-2xl">
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100 mb-2">
          Appearance Settings
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          Customize the look and feel of your application
        </p>
      </div>

      {message && (
        <div className={`mb-6 p-4 rounded-lg flex items-center gap-2 ${
          message.type === 'success' ? 'message-success' : 'message-error'
        }`}>
          {message.type === 'success' ? (
            <div className="w-4 h-4 rounded-full bg-green-500"></div>
          ) : (
            <div className="w-4 h-4 rounded-full bg-red-500"></div>
          )}
          {message.text}
        </div>
      )}

      <div className="space-y-6">
        {/* Theme Settings */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">
            Theme
          </h3>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium form-label mb-3">
                Choose your preferred theme
              </label>
              <div className="grid grid-cols-3 gap-3">
                <button
                  onClick={() => handleConfigChange('Theme', 'light')}
                  className={`flex flex-col items-center gap-3 p-4 rounded-lg border-2 transition-colors ${
                    config.Theme === 'light'
                      ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20 text-primary-700 dark:text-primary-300'
                      : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
                  }`}
                >
                  <Sun className="w-6 h-6" />
                  <span className="text-sm font-medium">Light</span>
                  <span className="text-xs text-gray-500 dark:text-gray-400">Always light</span>
                </button>
                <button
                  onClick={() => handleConfigChange('Theme', 'dark')}
                  className={`flex flex-col items-center gap-3 p-4 rounded-lg border-2 transition-colors ${
                    config.Theme === 'dark'
                      ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20 text-primary-700 dark:text-primary-300'
                      : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
                  }`}
                >
                  <Moon className="w-6 h-6" />
                  <span className="text-sm font-medium">Dark</span>
                  <span className="text-xs text-gray-500 dark:text-gray-400">Always dark</span>
                </button>
                <button
                  onClick={() => handleConfigChange('Theme', 'auto')}
                  className={`flex flex-col items-center gap-3 p-4 rounded-lg border-2 transition-colors ${
                    config.Theme === 'auto'
                      ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20 text-primary-700 dark:text-primary-300'
                      : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
                  }`}
                >
                  <Monitor className="w-6 h-6" />
                  <span className="text-sm font-medium">Auto</span>
                  <span className="text-xs text-gray-500 dark:text-gray-400">Follow system</span>
                </button>
              </div>
              <p className="text-sm form-description mt-3">
                Auto theme will automatically switch between light and dark based on your system preference
              </p>
            </div>
          </div>
        </div>

        {/* Save Button */}
        <div className="flex justify-end">
          <button
            onClick={handleSaveConfig}
            disabled={saving}
            className="btn-primary flex items-center gap-2"
          >
            {saving ? (
              <div className="w-4 h-4 border-2 border-gray-300 border-t-gray-600 rounded-full animate-spin" />
            ) : (
              <div className="w-4 h-4 rounded-full bg-white"></div>
            )}
            {saving ? 'Saving...' : 'Save Appearance Settings'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default AppearanceSettings
