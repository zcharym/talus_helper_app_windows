import { useState } from 'react'
import { OptimizePrompt } from '@wailsjs/go/main/App'
import { Sparkles, Copy, Check, AlertCircle, Loader2 } from 'lucide-react'

type OptimizationStyle = 'clarity' | 'specificity' | 'ai-focused' | 'concise' | 'detailed'

const optimizationStyles: { value: OptimizationStyle; label: string; description: string }[] = [
  { value: 'clarity', label: 'Clarity', description: 'Improve clarity and readability' },
  { value: 'specificity', label: 'Specificity', description: 'Add more specific details' },
  { value: 'ai-focused', label: 'AI-Focused', description: 'Optimize for AI/LLM understanding' },
  { value: 'concise', label: 'Concise', description: 'Make more concise' },
  { value: 'detailed', label: 'Detailed', description: 'Add more detail and context' }
]

function PromptOptimizer() {
  const [originalPrompt, setOriginalPrompt] = useState('')
  const [optimizedPrompt, setOptimizedPrompt] = useState('')
  const [selectedStyle, setSelectedStyle] = useState<OptimizationStyle>('clarity')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [copied, setCopied] = useState(false)

  const handleOptimize = async () => {
    if (!originalPrompt.trim()) {
      setError('Please enter a prompt to optimize')
      return
    }

    setLoading(true)
    setError(null)
    setOptimizedPrompt('')

    try {
      const result = await OptimizePrompt(originalPrompt.trim(), selectedStyle)
      setOptimizedPrompt(result)
    } catch (err) {
      console.error('Failed to optimize prompt:', err)
      setError(err instanceof Error ? err.message : 'Failed to optimize prompt')
    } finally {
      setLoading(false)
    }
  }

  const handleCopy = async () => {
    if (!optimizedPrompt) return

    try {
      await navigator.clipboard.writeText(optimizedPrompt)
      setCopied(true)
      setTimeout(() => setCopied(false), 2000)
    } catch (err) {
      console.error('Failed to copy:', err)
    }
  }

  return (
    <div className="px-4 py-6 sm:px-0">
      <div className="max-w-4xl mx-auto">
        <div className="mb-8">
          <h2 className="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-2 flex items-center gap-2">
            <Sparkles className="w-8 h-8" />
            Prompt Optimizer
          </h2>
          <p className="text-gray-600 dark:text-gray-400">
            Enter your prompt below and optimize it for better clarity, specificity, or AI understanding.
          </p>
        </div>

        {/* Error Message */}
        {error && (
          <div className="mb-4 p-4 message-error rounded-lg flex items-center gap-2">
            <AlertCircle className="w-5 h-5" />
            {error}
          </div>
        )}

        {/* Optimization Style Selector */}
        <div className="mb-6">
          <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
            Optimization Style
          </label>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
            {optimizationStyles.map((style) => {
              const isSelected = selectedStyle === style.value
              return (
                <button
                  key={style.value}
                  onClick={() => !loading && setSelectedStyle(style.value)}
                  disabled={loading}
                  className={`
                    relative p-4 rounded-lg border-2 transition-all duration-200 text-left
                    ${isSelected
                      ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20 shadow-md'
                      : 'border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 hover:border-primary-300 dark:hover:border-primary-700 hover:shadow-sm'
                    }
                    ${loading ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer'}
                  `}
                >
                  {isSelected && (
                    <div className="absolute top-2 right-2">
                      <div className="w-5 h-5 rounded-full bg-primary-500 flex items-center justify-center">
                        <Check className="w-3 h-3 text-white" />
                      </div>
                    </div>
                  )}
                  <div className="pr-6">
                    <div className={`font-semibold mb-1 ${isSelected ? 'text-primary-700 dark:text-primary-300' : 'text-gray-900 dark:text-gray-100'}`}>
                      {style.label}
                    </div>
                    <div className="text-sm text-gray-600 dark:text-gray-400">
                      {style.description}
                    </div>
                  </div>
                </button>
              )
            })}
          </div>
        </div>

        {/* Original Prompt Input */}
        <div className="mb-6">
          <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Original Prompt
          </label>
          <textarea
            value={originalPrompt}
            onChange={(e) => setOriginalPrompt(e.target.value)}
            placeholder="Enter your prompt here..."
            className="input-field min-h-[150px] resize-y"
            disabled={loading}
          />
        </div>

        {/* Optimize Button */}
        <div className="mb-6">
          <button
            onClick={handleOptimize}
            disabled={loading || !originalPrompt.trim()}
            className="btn-primary flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {loading ? (
              <>
                <Loader2 className="w-4 h-4 animate-spin" />
                Optimizing...
              </>
            ) : (
              <>
                <Sparkles className="w-4 h-4" />
                Optimize Prompt
              </>
            )}
          </button>
        </div>

        {/* Optimized Prompt Display */}
        {optimizedPrompt && (
          <div className="card">
            <div className="flex items-center justify-between mb-4">
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                Optimized Prompt
              </label>
              <button
                onClick={handleCopy}
                className="btn-secondary flex items-center gap-2 text-sm"
                title="Copy optimized prompt"
              >
                {copied ? (
                  <>
                    <Check className="w-4 h-4" />
                    Copied!
                  </>
                ) : (
                  <>
                    <Copy className="w-4 h-4" />
                    Copy
                  </>
                )}
              </button>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700">
              <p className="text-gray-900 dark:text-gray-100 whitespace-pre-wrap">
                {optimizedPrompt}
              </p>
            </div>
          </div>
        )}

        {/* Empty State */}
        {!optimizedPrompt && !loading && (
          <div className="card text-center py-12">
            <div className="text-gray-400 dark:text-gray-600 mb-4">
              <Sparkles className="w-16 h-16 mx-auto" />
            </div>
            <h3 className="text-lg font-medium text-gray-900 dark:text-gray-100 mb-2">
              Ready to optimize
            </h3>
            <p className="text-gray-500 dark:text-gray-400">
              Enter a prompt above and click "Optimize Prompt" to get started.
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

export default PromptOptimizer

