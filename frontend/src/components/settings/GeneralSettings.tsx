import { useState, useEffect } from 'react'
import { GetConfig, SaveConfig } from '@wailsjs/go/main/App'
import { AppConfig } from '../../types'
import { Settings as SettingsIcon, Bell, Save } from 'lucide-react'

function GeneralSettings() {
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
  }

  const handleSaveConfig = async () => {
    if (!config) return

    try {
      setSaving(true)
      await SaveConfig(config)
      setMessage({ type: 'success', text: 'General settings saved successfully!' })
      setTimeout(() => setMessage(null), 3000)
    } catch (error) {
      console.error('Failed to save config:', error)
      setMessage({ type: 'error', text: 'Failed to save general settings' })
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
          General Settings
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          Configure general application behavior and preferences
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
        {/* Auto-save Settings */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <Save className="w-5 h-5" />
            Auto-save Behavior
          </h3>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <label className="text-sm font-medium form-label">
                  Auto-save changes
                </label>
                <p className="text-sm form-description">
                  Automatically save todos and settings changes without manual confirmation
                </p>
              </div>
              <label className="relative inline-flex items-center cursor-pointer">
                <input
                  type="checkbox"
                  checked={config.AutoSave}
                  onChange={(e) => handleConfigChange('AutoSave', e.target.checked)}
                  className="sr-only peer"
                />
                <div className="w-11 h-6 toggle-bg peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 dark:after:border-gray-600 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:toggle-checked"></div>
              </label>
            </div>
          </div>
        </div>

        {/* Notification Settings */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <Bell className="w-5 h-5" />
            Notifications
          </h3>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <label className="text-sm font-medium form-label">
                  System notifications
                </label>
                <p className="text-sm form-description">
                  Show system notifications for completed tasks and important events
                </p>
              </div>
              <label className="relative inline-flex items-center cursor-pointer">
                <input
                  type="checkbox"
                  checked={config.Notifications}
                  onChange={(e) => handleConfigChange('Notifications', e.target.checked)}
                  className="sr-only peer"
                />
                <div className="w-11 h-6 toggle-bg peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 dark:after:border-gray-600 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:toggle-checked"></div>
              </label>
            </div>
          </div>
        </div>

        {/* Application Information */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <SettingsIcon className="w-5 h-5" />
            Application Information
          </h3>
          <div className="space-y-3 text-sm">
            <div className="flex justify-between">
              <span className="text-gray-600 dark:text-gray-400">Version:</span>
              <span className="text-gray-900 dark:text-gray-100">1.0.0</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600 dark:text-gray-400">Build:</span>
              <span className="text-gray-900 dark:text-gray-100">Production</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600 dark:text-gray-400">Platform:</span>
              <span className="text-gray-900 dark:text-gray-100">Windows</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600 dark:text-gray-400">Framework:</span>
              <span className="text-gray-900 dark:text-gray-100">Wails v2</span>
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
            {saving ? 'Saving...' : 'Save General Settings'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default GeneralSettings
