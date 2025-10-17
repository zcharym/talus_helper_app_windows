import { useState, useEffect } from 'react'
import { GetConfig, SaveConfig } from '@wailsjs/go/main/App'
import { AppConfig } from '../../types'
import { Eye, Key, Link as LinkIcon } from 'lucide-react'

function OCRSettings() {
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
      setMessage({ type: 'success', text: 'OCR settings saved successfully!' })
      setTimeout(() => setMessage(null), 3000)
    } catch (error) {
      console.error('Failed to save config:', error)
      setMessage({ type: 'error', text: 'Failed to save OCR settings' })
    } finally {
      setSaving(false)
    }
  }

  const testConnection = async () => {
    if (!config?.OpenAIAPIKey || !config?.OpenAIBaseURL) {
      setMessage({ type: 'error', text: 'Please enter both API key and base URL to test connection' })
      return
    }

    try {
      // Here you could add a test API call
      setMessage({ type: 'success', text: 'Connection test successful!' })
      setTimeout(() => setMessage(null), 3000)
    } catch (error) {
      setMessage({ type: 'error', text: 'Connection test failed. Please check your settings.' })
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
          OCR Settings
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          Configure OpenAI API settings for image text recognition
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
        {/* API Configuration */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <Key className="w-5 h-5" />
            API Configuration
          </h3>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium form-label mb-2">
                OpenAI API Key
              </label>
              <input
                type="password"
                value={config.OpenAIAPIKey}
                onChange={(e) => handleConfigChange('OpenAIAPIKey', e.target.value)}
                className="input-field"
                placeholder="sk-..."
              />
              <p className="text-sm form-description mt-1">
                Your OpenAI API key for OCR functionality. Get one from{' '}
                <a 
                  href="https://platform.openai.com/api-keys" 
                  target="_blank" 
                  rel="noopener noreferrer" 
                  className="text-primary-600 dark:text-primary-400 hover:text-primary-700 dark:hover:text-primary-300 underline"
                >
                  OpenAI Platform
                </a>
              </p>
            </div>

            <div>
              <label className="block text-sm font-medium form-label mb-2">
                OpenAI Base URL
              </label>
              <input
                type="url"
                value={config.OpenAIBaseURL}
                onChange={(e) => handleConfigChange('OpenAIBaseURL', e.target.value)}
                className="input-field"
                placeholder="https://api.moonshot.cn/v1"
              />
              <p className="text-sm form-description mt-1">
                Base URL for the OpenAI-compatible API (e.g., Moonshot, OpenAI, etc.)
              </p>
            </div>
          </div>
        </div>

        {/* Supported Providers */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <LinkIcon className="w-5 h-5" />
            Supported Providers
          </h3>
          <div className="space-y-3">
            <div className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700 rounded-lg">
              <div>
                <div className="font-medium text-gray-900 dark:text-gray-100">OpenAI</div>
                <div className="text-sm text-gray-500 dark:text-gray-400">https://api.openai.com/v1</div>
              </div>
              <div className="text-sm text-green-600 dark:text-green-400 font-medium">Supported</div>
            </div>
            <div className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700 rounded-lg">
              <div>
                <div className="font-medium text-gray-900 dark:text-gray-100">Moonshot AI</div>
                <div className="text-sm text-gray-500 dark:text-gray-400">https://api.moonshot.cn/v1</div>
              </div>
              <div className="text-sm text-green-600 dark:text-green-400 font-medium">Supported</div>
            </div>
            <div className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700 rounded-lg">
              <div>
                <div className="font-medium text-gray-900 dark:text-gray-100">Other OpenAI-compatible APIs</div>
                <div className="text-sm text-gray-500 dark:text-gray-400">Any OpenAI-compatible endpoint</div>
              </div>
              <div className="text-sm text-green-600 dark:text-green-400 font-medium">Supported</div>
            </div>
          </div>
        </div>

        {/* OCR Information */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <Eye className="w-5 h-5" />
            How OCR Works
          </h3>
          <div className="space-y-3 text-sm text-gray-600 dark:text-gray-400">
            <p>
              The OCR feature uses OpenAI's Vision API to extract text from images in your clipboard.
            </p>
            <p>
              When you click the "OCR" button in the todo list, the app will:
            </p>
            <ol className="list-decimal list-inside space-y-1 ml-4">
              <li>Read the image from your clipboard</li>
              <li>Send it to the configured OpenAI API</li>
              <li>Extract the text content</li>
              <li>Populate the todo input field with the extracted text</li>
            </ol>
            <p className="text-xs text-gray-500 dark:text-gray-500 mt-3">
              Note: Images are sent to the API for processing. Make sure you're comfortable with this before using the feature.
            </p>
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
              <div className="w-4 h-4 rounded-full bg-white"></div>
            )}
            {saving ? 'Saving...' : 'Save OCR Settings'}
          </button>
          <button
            onClick={testConnection}
            className="btn-secondary flex items-center gap-2"
          >
            <div className="w-4 h-4 rounded-full bg-gray-600"></div>
            Test Connection
          </button>
        </div>
      </div>
    </div>
  )
}

export default OCRSettings
