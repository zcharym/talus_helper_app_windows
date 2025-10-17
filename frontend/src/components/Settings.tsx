import { useState, useEffect } from 'react'
import { GetConfig, SaveConfig } from '@wailsjs/go/main/App'
import { AppConfig } from '../types'
import { Check, AlertCircle, Save, RotateCcw } from 'lucide-react'

const defaultConfig: AppConfig = {
  Theme: 'light',
  AutoSave: true,
  Notifications: true,
  OpenAIAPIKey: '',
  OpenAIBaseURL: 'https://api.moonshot.cn/v1',
  DefaultTodoCategory: 'General',
  MaxTodos: 100,
  Language: 'en'
}

function Settings() {
  const [config, setConfig] = useState<AppConfig>(defaultConfig)
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

  const handleSaveConfig = async () => {
    try {
      setSaving(true)
      await SaveConfig(config)
      setMessage({ type: 'success', text: 'Configuration saved successfully!' })
      setTimeout(() => setMessage(null), 3000)
    } catch (error) {
      console.error('Failed to save config:', error)
      setMessage({ type: 'error', text: 'Failed to save configuration' })
    } finally {
      setSaving(false)
    }
  }

  const handleConfigChange = (key: keyof AppConfig, value: any) => {
    setConfig((prev: AppConfig) => ({ ...prev, [key]: value }))
  }

  const resetToDefaults = () => {
    setConfig(defaultConfig)
  }

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
      </div>
    )
  }

  return (
    <div className="px-4 py-6 sm:px-0">
      <div className="max-w-2xl mx-auto">
        <div className="mb-8">
          <h2 className="text-3xl font-bold text-gray-900 mb-2">Settings</h2>
          <p className="text-gray-600">Configure your Talus Helper preferences</p>
        </div>

        {message && (
          <div className={`mb-6 p-4 rounded-lg flex items-center gap-2 ${
            message.type === 'success' 
              ? 'bg-green-50 border border-green-200 text-green-800' 
              : 'bg-red-50 border border-red-200 text-red-800'
          }`}>
            {message.type === 'success' ? (
              <Check className="w-5 h-5" />
            ) : (
              <AlertCircle className="w-5 h-5" />
            )}
            {message.text}
          </div>
        )}

        <div className="space-y-6">
          {/* Theme Settings */}
          <div className="card">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Appearance</h3>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Theme
                </label>
                <select
                  value={config.Theme}
                  onChange={(e) => handleConfigChange('Theme', e.target.value)}
                  className="input-field"
                >
                  <option value="light">Light</option>
                  <option value="dark">Dark</option>
                  <option value="auto">Auto (System)</option>
                </select>
              </div>
            </div>
          </div>

          {/* General Settings */}
          <div className="card">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">General</h3>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <div>
                  <label className="text-sm font-medium text-gray-700">
                    Auto-save changes
                  </label>
                  <p className="text-sm text-gray-500">
                    Automatically save todos and settings changes
                  </p>
                </div>
                <label className="relative inline-flex items-center cursor-pointer">
                  <input
                    type="checkbox"
                    checked={config.AutoSave}
                    onChange={(e) => handleConfigChange('AutoSave', e.target.checked)}
                    className="sr-only peer"
                  />
                  <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-600"></div>
                </label>
              </div>

              <div className="flex items-center justify-between">
                <div>
                  <label className="text-sm font-medium text-gray-700">
                    Notifications
                  </label>
                  <p className="text-sm text-gray-500">
                    Show system notifications for completed tasks
                  </p>
                </div>
                <label className="relative inline-flex items-center cursor-pointer">
                  <input
                    type="checkbox"
                    checked={config.Notifications}
                    onChange={(e) => handleConfigChange('Notifications', e.target.checked)}
                    className="sr-only peer"
                  />
                  <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-600"></div>
                </label>
              </div>
            </div>
          </div>

          {/* Todo Settings */}
          <div className="card">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Todo Settings</h3>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Default Category
                </label>
                <input
                  type="text"
                  value={config.DefaultTodoCategory}
                  onChange={(e) => handleConfigChange('DefaultTodoCategory', e.target.value)}
                  className="input-field"
                  placeholder="General"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Maximum Todos
                </label>
                <input
                  type="number"
                  value={config.MaxTodos}
                  onChange={(e) => handleConfigChange('MaxTodos', parseInt(e.target.value) || 100)}
                  className="input-field"
                  min="1"
                  max="1000"
                />
                <p className="text-sm text-gray-500 mt-1">
                  Maximum number of todos to keep in the list
                </p>
              </div>
            </div>
          </div>

          {/* OpenAI Settings */}
          <div className="card">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">OCR Settings</h3>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  OpenAI API Key
                </label>
                <input
                  type="password"
                  value={config.OpenAIAPIKey}
                  onChange={(e) => handleConfigChange('OpenAIAPIKey', e.target.value)}
                  className="input-field"
                  placeholder="sk-..."
                />
                <p className="text-sm text-gray-500 mt-1">
                  Your OpenAI API key for OCR functionality. Get one from{' '}
                  <a href="https://platform.openai.com/api-keys" target="_blank" rel="noopener noreferrer" className="text-primary-600 hover:text-primary-700 underline">
                    OpenAI Platform
                  </a>
                </p>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  OpenAI Base URL
                </label>
                <input
                  type="url"
                  value={config.OpenAIBaseURL}
                  onChange={(e) => handleConfigChange('OpenAIBaseURL', e.target.value)}
                  className="input-field"
                  placeholder="https://api.moonshot.cn/v1"
                />
                <p className="text-sm text-gray-500 mt-1">
                  Base URL for the OpenAI-compatible API (e.g., Moonshot, OpenAI, etc.)
                </p>
              </div>
            </div>
          </div>

          {/* Language Settings */}
          <div className="card">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Language</h3>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Interface Language
              </label>
              <select
                value={config.Language}
                onChange={(e) => handleConfigChange('Language', e.target.value)}
                className="input-field"
              >
                <option value="en">English</option>
                <option value="es">Español</option>
                <option value="fr">Français</option>
                <option value="de">Deutsch</option>
                <option value="zh">中文</option>
              </select>
            </div>
          </div>

          {/* Action Buttons */}
          <div className="flex gap-4">
            <button
              onClick={handleSaveConfig}
              disabled={saving}
              className="btn-primary flex items-center gap-2"
            >
                  {saving ? (
                    <div className="w-4 h-4 border-2 border-gray-300 border-t-gray-600 rounded-full animate-spin" />
                  ) : (
                    <Save className="w-4 h-4" />
                  )}
                  {saving ? 'Saving...' : 'Save Settings'}
                </button>
                <button
                  onClick={resetToDefaults}
                  className="btn-secondary flex items-center gap-2"
                >
                  <RotateCcw className="w-4 h-4" />
                  Reset to Defaults
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Settings
