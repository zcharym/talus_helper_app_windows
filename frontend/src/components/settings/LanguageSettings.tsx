import { useState, useEffect } from 'react'
import { GetConfig, SaveConfig } from '@wailsjs/go/main/App'
import { AppConfig } from '../../types'
import { Globe, Check } from 'lucide-react'

const languages = [
  { code: 'en', name: 'English', flag: '🇺🇸' },
  { code: 'es', name: 'Español', flag: '🇪🇸' },
  { code: 'fr', name: 'Français', flag: '🇫🇷' },
  { code: 'de', name: 'Deutsch', flag: '🇩🇪' },
  { code: 'zh', name: '中文', flag: '🇨🇳' },
  { code: 'ja', name: '日本語', flag: '🇯🇵' },
  { code: 'ko', name: '한국어', flag: '🇰🇷' },
  { code: 'pt', name: 'Português', flag: '🇵🇹' },
  { code: 'ru', name: 'Русский', flag: '🇷🇺' },
  { code: 'it', name: 'Italiano', flag: '🇮🇹' }
]

function LanguageSettings() {
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
      setMessage({ type: 'success', text: 'Language settings saved successfully!' })
      setTimeout(() => setMessage(null), 3000)
    } catch (error) {
      console.error('Failed to save config:', error)
      setMessage({ type: 'error', text: 'Failed to save language settings' })
    } finally {
      setSaving(false)
    }
  }

  const getCurrentLanguage = () => {
    return languages.find(lang => lang.code === config?.Language) || languages[0]
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
          Language Settings
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          Choose your preferred interface language
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
        {/* Current Language */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4 flex items-center gap-2">
            <Globe className="w-5 h-5" />
            Current Language
          </h3>
          <div className="flex items-center gap-3 p-4 bg-primary-50 dark:bg-primary-900/20 rounded-lg border border-primary-200 dark:border-primary-800">
            <span className="text-2xl">{getCurrentLanguage().flag}</span>
            <div>
              <div className="font-medium text-primary-900 dark:text-primary-100">
                {getCurrentLanguage().name}
              </div>
              <div className="text-sm text-primary-700 dark:text-primary-300">
                {getCurrentLanguage().code.toUpperCase()}
              </div>
            </div>
            <Check className="w-5 h-5 text-primary-600 dark:text-primary-400 ml-auto" />
          </div>
        </div>

        {/* Language Selection */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">
            Available Languages
          </h3>
          <div className="grid grid-cols-1 gap-2">
            {languages.map((language) => (
              <button
                key={language.code}
                onClick={() => handleConfigChange('Language', language.code)}
                className={`flex items-center gap-3 p-3 rounded-lg border-2 transition-colors text-left ${
                  config.Language === language.code
                    ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20 text-primary-700 dark:text-primary-300'
                    : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
                }`}
              >
                <span className="text-xl">{language.flag}</span>
                <div className="flex-1">
                  <div className="font-medium text-gray-900 dark:text-gray-100">
                    {language.name}
                  </div>
                  <div className="text-sm text-gray-500 dark:text-gray-400">
                    {language.code.toUpperCase()}
                  </div>
                </div>
                {config.Language === language.code && (
                  <Check className="w-5 h-5 text-primary-600 dark:text-primary-400" />
                )}
              </button>
            ))}
          </div>
        </div>

        {/* Language Information */}
        <div className="card">
          <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">
            About Language Support
          </h3>
          <div className="space-y-3 text-sm text-gray-600 dark:text-gray-400">
            <p>
              Talus Helper supports multiple languages for the interface. Changing the language will update:
            </p>
            <ul className="list-disc list-inside space-y-1 ml-4">
              <li>Menu items and navigation</li>
              <li>Button labels and tooltips</li>
              <li>Settings descriptions</li>
              <li>Error messages and notifications</li>
            </ul>
            <p className="text-xs text-gray-500 dark:text-gray-500 mt-3">
              Note: Some features like OCR text extraction may still work in the original language of the content.
            </p>
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
            {saving ? 'Saving...' : 'Save Language Settings'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default LanguageSettings
